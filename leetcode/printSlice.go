package leetcode

import "fmt"

func PrintSlice() {
	s1 := []int{1, 2, 3}
	s2 := make([]int, 3)

	s2 = append(s2, s1...)
	// s2 = 0, 0, 0, 1, 2, 3,

	s3 := s1[1:]
	s3 = append(s3, 4)
	// s3 = 2, 3, 4

	for _, v := range s2 {
		v += 10
	}

	for i := range s2 {
		s2[i] += 10
	}
	// 10, 10, 10, 11, 12, 13
	fmt.Println(s2)

	for _, v := range s3 {
		v += 10
	}

	for i := range s3 {
		s3[i] += 10
	}
	// 12, 13, 14
	fmt.Println(s3)
}
