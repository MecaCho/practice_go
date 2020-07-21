package CrackinfInterview

import (
	"fmt"
	"math"
)

func getKthMagicNumber(k int) int {
	aNum, bNum, cNum := 0, 0, 0
	res := []int{1}

	for i := 0; i < k; i++ {
		min_num := min(res[aNum]*3, res[bNum]*5, res[cNum]*7)

		if min_num == res[aNum]*3 {
			aNum++
		}
		if min_num == res[bNum]*5 {
			bNum++
		}
		if min_num == res[cNum]*7 {
			cNum++
		}
		res = append(res, min_num)
	}

	return res[k-1]
}

func GetKthMagicNumber1(k int) int {
	aNum, bNum, cNum := 0, 0, 0
	res := []int{1}
	// kk := k - 1
	for i := 0; i < k; i++ {
		min_num := min(res[aNum]*3, res[bNum]*5, res[cNum]*7)
		// fmt.Println(i, num1, num2, num3)
		// nums := []int{res[aNum]*3, res[bNum]*5, res[cNum]*7}
		// fmt.Println(nums)
		// sort.Ints(nums)
		// min_num := nums[0]
		// fmt.Println(nums)
		// fmt.Printf("%d %d %d   ", aNum, bNum, cNum)
		// aNum_, bNum_, cNum_ := aNum, bNum, cNum
		// switch {
		// case min_num == res[aNum]*3:
		// 	aNum++
		// case min_num == res[bNum]*5:
		// 	bNum++
		// case min_num == res[cNum]*7:
		// 	cNum++
		// }
		// fmt.Println(aNum, bNum, cNum, min_num)
		if min_num == res[aNum]*3 {
			aNum++
		}
		if min_num == res[bNum]*5 {
			bNum++
		}
		if min_num == res[cNum]*7 {
			cNum++
		}
		res = append(res, min_num)
	}
	fmt.Println(res)
	return res[k-1]

}

func min(a, b, c int) int {
	return int(math.Min(float64(a), math.Min(float64(b), float64(c))))
}

func Switch() {

	a, b := 1, 2

	switch {
	case 1 == a:
		fmt.Println("1")
		fallthrough
	case 4 == b:
		fmt.Println("2")

	}
}
