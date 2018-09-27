package main

import "fmt"

func main() {
	link := ListNode{
		Val: 1,
	}
	removeNthFromEnd(&link, 1)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 会有两种特殊情况，1.链表为空；2.链表只有一个元素
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 先考虑特殊情况
	if head == nil {
		return head
	}
	var (
		pre    *ListNode
		remove = head
	)
	for l := head; l != nil; l = l.Next {
		if n == 0 {
			pre = remove
			remove = remove.Next
		} else {
			n--
		}
	}
	if pre == nil {
		// 如果链表只有一个元素，第一次循环，
		// 就跳出了，故pre没有被赋值
		fmt.Printf("pre: %v \n", pre)
		head = head.Next
	} else {
		pre.Next = remove.Next
	}
	return head
}
