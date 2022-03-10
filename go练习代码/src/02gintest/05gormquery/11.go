package main

import "fmt"

func twoSum(nums []int, target int) []int {
	sumMap :=  make(map[int]int)
	for i, num := range nums {
		fmt.Println(sumMap)
		index,ok := sumMap[target-num]
		if ok{
			return []int{index,i}
		}
		sumMap[num]=i
	}
	return nil
}

func lengthOfLastWord(s string) int {
	ans := 0
	lenS := len(s) -1
	for s[lenS] == ' '{
		lenS -=1
	}
	for lenS>=0 && s[lenS]!= ' '{
		ans+=1
		lenS -=1
	}
	return ans
}


func climbStairs(n int) int {
	if n <=2{
		return n
	}else{
		countList := make([]int,n,n)
		countList[0]=1
		countList[1]=2
		for i:=2;i<n;i++{
			countList[i]=countList[i-1]+countList[i-2]
		}
		return countList[n-1]
	}
}

func main()  {
	//nums := []int{2,7,11,15}
	//target := 9
	//fmt.Println(twoSum(nums,target))
	fmt.Println(climbStairs(4))
}