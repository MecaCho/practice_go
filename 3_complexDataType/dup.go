package __complexDataType

import (
	"bufio"
	"fmt"
	"os"
)

// func main() {
// 	DupLinesInFiles()
// }

// package main
//
// import (
// "fmt"
// "runtime"
// )

func reverse(a []int) {
	length := len(a)
	for i := 0; i < length/2; i++ {
		tmp := a[i]
		a[i] = a[length-i-1]
		a[length-i-1] = tmp
	}
}

func main() {
	a := []int{1, 2}
	b := []int{3, 4}
	check := a
	a = b

	fmt.Println(a, b, check)
	// [3, 4], [3, 4], [1, 2]
	example2()
}

func example2() {
	a := []int{1, 2}
	b := []int{3, 4}
	check := a
	copy(a, b)

	// fmt.Println(&a, &b, &check)
	fmt.Println(a, b, check)
	// [3, 4], [3, 4], [3, 4]

	// addA := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	// addB := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	// addCheck := (*reflect.SliceHeader)(unsafe.Pointer(&check))
	//
	// fmt.Println(&addA.Data, &addB.Data, &addCheck.Data)
}

// func main() {
// 	a := []int{1, 2, 3}
// 	reverse(a)
// 	fmt.Println(a)
// 	// Output: [3 2 1]
// }

// func main() {
// 	runtime.GOMAXPROCS(1)
// 	int_chan := make(chan int, 1)
// 	string_chan := make(chan string, 1)
// 	int_chan <- 1
// 	string_chan <- "hello"
// 	select {
// 	case value := <-int_chan:
// 		fmt.Println(value)
// 	case value := <-string_chan:
// 		panic(value)
// 	}
// }

func DupLinesInFiles() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		return
	} else {
		for _, file := range files {
			fp, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup lines in file %s, error: %v\n", file, err)
				continue
			}
			countLines(fp, counts)
			fp.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}

}
