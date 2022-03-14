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

type Checker struct {
}

type LineParser struct {
}

const (
	COMMENT = iota
	EMPTY
	NODE
	KV
)

func (c *Checker) judge_line_kind(line Line) (kind int, err error) {
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

func (c *Checker) check_node_type(data interface{}) (err error) {
	// 检查待解析的结构体类型
	t := reflect.TypeOf(data)
	if t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct {
		err = nil
	} else {
		err = fmt.Errorf("data[%v] should be a struct pointer", data)
	}
	return err
}

func (c *Checker) check_kv_type(nodeStruct reflect.Value) (err error) {
	// 检查kv是否是struct类型
	nodeType := reflect.TypeOf(nodeStruct)
	if nodeType.Kind() != reflect.Struct {
		err = fmt.Errorf("node[%v] is not struct", nodeType.Name())
	}
	return err
}

func (p *LineParser) parse_k_v(line Line) (k, v string) {
	idx := strings.Index(line.content, "=")
	k = strings.TrimSpace(line.content[:idx])
	v = strings.TrimSpace(line.content[idx+1:])
	return k, v
}

func (p *LineParser) parse_node(line Line) (node string) {
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

func loadIni(path string, data interface{}) (err error) {
	// node --- input: nodeTag(field tag) nodeString(ini value)  middle: nodeType(field type)  output: nodeValue(field value)
	checker := Checker{}
	parser := LineParser{}
	var nodeStruct reflect.Value
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
		if kind, e := checker.judge_line_kind(line); e == nil {
			switch kind {
			case COMMENT, EMPTY:
				continue
			case KV:
				tag, value := parser.parse_k_v(line)
				// 反射值对象调Type()方法得到反射类型对象
				if err = checker.check_kv_type(nodeStruct); err != nil {
					return err
				}
				field := getFieldByTag(nodeStruct, tag)
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
				// 传入的 data 参数是否是一个结构体指针类型（要对结构体进行字段补充）
				if err = checker.check_node_type(data); err != nil {
					return err
				}
				nodeTagString := parser.parse_node(line)
				nodeStruct = getFieldByTag(reflect.ValueOf(data), nodeTagString)
			}
		}
	}
	return
}
