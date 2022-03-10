package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	vis := map[*ListNode]bool{}
	for tmp := headA; tmp != nil; tmp = tmp.Next {
		vis[tmp] = true
	}
	fmt.Println(vis)
	for tmp := headB; tmp != nil; tmp = tmp.Next {
		if vis[tmp] {
			return tmp
		}
	}
	return nil
}

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headB == nil || headA == nil{
		return nil
	}
	pa,pb := headA,headB
	for pa != pb{
		if pa == nil{
			pa = headB
		}else {
			pa=pa.Next
		}
		if pb == nil{
			pb = headA
		}else {
			pb =pb.Next
		}

	}
	return pa
}


func main() {
	//intersectVal = 8, listA = [4,1,8,4,5], listB = [5,6,1,8,4,5], skipA = 2, skipB = 3
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

	list6 := ListNode{
		Val:  6,
		Next: &list1,
	}

	list5_1 := ListNode{
		Val:  5,
		Next: &list6,
	}

	//headA := &[]int{4,1,8,4,5}
	//headB := &[]int{5,6,1,8,4,5}
	fmt.Println(getIntersectionNode2(&list4_1, &list5_1))

}
