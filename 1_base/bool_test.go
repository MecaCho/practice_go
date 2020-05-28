package __base

import (
	"fmt"
	"testing"
)

func TestReverseBetween(t *testing.T) {
	l := ListNode{0, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}}
	res := ReverseBetween(&l, 2, 4)
	// fmt.Println(res)
	fmt.Println(res.Val)
	for res.Next != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
	fmt.Printf("%+v", *res)

	// h, tail := ReverseK(&l, 2)
	// fmt.Println(h, "xxxxxxxx", tail)
}
