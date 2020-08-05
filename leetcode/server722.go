package leetcode

import (
	"fmt"
)

func MergeArrr(arr1, arr2 []int) []int {
	arr3 := []int{}
	i, j := len(arr1)-1, len(arr2)-1
	for i >= 0 && j >= 0 {
		if arr2[j] > arr1[i] {
			arr3 = append(arr3, arr2[j])
			j--
		} else {
			arr3 = append(arr3, arr1[i])
			i--
		}
	}
	for i >= 0 {
		arr3 = append(arr3, arr1[i])
		i--
	}

	for j >= 0 {
		arr3 = append(arr3, arr2[j])
		j--
	}

	fmt.Println(arr3)
	return arr3

}

func FiveGroutinue() {
	// //  100个任务
	// //  5个线程
	chans := make(chan int, 5)
	//tasks := []int{}

	done := make(chan int)

	fmt.Println("hello")
	go func() {
		for i := 0; i < 100; i++ {
			//tasks = append(tasks, i)
			chans <- i
		}
	}()

	go func() {
		for {
			res := <-chans
			fmt.Println(res)
			if res >= 99 {
				done <- 1
			}

		}

	}()

	<-done

	// time.Sleep(1*time.Second)

	MergeArrr([]int{1, 2, 3, 8, 9}, []int{1, 2, 3, 4, 5, 6, 7})

}
