package link_list

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func Partition(head *ListNode, x int) *ListNode {

	pre := head
	cur := head

	for cur != nil {

		if cur.Next != nil && cur.Next.Val < x && cur.Next.Val < pre.Next.Val {
			fmt.Println(pre.Val, cur.Val)
			tmp := pre.Next
			pre.Next = cur.Next
			tmp1 := cur.Next.Next
			cur.Next.Next = tmp

			cur.Next = tmp1

			fmt.Println(tmp, tmp1)

			cur = tmp1
			pre = pre.Next

		} else {
			cur = cur.Next
		}

	}

	return head

}
