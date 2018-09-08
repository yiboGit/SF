package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = fast.Next
		if fast == slow {
			return true
		}
	}
	return false
}

// java 代码思路
// public boolean hasCycle(ListNode head) {

// 	ListNode fast = head, slow = head;
// 	while (fast != null && fast.next != null) {
// 		fast = fast.next.next;
// 		slow = slow.next;
// 		if (slow == fast) {
// 			return true;
// 		}
// 	}
// 	return false;
// }
