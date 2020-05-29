package __baseDataType

import (
	"fmt"
	"os"
)

func DeferReturn() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("fatal")
		}
	}()

	defer func() {
		panic("defer panic")
	}()
	panic("panic")
}

func Deferstack() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func CalDefer() {

	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1

}

func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}

func DeferFun() {

	println(DeferFunc1(1))
	println(DeferFunc2(1))
	println(DeferFunc3(1))

}

func VarInDefer() {

	var i = 1
	defer fmt.Println("result: ", func() int { return i * 2 }())
	i++

}

func DeferInFor() {
	files := []string{"a", "b", "c"}

	for k := range files {
		f, _ := os.Open(files[k])
		defer f.Close()
	}

	for k := range files {
		func() {
			f, _ := os.Open(files[k])
			defer f.Close()
		}()
	}

}
