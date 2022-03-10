package main

import (
	"fmt"
	"sort"
)

func dingyi() {
	var name []string
	var ageSlice []int
	fmt.Println(name, ageSlice)
}

func lenquery(all []int) {
	for i, i2 := range all {
		fmt.Println(i, i2)
	}
}

func getSliceByArray() {
	a := [5]int{1, 2, 3, 4, 5}
	sliceA := a[:3]
	a[0] = 10
	fmt.Printf("%#v,%T,%p,%d,%d \n", sliceA, sliceA, sliceA, len(sliceA), cap(sliceA))
	for i := 0; i < 10; i++ {
		sliceA = append(sliceA, i)
	}
	fmt.Printf("%#v,%T,%p,%d,%d \n", sliceA, sliceA, sliceA, len(sliceA), cap(sliceA))
	fmt.Printf("%#v,%T,%p,%d,%d \n", sliceA, sliceA, sliceA, len(sliceA), cap(sliceA))
}

func getSliceBySlice() {
	a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("%#v,%T,%p,%d,%d \n", a, a, &a, len(a), cap(a))
	slice1 := a[:3]
	fmt.Printf("%#v,%T,%p,%d,%d \n", slice1, slice1, slice1, len(slice1), cap(slice1))
	slice2 := slice1[:8]
	fmt.Printf("%#v,%T,%p,%d,%d \n", slice2, slice2, slice2, len(slice2), cap(slice2))
}

func getslicebymake() {
	//通过make构建切片
	slice1 := make([]int, 10, 10)
	fmt.Printf("%#v,%T,%p,%d,%d \n", slice1, slice1, slice1, len(slice1), cap(slice1))
	slice2 := make([]int, 0, 10)
	fmt.Printf("%#v,%T,%p,%d,%d \n", slice2, slice2, slice2, len(slice2), cap(slice2))
}

func ifSliceNil() {
	//判断切片是不是为空 为nil
	var slice1 []int
	if slice1 == nil {
		fmt.Println("slice 是个nil")
	}
	fmt.Printf("%#v,%T,%p,%d,%d \n", slice1, slice1, slice1, len(slice1), cap(slice1))
	slice2 := make([]int, 0, 0)
	if slice2 == nil {
		fmt.Println("slice2 是个nil")
	}
	fmt.Printf("%#v,%T,%p,%d,%d \n", slice2, slice2, slice2, len(slice2), cap(slice2))
}

func copySlice() {
	slice1 := make([]int, 3)
	slice2 := slice1
	fmt.Printf("更改之前 slice1 %#v,%T,%p,%d,%d \n", slice1, slice1, slice1, len(slice1), cap(slice1))
	fmt.Printf("更改之前 slice2  %#v,%T,%p,%d,%d \n", slice2, slice2, slice2, len(slice2), cap(slice2))
	slice2[0] = 100
	fmt.Printf("更改过后 slice1 %#v,%T,%p,%d,%d \n", slice1, slice1, slice1, len(slice1), cap(slice1))
	fmt.Printf("更改过后 slice2  %#v,%T,%p,%d,%d \n", slice2, slice2, slice2, len(slice2), cap(slice2))
}

func addaSlice() {
	a := make([]int, 3, 3)
	b := a
	b[0] = 100
	a = append(a, 1)

	fmt.Printf("%#v,%T,%p,%d,%d \n", a, a, a, len(a), cap(a))
	fmt.Printf("%#v,%T,%p,%d,%d \n", b, b, b, len(b), cap(b))
}

func addMoreSlice() {
	a := make([]int, 3, 3)
	b := make([]int, 3, 3)
	a = append(a, b...)
	fmt.Printf("%#v,%T,%p,%d,%d \n", a, a, a, len(a), cap(a))
}

func copytoSlice() {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	slice2 := make([]int, 3, 9)
	copy(slice2, slice1)
	fmt.Printf("%#v,%T,%p,%d,%d \n", slice2, slice2, slice2, len(slice2), cap(slice2))
	slice3 := slice2[1:]
	fmt.Printf("%#v,%T,%p,%d,%d \n", slice3, slice3, slice3, len(slice3), cap(slice3))
}

func deleteSlice() {
	slice1 := []int{1, 2, 3, 4, 5, 2, 6, 7, 8}
	//for i := 0; i < len(slice1); {
	//	if slice1[i] == 2 {
	//		slice1 = append(slice1[:i], slice1[i+1:]...)
	//	} else {
	//		i++
	//	}
	//}
	for i, k := range slice1 {
		if k == 2 {
			slice1 = append(slice1[:i], slice1[i+1:]...)
		}
	}
	fmt.Println(slice1)
}

func ceshi() {
	slice := []int{10, 20, 30, 40, 50}
	newSlice := slice[1:3]
	newSlice = append(newSlice, 60)
	fmt.Println(slice, newSlice)
}

func zuoye1() {
	var a = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(len(a))
}

func zuoye2()  {
	var a = [...]int{3, 7, 8, 9, 1}
	a1 := a[:]
	sort.Ints(a1)
	fmt.Println(a1)
}

func main() {
	//getSliceByArray()
	//getSliceBySlice()
	//getslicebymake()
	//ifSliceNil()
	//copySlice()
	//addaSlice()
	//addMoreSlice()
	//copytoSlice()
	//deleteSlice()
	//ceshi()
	//zuoye1()
	zuoye2()
}
