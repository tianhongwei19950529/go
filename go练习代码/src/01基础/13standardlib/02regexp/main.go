package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18"
	//searchIn2 := "John: 2578a34 William: 4567.23 Steve: 5632.18"
	//pat := "[0-9]+.[0-9]+"
	pat := `[0-9]+\.[0-9]+`

	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		return strconv.FormatFloat(v*2, 'f', 2, 32)
	}

	// byte 是否匹配
	if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
		fmt.Println("Match Found")
	}

	// 字符串是否匹配
	if ok2, _ := regexp.MatchString(pat, searchIn); ok2 {
		fmt.Println("Match Found")
	}

	re, _ := regexp.Compile(pat)
	// 将匹配到的部分替换为"##.#"
	str := re.ReplaceAllString(searchIn, "##.#")
	fmt.Println(str)
	//	匹配部分传入函数进行替换
	str2 := re.ReplaceAllStringFunc(searchIn, f)
	fmt.Println(str2)
	// 提取所有匹配的字符串, 返回一个数组
	res := re.FindAllString(searchIn, -1) // 第二个参数为索引长度
	fmt.Println("reuslt = ", res)
	// 提取所有匹配的字符串, 返回一个二维数组
	result := re.FindAllStringSubmatch(searchIn, -1)
	fmt.Println("reuslt = ", result)
	result2 := re.FindAllStringIndex(searchIn, -1)
	fmt.Println("reuslt = ", result2)
	result3 := re.FindAllStringSubmatchIndex(searchIn, -1)
	fmt.Println("reuslt = ", result3)
}
