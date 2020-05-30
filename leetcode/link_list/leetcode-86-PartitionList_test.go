package link_list

import (
	"fmt"
	"testing"
)

func TestPartition(t *testing.T) {

	h := &ListNode{1, &ListNode{4, &ListNode{3,
		&ListNode{2, &ListNode{5, &ListNode{2, nil}}}}}}

	new_h := Partition(h, 3)

	for new_h != nil {
		fmt.Println(new_h.Val)
		new_h = new_h.Next
	}
}
