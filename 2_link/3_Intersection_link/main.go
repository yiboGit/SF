package main

import "fmt"

func main() {
	c := &ListNode{
		"c1",
		&ListNode{
			"c2",
			&ListNode{},
		},
	}
	one := ListNode{
		"a1",
		&ListNode{
			"c1",
			c,
		},
	}
	two := ListNode{
		"b1",
		&ListNode{
			"b2",
			&ListNode{
				"b3",
				c,
			},
		},
	}

	fmt.Printf("%v \n", getIntersectionNode(&one, &two))
}

type ListNode struct {
	Val  string
	Next *ListNode
}

// 算法思路非常简单， 两个链表循环，
// 如果循环完，为空 那么给 另一条 链表重新循环，
// 两个链表一起循环 比较节点是否是相交
// 如果没有相交部分，最终会一同结束，返回nil
// 如下： A : 1 -> 4 -> 8 -> 6 -> 3  B : 2 -> 6 -> 3
//  1 -> 4 -> 8 -> 6 -> 3 || -> 2 -> 6 -> 3
//  2 -> 6 -> 3 || -> 1 -> 4 -> 8 -> 6 -> 3
// 一目了然 毕竟 m+n == n + m 啊
func getIntersectionNode(head1, head2 *ListNode) *ListNode {
	cur1 := head1
	cur2 := head2
	for cur1 != cur2 {
		if cur1 != nil {
			cur1 = cur1.Next
		} else {
			cur1 = head2
		}
		if cur2 != nil {
			cur2 = cur2.Next
		} else {
			cur2 = head1
		}
	}
	return cur1
}
