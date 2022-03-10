

1.  map的定义

    1.  map声明之后,需要进行初始化.在未初始化之前一直都是nil.并且不能直接用键值对.

    2.  通过make初始化的时候,可以指定容量

    3.  可以一直增加键值对,系统会自动扩容

        ```go
        	var stuMap  map[string]int
        	fmt.Println(stuMap==nil)//true
        	stuMap = make(map[string]int,3)
        	stuMap["thw"]=100
        	stuMap["thw1"]=100
        	fmt.Println(stuMap)
        	stuMap1 := map[int]bool{
        		1:true,
        		2:false,
        	}
        	stuMap1[3]=true
        	fmt.Println(stuMap1)
        ```

2.  常见操作

    1.  判断某个键在不在map里面

        ```go
        	var b  []int
        	var scoreMap = map[string][]int{
        		"孙悟空":b,
        		"如来":make([]int,10),
        		"唐三藏":make([]int,10),
        		"女儿国国王":make([]int,10),
        	}
        	value,ok := scoreMap["孙悟空"]
        	if ok{
        		fmt.Println(value)
        	} else {
        		fmt.Println("Nil")
        	}
        ```

    2.  删除map里面的某个值

        ```go
        	//删除已经存在的键值对直接删除,删除不存在的没有影响
        	var b = map[string]int{
        		"123":123,
        		"321":321,
        		"12":12,
        	}
        	delete(map1,"123")
        
        ```

3.  map遍历

    ```go
    	var scoreMap = map[string]int{
    		"孙悟空":100,
    		"如来":90,
    		"唐三藏":80,
    		"女儿国国王":70,
    	}
    	//-----------------------------------
    	fmt.Println("拿到全部")
    	for index, value := range scoreMap {
    		fmt.Println(index,value)
    	}
    	//-----------------------------------
    	fmt.Println("拿到key值")
    	for index := range scoreMap {
    		fmt.Println(index)
    	}
    	//-----------------------------------
    	fmt.Println("拿到value值")
    	for _, value := range scoreMap {
    		fmt.Println(value)
    	}
    ```

4.  按key的顺序遍历

    ```go
    	scoreMap := make(map[string]int,50)
    	for i := 1; i <= 50; i++ {
    		key:= fmt.Sprintf("stu%02d",i)
    		value := rand.Intn(100)
    		scoreMap[key]=value
    	}
    	fmt.Println(scoreMap)
    	//构建可以排序的切片
    	var orderSlice = make([]string,0,100)
    	for i := range scoreMap {
    		orderSlice=append(orderSlice,i)
    	}
    	//切片排序
    	sort.Strings(orderSlice)
    	//根据切片输出for循环的值
    	for i := 0; i < len(orderSlice)-1; i++ {
    		fmt.Println(orderSlice[i],scoreMap[orderSlice[i]])
    	}
    ```

5.  元素为map类型的切片

    ```go
    	Mapslice := make([]map[string]string,10)
    	//Map1 := make(map[string]string,3)
    	//Map1["name"] = "thw"
    	//Mapslice =append(Mapslice,Map1)
    	//fmt.Println(Mapslice[0]["name"])
    	Mapslice[0] =make(map[string]string,5)
    	Mapslice[0]["name"]="beijing"
    	fmt.Println(Mapslice[0])
    ```

6.  元素为切片类型的map

    ```go
    	var sliceMap = make(map[string][]string, 3)
    	fmt.Println(sliceMap)
    	fmt.Println("after init")
    	key := "中国"
    	value, ok := sliceMap[key]
    	if !ok {
    		value = make([]string, 0, 2)
    	}
    	value = append(value, "北京", "上海")
    	sliceMap[key] = value
    	fmt.Println(sliceMap)
    ```

    

