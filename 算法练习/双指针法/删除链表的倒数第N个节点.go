// author:zfy  date:2022/8/23

package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{}
	dummyHead.Next = head
	cur := head
	prev := dummyHead
	i := 1
	for cur != nil {
		cur = cur.Next
		if i > n {
			prev = prev.Next
		}
		i++
	}
	prev.Next = prev.Next.Next
	return dummyHead.Next
}

// 双指针法

func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	dummyNode := &ListNode{0, head}
	fast, slow := dummyNode, dummyNode
	for i := 0; i <= n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummyNode.Next
}
