package __base

import "fmt"

func BreakLoop1() {
loop:
	for {
		switch {
		case true:
			fmt.Println("breaking out...")
			// break    // 死循环，一直打印 breaking out...
			break loop
		}
	}
	fmt.Println("out...")
}
