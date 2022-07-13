// author:zfy  date:2022/7/13

package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	cur := head
	per := dummy
	i := 1
	for cur != nil {
		cur = cur.Next
		if i > n {
			per = per.Next
		}
		i++
	}
	per.Next = per.Next.Next
	return dummy.Next
}
