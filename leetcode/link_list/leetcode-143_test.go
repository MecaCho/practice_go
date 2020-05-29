package link_list

import (
	"fmt"
	"testing"
)

func TestReorderList(t *testing.T) {
	h := &ListNode{0, &ListNode{1, &ListNode{2,
		&ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}}

	// ReorderList(h)
	ReorderList1(h)

	for h != nil {
		fmt.Println(h.Val)
		h = h.Next
	}
}

func TestReorderList1(t *testing.T) {
	h := &ListNode{0, &ListNode{1, &ListNode{2,
		&ListNode{3, nil}}}}

	// ReorderList(h)
	ReorderList1(h)

	for h != nil {
		fmt.Println(h.Val)
		h = h.Next
	}
}
