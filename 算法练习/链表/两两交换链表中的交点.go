// author:zfy  date:2022/7/11

package main

func swapParis(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	per := dummy
	for head != nil && head.Next != nil {
		per.Next = head.Next
		next := head.Next.Next
		head.Next.Next = head
		head.Next = next
		per = head
		head = next
	}
	return dummy.Next
}
