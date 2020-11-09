package link_list

// type ListNode struct{
// 	Val int
// 	Next *ListNode
// }

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tail *ListNode

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val < l2.Val {
		head, tail, l1 = l1, l1, l1.Next
	} else {
		head, tail, l2 = l2, l2, l2.Next
	}

	//循环,直到某一个链表已遍历完
	for l1 != nil && l2 != nil {
		//找到下一个节点,添加到新链表的尾
		if l1.Val < l2.Val {
			tail.Next, l1 = l1, l1.Next
		} else {
			tail.Next, l2 = l2, l2.Next
		}
		//更新tail
		tail = tail.Next
	}

	//剩下的节点字节拼接到新链表尾部
	if l1 != nil {
		tail.Next = l1
	}
	if l2 != nil {
		tail.Next = l2
	}

	return head
}
