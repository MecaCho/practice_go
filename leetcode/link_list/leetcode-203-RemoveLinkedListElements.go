package link_list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {

	pre := &ListNode{}
	var dummy *ListNode

	dummy = head
	pre.Next = head
	for head != nil {
		for head != nil && head.Val == val {
			pre.Next = head.Next
			head = head.Next
		}
		if head != nil {
			head = head.Next
			pre = pre.Next
		}

	}

	for dummy != nil && dummy.Val == val {
		dummy = dummy.Next
	}

	return dummy.Next
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements1(head *ListNode, val int) *ListNode {

	dummy := new(ListNode)
	dummy.Next = head
	pre := dummy

	for head != nil {
		if head.Val != val {
			head = head.Next
			pre = pre.Next
		} else {
			pre.Next = head.Next
			head = head.Next
		}
	}

	return dummy.Next
}
