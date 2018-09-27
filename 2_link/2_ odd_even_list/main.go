package main

func main() {

}

type ListNode struct {
	Val  string
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	odd := head
	evenHead := head.Next
	even := evenHead

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}

	// 奇数链 -> 链接 -> 偶数链
	odd.Next = evenHead
	return head
}
