package parseini

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

type MysqlConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Database int    `ini:"database"`
	Password string `ini:"password"`
}

type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

type Line struct {
	no      int
	content string
}

const (
	COMMENT = iota
	EMPTY
	NODE
	KV
)

type IChecker interface {
	check() (err error)
	setV(v reflect.Value)
}

type NodeChecker struct {
	v reflect.Value
}

type KVChecker struct {
	v reflect.Value
}

func (c NodeChecker) check() (err error) {
	// 检查待解析的结构体类型
	//t := reflect.TypeOf(data)
	t := c.v.Type()
	if t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct {
		err = nil
	} else {
		err = fmt.Errorf("data[%v] should be a struct pointer", t.Name())
	}
	return err
}

func (c *NodeChecker) setV(v reflect.Value) {
	(*c).v = v
}

func (c KVChecker) check() (err error) {
	// 检查kv是否是struct类型
	//nodeType := reflect.TypeOf(c.v)
	t := c.v.Type()
	if t.Kind() != reflect.Struct {
		err = fmt.Errorf("node[%v] is not struct", t.Name())
	}
	return err
}

func (c *KVChecker) setV(v reflect.Value) { // 法克！！！Fuck 指针赋值！！！
	(*c).v = v
}

func (line *Line) parse_kind() (kind int, err error) {
	// 检查ini中每一行的类型
	stripedLine := strings.TrimSpace(line.content)
	if len(stripedLine) == 0 {
		kind = EMPTY
	} else if strings.HasPrefix(stripedLine, ";") {
		kind = COMMENT
	} else if strings.HasPrefix(stripedLine, "[") {
		kind = NODE
		if stripedLine[0] != '[' || stripedLine[len(stripedLine)-1] != ']' {
			err = fmt.Errorf("line no[%d] content[%s] error: 两侧[]有误", line.no, line.content)
			return kind, err
		}
		if sectionName := strings.TrimSpace(stripedLine[1 : len(stripedLine)-1]); len(sectionName) == 0 {
			err = fmt.Errorf("line no[%d] content[%s] error: 中间值有误", line.no, line.content)
			return kind, err
		}
	} else if strings.Index(stripedLine, "=") > 0 {
		kind = KV
		if strings.HasPrefix(stripedLine, "=") || strings.HasSuffix(stripedLine, "=") {
			err = fmt.Errorf("line no[%d] content[%s] error: 缺少=", line.no, line.content)
			return kind, err
		}
	}
	return kind, nil
}

func (line *Line) parse_k_v() (k, v string) {
	idx := strings.Index(line.content, "=")
	k = strings.TrimSpace(line.content[:idx])
	v = strings.TrimSpace(line.content[idx+1:])
	return k, v
}

func (line *Line) parse_node() (node string) {
	node = strings.TrimSpace(line.content)[1 : len(line.content)-1]
	return node
}

func getFieldByTag(v reflect.Value, tag string) (field reflect.Value) {
	// 根据 string 取 node 的 field 名; 指针 or 结构体
	t := v.Type()
	var fieldName string
	switch v.Kind() {
	case reflect.Ptr:
		for i := 0; i < v.Elem().NumField(); i++ {
			if f := t.Elem().Field(i); f.Tag.Get("ini") == tag {
				fieldName = f.Name
				break
			}
		}
		if fieldName != "" {
			field = v.Elem().FieldByName(fieldName)
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			if f := t.Field(i); f.Tag.Get("ini") == tag {
				fieldName = f.Name
				break
			}
		}
		if fieldName != "" {
			field = v.FieldByName(fieldName)
		}
	}
	return
}

func doCheck(checker IChecker, v reflect.Value) (err error) {
	checker.setV(v)
	err = checker.check()
	return
}

func loadIni(path string, data interface{}) (err error) {
	// node --- input: nodeTag(field tag) nodeString(ini value)  middle: nodeType(field type)  output: nodeValue(field value)
	var v reflect.Value
	// 1.读文件，得到字节类型，转化为字符串
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	// 2. 一行一行读文件，如果是注释就忽略
	// 3. 一行一行读数据，如果是[开头则代表就是一个节点
	// 4. 一行一行读数据，如果不是[开头的就是=分割的键值对
	lineSlice := strings.Split(string(b), "\n")
	for idx, lineStr := range lineSlice {
		line := Line{no: idx, content: lineStr}
		if kind, e := line.parse_kind(); e == nil {
			switch kind {
			case COMMENT, EMPTY:
				continue
			case KV:
				tag, value := line.parse_k_v()
				if err = doCheck(new(KVChecker), v); err != nil {
					//if err = doCheck(&KVChecker{}, v); err != nil {  // 被指针赋值坑了
					return
				}
				field := getFieldByTag(v, tag)
				// 赋值
				switch field.Type().Kind() {
				case reflect.String:
					field.SetString(value)
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					var valueInt int64
					valueInt, err = strconv.ParseInt(value, 10, 64)
					if err != nil {
						err = fmt.Errorf("line:%d value type err", idx+1)
						return
					}
					field.SetInt(valueInt)
				case reflect.Bool:
					var valueBool bool
					valueBool, err = strconv.ParseBool(value)
					if err != nil {
						err = fmt.Errorf("line:%d value type err", idx+1)
						return err
					}
					field.SetBool(valueBool)
				case reflect.Float32, reflect.Float64:
					var valueFloat float64
					valueFloat, err = strconv.ParseFloat(value, 64)
					if err != nil {
						err = fmt.Errorf("line:%d value type err", idx+1)
						return err
					}
					field.SetFloat(valueFloat)
				}
			case NODE:
				//	校验参数
				v = reflect.ValueOf(data)
				if err = doCheck(new(NodeChecker), v); err != nil {
					return
				}
				nodeTagString := line.parse_node()
				v = getFieldByTag(v, nodeTagString)
			}
		}
	}
	return
}
