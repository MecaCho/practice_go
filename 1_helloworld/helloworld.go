package __helloworld

import (
	"bufio"
	"fmt"
	"os"
)

// import "fmt"
//
// func main()  {
// 	fmt.Println("Hello, World!")
// 	fmt.Println("你好, 世界!")
// }

func DupFunc() {

	// fmt.Println(strings.Join(os.Args[1:], " "))
	//
	// var s, sep string
	//
	// for i := 1; i < len(os.Args); i++ {
	// 	s += sep + os.Args[i]
	// 	sep = ""
	// }
	// fmt.Println(s)

	Dup()
}

func Dup() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	// res := input.Scan()
	// fmt.Println(res, input.Text())
	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}
