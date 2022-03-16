package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

func initMap() {
	var stuMap map[string]int
	fmt.Println(stuMap == nil) //true
	stuMap = make(map[string]int, 3)
	stuMap["thw"] = 100
	stuMap["thw1"] = 100
	fmt.Println(stuMap)
	stuMap1 := map[int]bool{
		1: true,
		2: false,
	}
	stuMap1[3] = true
	fmt.Println(stuMap1)
}

func hasKeyMap() {
	var b []int
	var scoreMap = map[string][]int{
		"孙悟空":   b,
		"如来":    make([]int, 10),
		"唐三藏":   make([]int, 10),
		"女儿国国王": make([]int, 10),
	}
	value, ok := scoreMap["孙悟空"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("Nil")
	}
}

func deleteMap(map1 map[string]int) map[string]int {
	delete(map1, "12")
	return map1
}

func forMap() {
	var scoreMap = map[string]int{
		"孙悟空":   100,
		"如来":    90,
		"唐三藏":   80,
		"女儿国国王": 70,
	}
	fmt.Println("拿到全部")
	for index, value := range scoreMap {
		fmt.Println(index, value)
	}
	fmt.Println("拿到key值")
	for index := range scoreMap {
		fmt.Println(index)

	}
	fmt.Println("拿到value值")
	for _, value := range scoreMap {
		fmt.Println(value)
	}
}

func forOrderMap() {
	scoreMap := make(map[string]int, 50)
	for i := 1; i <= 50; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		scoreMap[key] = value
	}
	fmt.Println(scoreMap)
	//构建可以排序的切片
	var orderSlice = make([]string, 0, 100)
	for i := range scoreMap {
		orderSlice = append(orderSlice, i)
	}
	//切片排序
	sort.Strings(orderSlice)
	//根据切片输出for循环的值
	for i := 0; i < len(orderSlice)-1; i++ {
		fmt.Println(orderSlice[i], scoreMap[orderSlice[i]])
	}
}

func createMapSlice() {
	Mapslice := make([]map[string]string, 10)
	//Map1 := make(map[string]string,3)
	//Map1["name"] = "thw"
	//Mapslice =append(Mapslice,Map1)
	//fmt.Println(Mapslice[0]["name"])
	Mapslice[0] = make(map[string]string, 5)
	Mapslice[0]["name"] = "beijing"
	fmt.Println(Mapslice[0])
}

func createSliceMap() {
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	//key := "中国"
	//value, ok := sliceMap[key]
	//if !ok {
	//	value = make([]string, 0, 2)
	//}
	//value = append(value, "北京", "上海")
	//sliceMap[key] = value
	for _, i2 := range sliceMap {
		if i2 == nil {
			i2 = make([]string, 8)
		}
	}
	fmt.Println(sliceMap)
}
func zuoye1() {
	str1 := "how do you do"
	strSlice := strings.Split(str1, " ")
	var strMap = map[string]int{}
	for _, v := range strSlice {
		_, ok := strMap[v]
		if ok {
			strMap[v] += 1
		} else {
			strMap[v] = 1
		}
	}
	fmt.Println(strMap)
}

func zuoye2() {
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s) //{1,2,3}
	m["q1mi"] = s          //{"qimi":[1,2,3]}
	fmt.Println(m)
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\n", s)         //[1,3]
	fmt.Printf("%+v\n", m["q1mi"]) //{"qimi":[1,3,3]}
}
func main() {
	//initMap()
	//hasKeyMap()
	var b = map[string]int{
		"123": 123,
		"321": 321,
		"12":  12,
	}
	fmt.Println(b)
	deleteMap(b)
	//b["123"] = 1234
	fmt.Println(b)
	//forMap()
	//forOrderMap()
	//zuoye1()
	//zuoye2()
	//createMapSlice()
	createSliceMap()
}
