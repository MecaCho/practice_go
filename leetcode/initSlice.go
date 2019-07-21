package leetcode

import (
	"fmt"
	"os/exec"
	"time"
)

func InitSlice()  {

	var a []int
	fmt.Println("Slice init : ", a, len(a), cap(a))

	s := make([]int, 0)
	fmt.Println(s, len(s), cap(s))

	s1 := make([]int, 5, 10)
	s2 := []int{1, 2, 3, 4, 5}
	//s3 := make([]int)

	fmt.Println(s, s1, s2)
	exec.Command("/mnt/get_blog.sh")
	time.Now()

}

