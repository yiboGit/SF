package main

import "fmt"

func main() {
	a := ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				Val: 4,
			},
		},
	}
	b := ListNode{
		1,
		&ListNode{
			3,
			&ListNode{
				Val: 4,
			},
		},
	}

	result := mergeTwoLists(&a, &b)
	for result.Next != nil {
		fmt.Printf("%v", result.Val)
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	cur := result
	for l1 != nil || l2 != nil {
		if l1 == nil {
			cur.Next = l2
			return result.Next
		}
		if l2 == nil {
			cur.Next = l1
			return result.Next
		}
		// 不能这样写 因为此时 cur 是一个新的链表 cur.Next什么都没有
		// 这里的 cur= cur.Next 想得到的是拼接之后的下一个节点，作为下一次循环的当前节点
		// if l1.Val > l2.Val {
		// 	cur.Next, cur, l2 = l2, cur.Next, l2.Next
		// } else {
		// 	cur.Next, cur, l1 = l1, cur.Next, l1.Next
		// }
		if l1.Val > l2.Val {
			cur.Next = l2
			cur, l2 = cur.Next, l2.Next
		} else {
			cur.Next = l1
			cur, l1 = cur.Next, l1.Next
		}
	}

	return result.Next
}
