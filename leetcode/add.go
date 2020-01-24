package leetcode

import "fmt"

func add(args ...int) int {
	sum := 0
	for _, arg := range args {
		sum += arg
	}
	fmt.Println("add int : ", args, ", sum : ", sum)
	return sum
}
func AddTest() {

	add(3, 4)
	add(5, 6, 7)
	//add([]int{3, 4})
	add([]int{5, 6, 7}...)

}
