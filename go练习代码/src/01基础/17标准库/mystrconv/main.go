// strconv 包实现了基本数据类型与其字符串表示的转换，常用：Atoi，Itoa，parse系列，format系列，append系列
package main

import (
	"fmt"
	"strconv"
)

// Atoi  asicc to int, 要接收 err, 因为可能出错
func main1() {
	s1 := "100"
	i1, err := strconv.Atoi(s1)
	fmt.Println(i1, err)
}

// Itoa  int to ascii，不接收err，因为不会不错
func main2() {
	i2 := 100
	s2 := strconv.Itoa(i2)
	fmt.Println(s2)
}

// parse 系列，用于转换字符串为给定类型的值：ParseBool，ParseFloat, ParseInt, parseUint
func main3() {
	b, err := strconv.ParseBool("true")
	f, err := strconv.ParseFloat("3.14", 64) // bitSize: 位
	i, err := strconv.ParseInt("-1", 10, 64)
	u, err := strconv.ParseUint("2", 10, 64)
	fmt.Println(b, f, i, u, err)
}

// format 系列函数实现了将给定类型数据格式化为string类型数据的功能
func main4() {
	s1 := strconv.FormatBool(true)
	s2 := strconv.FormatFloat(3.14, 'E', -1, 64)
	s3 := strconv.FormatInt(-2, 16)    // -2转成16进制字符串
	s4 := strconv.FormatUint(2000, 16) // 2 转成无符号16进制
	fmt.Println(s1, s2, s3, s4)
}

// isPrint() 返回一个字符是否是可打印的
// CanBackquote() 返回一个字符串是否可以不被修改的 表示为一个单行的、没有空格和tab之外的控制字符的反引号字符串
func main5() {
	a := strconv.IsPrint('f')
	b := strconv.CanBackquote("fuck")
	fmt.Println(a, b)
}

// append系列，将值添加到现有字符串中
func main6() {
	s := []byte{'a', 'b'}
	s = strconv.AppendBool(s, true)
	s = strconv.AppendInt(s, 4567, 10) // 10 进制
	fmt.Println(s)
}

// quote 系列，Quote: 将``转换为""字符串; QuoteToASCII 将字符串s转换为双引号引起来的ascii字符串(unicode)
func main7() {
	fmt.Println(strconv.Quote(`fuck`))
	fmt.Println(`fuck`)
	fmt.Println(`"fuck"`)
	for i, c := range "fuck" {
		fmt.Printf("index: %d, char: %c\n", i, c)
	}
	for i, c := range strconv.Quote(`fuck`) {
		fmt.Printf("index: %d, char: %c\n", i, c)
	}
	fmt.Println(strconv.QuoteToASCII(`宝贝`))

}

func main() {
	//main1()
	//main2()
	//main3()
	//main4()
	//main5()
	//main6()
	main7()
}
