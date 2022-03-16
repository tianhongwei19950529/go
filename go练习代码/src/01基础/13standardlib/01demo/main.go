package main

import (
	"fmt"
	"unsafe"
	//"container/list"
)

func main() {
	var i int64
	fmt.Printf("%d bytes of int64", unsafe.Sizeof(i))
}
