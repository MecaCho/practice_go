package __base

import "fmt"

func LoopGoTo() {
	//  放在for前面，此例会一直循环下去
Loop:
	fmt.Println("test")
	for a := 0; a < 5; a++ {
		fmt.Println(a)
		if a > 3 {
			goto Loop
		}
	}
}

func GoToLoop() {
	for a := 0; a < 5; a++ {
		fmt.Println(a)
		if a > 3 {
			goto Loop
		}
	}
Loop: //放在for后边
	fmt.Println("test, loop after for")
}

func BreakLoop() {
Loop:
	for j := 0; j < 3; j++ {
		fmt.Println(j)
		for a := 0; a < 5; a++ {
			fmt.Println(a)
			if a > 3 {
				break Loop
			}
		}
	}
}
