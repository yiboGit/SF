package main

func main() {

}

// 给定两个非空链表来表示两个非负整数。位数按照逆序方式存储，它们的每个节点只存储单个数字。将两数相加返回一个新的链表。
// 你可以假设除了数字 0 之外，这两个数字都不会以零开头。
// 示例：
// 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
// 输出：7 -> 0 -> 8
// 原因：342 + 465 = 807

type ListNode struct {
	Val  int
	Next *ListNode
}

// 第二种思路使用循环 迭代
func addTwoNumbersTwo(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	res := result
	var carryInt int

	for l2 != nil || l1 != nil || carryInt != 0 {
		sum := carryInt
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carryInt = sum / 10
		//这里需要做判断 才能对下一个节点进行初始化创建
		//判断的是循环条件，如不判断，会建立一个空链表节点，
		//例如预期是 7，0，8  会输出7，0，8，0 ；
		//这里解决有两个思路，一个是在此处进行判断,通过判断 建立初始化Next节点，一个是直接在此处对next操作赋值，最终结果返回的是result.Next()
		// 1-----------  
		res.Val = sum % 10
		if l2 != nil || l1 != nil || carryInt != 0 {
			res.Next = &ListNode{}
			res = res.Next
		}
		// 2-----------***
		// res.Next = &ListNode{Val: sum % 10}
		// res = res.Next
	}
	return result
}

// 第一种思路 使用递归的方法 进行计算
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addNum(l1, l2, 0)
}

func addNum(l1 *ListNode, l2 *ListNode, carryInt int) *ListNode {
	res := ListNode{}

	sum := carryInt

	if l1 != nil {
		sum += l1.Val
		l1 = l1.Next
	}

	if l2 != nil {
		sum += l2.Val
		l2 = l2.Next
	}

	carryInt = sum / 10
	res.Val = sum % 10

	if l1 != nil || l2 != nil || carryInt != 0 {
		res.Next = addNum(l1, l2, carryInt)
	}
	return &res
}
