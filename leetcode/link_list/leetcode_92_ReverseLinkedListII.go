package link_list

// 92. 反转链表 II
// 反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
//
// 说明:
// 1 ≤ m ≤ n ≤ 链表长度。
//
// 示例:
//
// 输入: 1->2->3->4->5->NULL, m = 2, n = 4
// 输出: 1->4->3->2->5->NULL

// 92. Reverse Linked List II
// Reverse a linked list from position m to n. Do it in one-pass.
//
// Note: 1 ≤ m ≤ n ≤ length of list.
//
// Example:
//
// Input: 1->2->3->4->5->NULL, m = 2, n = 4
// Output: 1->4->3->2->5->NULL

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//
func reverseBetween(head *ListNode, m int, n int) *ListNode {

	dummy := &ListNode{}
	dummy.Next = head
	pre := dummy
	cur := head

	for i := 0; i < m-1; i++ {
		pre = pre.Next
		cur = cur.Next
	}

	h, t := pre, cur

	for i := 0; i <= n-m; i++ {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}

	h.Next = pre
	t.Next = cur

	return dummy.Next

}

// 作者：qiuwenqi
// 链接：https://leetcode-cn.com/problems/reverse-linked-list-ii/solution/goyu-yan-shuang-100-by-qiuwenqi/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
