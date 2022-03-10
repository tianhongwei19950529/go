package main

import (
	"fmt"
)

func count(lst1 [5]int) {
	//作业题
	num := 0
	for i := 0; i < len(lst1); i++ {
		num += lst1[i]
	}
	fmt.Println(num)

}

func forlist()  {
	//第一种for循环 利用长度
	lst1 := [...]string{"山东","辽宁","河南","河北"}
	for i := 0; i < len(lst1); i++ {
		fmt.Println(lst1[i])
	}
}

func forrange()  {
	lst1 := [...]string{"山东","辽宁","河南","河北"}
	for s1, s2 := range lst1 {
		fmt.Println(s1,s2)
	}
}

func cityList(){
	erList := [3][2]string{
		{"山东","济南"},
		{"辽宁","大连"},
		{"河北","石家庄"},
	}
	for _,k := range erList{
		for _,k1 := range k{
			fmt.Println(k1)
		}
	}
}

func updatedata(){
	updatelist := [5]int{}
	updatelist1 := [5]int{}
	fmt.Println(updatelist)
	updatelist[2]=10
	fmt.Println(updatelist)

	updatelist1 = updatelist
	fmt.Println(updatelist1)
	updatelist1[2] =100
	fmt.Println(updatelist)
	fmt.Println(updatelist1)
}


func lianxi2(){
	a := [5]int{1, 3, 5, 7, 8}
	for i := 0; i < len(a); i++ {
		for y := i; y < len(a); y++ {

			if a[i] + a[y] == 8{
				fmt.Println(i,y)
			}
		}
	}
}


func main() {
	//lst1 := [5]int{1, 3, 5, 7, 8}
	//count(lst1)
	//lst1 := [3]int{1,2,3}
	//lst2 := [...]int{1,4,7,9}
	//lst3 := [...]int{1:1,6:6}
	//fmt.Println(lst1)
	//fmt.Println(lst2)
	//fmt.Println(lst3)

	//forlist()
	//forrange()

	//erList := [3][2]string{
	//	{"山东","济南"},
	//	{"辽宁","大连"},
	//	{"河北","石家庄"},
	//}
	//fmt.Println(erList[0])
	//cityList()
	//allList := [3][2]int{}
	//fmt.Println(allList)
	//updatedata()
	lianxi2()
}
