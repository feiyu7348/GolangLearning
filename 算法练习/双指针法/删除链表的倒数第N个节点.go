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
