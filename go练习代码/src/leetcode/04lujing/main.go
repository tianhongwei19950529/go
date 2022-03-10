package main

import "fmt"

func uniquePaths(m int, n int) int {
	all := make([][]int,m)
	for i := range all {
		all[i] = make([]int, n)
		all[i][0] = 1
	}
	for j := 0; j < n; j++ {
		all[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for r := 1; r < n; r++ {
			all[i][r] = all[i-1][r]+all[i][r-1]
		}
	}
	fmt.Println(all)
	return all[m-1][n-1]
}


func uniquePaths2(m int, n int) int {
	/*
	1.二维dp转一维dp 共有两个参数m  n
	2.选取一个参数n 构建一个长度为n 的切片
	2.两次循环,第一次循环另外一个参数,第二次循环第一个参数
	 */
	all := make([]int,n)
	for i := range all {
		all[i] =1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			all[j] = all[j-1]+all[j]
		}
	}
	fmt.Println(all)
	return all[n-1]
}


func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n,m := len(obstacleGrid),len(obstacleGrid[0])
	all := make([]int,m)

	if obstacleGrid[0][0]==0{
			all[0]=1
		}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if obstacleGrid[i][j]==1{
				all[j]=0
				continue
			}
			if j-1>0 && obstacleGrid[i][j-1]==0{
				all[j] = all[j-1]+all[j]
			}
		}
	}
	return all[m-1]
}
func minimumTotal(triangle [][]int) int {
	h := len(triangle)
	all := make([][]int,h)
	for i := range all{
		all[i] = make([]int,len(triangle[i]))
	}
	fmt.Println(all)
	for i:=h-1;i>=0;i--{
		for j:=0;j<len(triangle[i]);j++{
			if i==h-1{
				all[i][j]=triangle[i][j]
			}else{
				all[i][j] = min1(all[i+1][j],all[i+1][j+1])+triangle[i][j]
			}
			fmt.Println(all)
		}
	}
	return all[0][0]
}

func minimumTotal2(triangle [][]int) int {
	h := len(triangle)
	all := make([]int,len(triangle[h-1]))
	for i := h-1; i >=0 ; i-- {
		fmt.Println(i)
		for j := 0; j < len(triangle[i]); j++ {
			if i == h-1{
				all[j] = triangle[i][j]
			}else {
				all[j] = min1(triangle[i][j]+all[j],triangle[i][j]+all[j+1])
			}
			fmt.Println(all)
		}
	}

	return all[0]
}

func min1(a,b int)int{
	if a>b{return b}else{return a}
}

func min(a,b int)int{
	if a>b{return a}else{return b}
}

func main()  {
	//fmt.Println(uniquePaths(8,3))
	//fmt.Println(uniquePaths2(8,3))
	//ob := [][]int{}
	//fmt.Println(uniquePathsWithObstacles(ob))
	all := [][]int{{2},{3,4},{6,5,7},{4,1,8,3}}
	minimumTotal2(all)
}
