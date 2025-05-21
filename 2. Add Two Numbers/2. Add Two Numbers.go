package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var c int
	res := &ListNode{}
	res.Val = l1.Val + l2.Val
	if res.Val >= 10 {
		res.Val -= 10
		c = 1
	}
	l1 = l1.Next
	l2 = l2.Next

	prev := res

	for l1 != nil || l2 != nil || c != 0 {
		var a, b int
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}

		cur := &ListNode{}
		cur.Val = a + b + c
		c = 0
		if cur.Val >= 10 {
			cur.Val -= 10
			c = 1
		}
		prev.Next = cur
		prev = cur
	}

	return res
}
