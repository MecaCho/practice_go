package leetcode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	r := res
	jin := 0
	for l1 != nil || l2 != nil || jin != 0 {
		value1, value2 := 0, 0
		if l1 != nil {
			value1 = l1.Val
		}
		if l2 != nil {
			value2 = l2.Val
		}
		value := (value1 + value2 + jin) % 10
		fmt.Println(value, jin)

		node := ListNode{Val: value}
		res.Next = &node
		res = res.Next

		jin = (value1 + value2 + jin) / 10

		if l1 != nil {
			l1 = l1.Next

		}
		if l2 != nil {
			l2 = l2.Next

		}
	}
	fmt.Println(r, res)
	// for ;r != nil;r = r.Next{
	// 	fmt.Println(r.Val)
	// }
	return r

}
