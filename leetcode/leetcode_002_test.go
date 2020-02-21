package leetcode

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	list1 := ListNode{Val: 1}
	list2 := ListNode{Val: 9}
	res := AddTwoNumbers(&list1, &list2)
	fmt.Printf("result: %+v.", res)
}
