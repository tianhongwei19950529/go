package main

import "fmt"

func main() {
	//a := "what are you talking about?"
	//b := ""
	//fmt.Println(len(a))
	//fmt.Println(a + b)
	//c := fmt.Sprintf("%s  - %s", a, b)
	//fmt.Println(c)
	//d := strings.Split(b, " ")
	//fmt.Println(d)
	////任何字符串都包含一个空字符串
	//fmt.Println(strings.Contains(a, b))
	////查看字符串出现的位置
	//fmt.Println(strings.Index(a, "are"))
	//charList := []string{"what", "are", "you"}
	//fmt.Println(strings.Join(charList,"+"))
	str1 := "hello 沙河"
	for i := 0; i < len(str1); i++ {
		fmt.Printf("%c", str1[i])
	}
	fmt.Println()
	for _, i := range str1 {
		fmt.Printf("%c", i)
	}
	fmt.Println()
	var rune1 = []rune(str1)
	fmt.Printf("%c \n", rune1[6])

	 var a int8
	a = 127
	fmt.Println(a+2)
}
