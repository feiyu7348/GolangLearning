// author:zfy  date:2022/7/10

package main

func reverseList(head *ListNode) *ListNode {
	var per *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = per
		per = cur
		cur = next
	}
	return per
}
