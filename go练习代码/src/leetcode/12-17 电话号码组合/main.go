package main

import "fmt"

/*
1. 看到所有组合 第一反应应该是搜索算法，深度优先（字典） (递归解决) 广度优先（队列）（两层for循环）
*/

var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var combinations []string

func letterCombinations(digits string) []string {
	combinations = []string{}
	if len(digits) == 0 {
		return combinations
	} else {
		backtrack(digits, 0, "")
		return combinations
	}
}

func backtrack(digits string, index int, combination string) {
	if index == len(digits) {
		combinations = append(combinations, combination)
	} else {
		digit := string(digits[index])
		letters := phoneMap[digit]
		lettersCount := len(letters)
		for i := 0; i < lettersCount; i++ {
			backtrack(digits, index+1, combination+string(letters[i]))
		}
	}
}

//func backtrack(digits string, index int, combination string) {
//	if index == len(digits) {
//		combinations = append(combinations, combination)
//	} else {
//		digit := string(digits[index])
//		letters := phoneMap[digit]
//		lettersCount := len(letters)
//		for i := 0; i < lettersCount; i++ {
//			backtrack(digits, index + 1, combination + string(letters[i]))
//		}
//	}
//}

func main() {
	fmt.Println(letterCombinations("23"))
}
