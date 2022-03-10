package main

import (
	"01基础/9package/calc"
	"fmt"
)

func main() {
	x, y := 20, 10
	fmt.Println(calc.Add(x, y))
	fmt.Println(calc.Red(x, y))
	fmt.Println(calc.Ride(x, y))
	fmt.Println(calc.Div(x, y))
}
