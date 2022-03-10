package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode

}
func reverseList(head *ListNode) *ListNode {
	/*
	1.定义一个当前节点的前节点
	2.当前节点不为空
	3.取出当前节点的后节点
	4.将当前节点的后节点给到前节点
	5.
	 */
	var prev *ListNode
	curr := head
	for curr != nil {
		//next := curr.Next
		//curr.Next = prev
		//prev = curr
		//curr = next
		prev, curr, curr.Next = curr, curr.Next, prev
	}
	return prev
}

func main()  {
	list5 := ListNode{
		Val:  5,
		Next: nil,
	}
	list4 := ListNode{
		Val:  4,
		Next: &list5,
	}
	list8 := ListNode{
		Val:  8,
		Next: &list4,
	}
	list1 := ListNode{
		Val:  1,
		Next: &list8,
	}
	list4_1 := ListNode{
		Val:  4,
		Next: &list1,
	}
	fmt.Println(reverseList(&list4_1))
}
