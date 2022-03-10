package main

import "fmt"

func main() {
	lst1 := []int{3, 3}
	b := 0
	for _, i2 := range lst1 {
		b = b ^ i2
	}
	fmt.Println(b)

}
