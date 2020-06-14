package __base

import "fmt"

// func NewBool() {
//
// 	a = bool(1 == 1)
// 	// a = bool(1)
// 	fmt.Println(bool(1 == 1))
//
// }

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseK(head *ListNode, k int) (*ListNode, *ListNode) {

	t := head
	tmp, pre, cur := &ListNode{}, &ListNode{}, head
	for k := k; k >= 0; k-- {
		tmp = cur.Next
		pre, cur.Next = cur, pre
		cur = tmp
	}
	return pre, t
}

// # Definition for singly-linked list.
// # class ListNode(object):
// #     def __init__(self, x):
// #         self.val = x
// #         self.next = None
//
// class Solution(object):
//    def reverseBetween(self, head, m, n):
//        """
//        :type head: ListNode
//        :type m: int
//        :type n: int
//        :rtype: ListNode
//        """
//        dummmy = ListNode(0)
//        dummmy.next = head
//        cur, pre = head, dummmy
//        while m>1:
//            cur = cur.next
//            pre = pre.next
//            m -= 1
//            n -= 1
//
//        h, t = pre, cur
//
//        while n:
//            tmp = cur.next
//            cur.next = pre
//            pre = cur
//            cur = tmp
//            n -= 1
//
//        h.next = pre
//        t.next = cur
//        return dummmy.next
//
//
func ReverseBetween1(head *ListNode, m int, n int) *ListNode {

	dummy := &ListNode{}
	dummy.Next = head
	pre := dummy
	cur := head

	for i := 0; i < m-1; i++ {
		pre = pre.Next
		cur = cur.Next
	}

	h, t := pre, cur

	for i := 0; i < n-m; i++ {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}

	h.Next = pre
	t.Next = cur

	fmt.Println(pre, cur, h, t)

	return dummy.Next
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//
func ReverseBetween(head *ListNode, m int, n int) *ListNode {

	dummy := &ListNode{}
	dummy.Next = head
	pre := dummy
	cur := head

	for i := 0; i < m-1; i++ {
		pre = pre.Next
		cur = cur.Next
	}
	fmt.Println(pre, cur)

	h, t := pre, cur

	for i := 0; i <= n-m; i++ {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}

	h.Next = pre
	t.Next = cur

	fmt.Println(pre, cur, h, t)

	return dummy.Next

}

func main() {

}

func MySort(arr []string) []string {
	length := len(arr)
	if length == 0 || length == 1 {
		return arr
	}
	len0 := len(arr[0])
	left := []string{}
	right := []string{}
	j := 1
	for j < len(arr) {
		// fmt.Println(arr[j])
		if len(arr[j]) > len0 {
			right = append(right, arr[j])
		} else {
			left = append(left, arr[j])
		}
		j++
	}

	// fmt.Println(left, right)

	sortedLeft := MySort(left)
	sortedLeft = append(sortedLeft, arr[0])
	sortedRight := MySort(right)
	for i := range sortedRight {
		sortedLeft = append(sortedLeft, sortedRight[i])
	}
	return sortedLeft
}
