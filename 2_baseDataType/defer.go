package __baseDataType

import (
	"fmt"
	"os"
	"time"
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
	return
	// 	=== RUN   TestDeferstack2
	// 打印后
	// 打印中
	// 打印前
	// --- FAIL: TestDeferstack2 (0.00s)
	// panic: 触发异常 [recovered]
	//	panic: 触发异常
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
		fmt.Println("t in defer1: ", t)
	}()
	t += 3
	return t
	// === RUN   TestDeferFunc1
	// t in defer1:  7
	// 7

	// t in defer1:  4
	// t in func result: 4
	// t in defer2:  4
	// t in func result: 1
	// t in defer3:  3
	// t in func result: 3
}

func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
		fmt.Println("t in defer2: ", t)
	}()
	return t
}

func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
		fmt.Println("t in defer3: ", t)
	}()
	return 2
}

func DeferFun() {

	println("t in func result:", DeferFunc1(1))
	println("t in func result:", DeferFunc2(1))
	println("t in func result:", DeferFunc3(1))
	// 	4
	// 1
	// 3
	// t in defer1:  4
	// t in func result: 4
	// t in defer2:  4
	// t in func result: 1
	// t in defer3:  3
	// t in func result: 3

}

func VarInDefer() {

	var i = 1
	defer func() {
		fmt.Println("result: ", func() int { return i * 2 }())
	}()
	i++

	startedAt := time.Now()
	defer func() { fmt.Println("defer time: ", time.Since(startedAt)) }()

	time.Sleep(time.Second)

	fmt.Println("return time: ", time.Since(startedAt))
	// 	return time:  1.003767335s
	// defer time:  1.003916954s

}

func VarInDefer1() {

	var i = 1
	defer fmt.Println("result: ", func() int { return i * 2 }())
	i++

	startedAt := time.Now()
	defer fmt.Println("defer time: ", time.Since(startedAt))

	time.Sleep(time.Second)
	fmt.Println("return time: ", time.Since(startedAt))

	// 	return time:  1.002410955s
	// defer time:  205ns
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
