package leetcode

import "fmt"

func ff1() {
	defer println("f1-begin")
	ff2()
	defer println("f1-end")
}

func ff2() {
	defer println("f2-begin")
	ff3()
	defer println("f2-end")
}

func ff3() {
	defer println("f3-begin")
	panic(0)
	defer println("f3-end")
}

func PrintPnic() {
	//f3-begin
	//f2-begin
	//f1-begin
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	ff1()
}
