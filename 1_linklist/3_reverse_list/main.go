package main

func main() {
	a := ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				3,
				&ListNode{
					Val: 4,
				},
			},
		},
	}
	reverseList(&a)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		// 第一个head指向原head的下一个节点(成为新的head)，
		// 第二个head.Next 是原 head.Next的指针 指向上一个节点
		// 第三个 prev 是 前一个节点指向 原head
		// 这三步是同时进行
		head, head.Next, prev = head.Next, prev, head
	}
	return prev
}
