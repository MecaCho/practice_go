package leetcode

import "fmt"

func PrintFunc(n int) {

	for i := 0; i < n; i++ {
		fmt.Printf("#  ")
		PrintFunc(n - 1)
	}

}

func RunPrint() {

	// 3 * (1 + 2 * (1 + 1)) = 15

	PrintFunc(3)

}
