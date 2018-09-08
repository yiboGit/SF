package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head
	var prev *ListNode

	// 这里的思路，设置两个指针；快指针指向Next.Next，两位
	// 满指针指向 Next，一位 
	// 这样 快指针指向结尾时候，慢指针正好是正中间，将链表断开，
	// 反转前一半，slow指向后一半，两半再进行比较

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		next := slow.Next
		slow.Next = prev
		prev = slow
		slow = next
	}

	first, second := prev, slow
	if fast != nil {
		second = second.Next
	}

	for second != nil {
		if first.Val != second.Val {
			return false
		}
		first, second = first.Next, second.Next
	}
	return true
}
