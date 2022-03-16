package main

import (
	"fmt"
	"math/rand"
)

// 判断两个二叉树是否相等

//				3
//		1				8
//	1		2		5		13
//
////
//							8
//				3						13
//		1				5
//	1		2

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func New(k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(1000000) {
		t = insert(t, (1+v)*k)
	}
	return t
}
func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}

// Walk 遍历树 t
// 将所有的值从 t 发送到 ch
func Walk(t *Tree, ch chan int) {
	if t == nil {
		return
	}
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same 检测树 t1 和 t2 是否表示相同的序列
func Same(t1, t2 *Tree) bool {
	ch1 := make(chan int)
	go func() {
		defer close(ch1)
		Walk(t1, ch1)
	}()
	ch2 := make(chan int)
	go func() {
		defer close(ch2)
		Walk(t2, ch2)
	}()
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if ok1 != ok2 || v1 != v2 {
			return false
		}
		if !ok1 {
			break
		}
	}
	return true
}

func main() {
	fmt.Println(rand.Perm(100))
	fmt.Println(Same(New(1), New(1)))
	fmt.Println(Same(New(2), New(1)))
}
