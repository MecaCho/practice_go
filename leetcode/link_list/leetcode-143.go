package link_list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func ReorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	var dummy ListNode
	dummy.Next = head
	pre := head
	cur := head
	for cur.Next.Next != nil {
		cur = cur.Next
	}

	last := cur.Next
	cur.Next = nil

	tmp := pre.Next

	ReorderList(tmp)

	pre.Next = last
	last.Next = tmp

}

func reorderList2(head *ListNode) {

	if head == nil || head.Next == nil {
		return
	}

	fast := head
	slow := head

	for fast != nil && fast.Next != nil && slow != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	right := slow.Next

	slow.Next = nil

	var pre *ListNode
	cur := right
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}

	right = pre

	for head != nil && right != nil {
		tmp := head.Next
		head.Next = right

		tmp1 := right.Next
		right.Next = tmp
		right = tmp1
		head = head.Next.Next
	}
}

func ReorderList1(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	fast := head
	slow := head

	for fast != nil && fast.Next != nil && slow != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	right := slow.Next

	slow.Next = nil

	right = reverseList(right)

	for head != nil && right != nil {
		tmp := head.Next
		head.Next = right

		tmp1 := right.Next
		right.Next = tmp
		right = tmp1
		head = head.Next.Next
	}

}

// 递归解法，找到最后一个节点连接到第一个节点后面，中间部分递归处理，该方法在Python会超时
//
// /**
// * Definition for singly-linked list.
// * type ListNode struct {
// *     Val int
// *     Next *ListNode
// * }
// */
// func ReorderList(head *ListNode) {
// 	if head == nil || head.Next == nil {
// 		return
// 	}
//
// 	var dummy ListNode
// 	dummy.Next = head
// 	pre := head
// 	cur := head
// 	for cur.Next.Next != nil {
// 		cur = cur.Next
// 	}
//
// 	last := cur.Next
// 	cur.Next = nil
//
// 	tmp := pre.Next
//
// 	ReorderList(tmp)
//
// 	pre.Next = last
// 	last.Next = tmp
//
// }
// 1.快慢指针找到中间节点，
// 2.对后半部分反转，
// 3.然后拼接前后两个部分
//
// func reorderList2(head *ListNode)  {
//
// 	if head == nil || head.Next == nil {
// 		return
// 	}
//
//
//    // 1.快慢指针找到中间节点，
// 	fast := head
// 	slow := head
//
// 	for fast != nil && fast.Next != nil && slow != nil {
// 		fast = fast.Next.Next
// 		slow = slow.Next
// 	}
//
// 	right := slow.Next
//
// 	slow.Next = nil
//
//    //2.对后半部分反转
// 	var pre *ListNode
// 	cur := right
// 	for cur != nil {
// 		tmp := cur.Next
// 		cur.Next = pre
// 		pre = cur
// 		cur = tmp
// 	}
//
// 	right = pre
//    // 3.然后拼接前后两个部分
// 	for head != nil && right != nil {
// 		tmp := head.Next
// 		head.Next = right
//
// 		tmp1 := right.Next
// 		right.Next = tmp
// 		right = tmp1
// 		head = head.Next.Next
// 	}
// }
//
// 作者：qiuwenqi
// 链接：https://leetcode-cn.com/problems/reorder-list/solution/golangliang-chong-fang-fa-di-gui-he-die-dai-by-qiu/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
