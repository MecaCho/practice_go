package __complexDataType

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func DupLinesInFiles1() {

	counts := make(map[string]int)
	for _, file := range os.Args[1:] {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup lines in file %s, error: %v\n", file, err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func main() {

	DupLinesInFiles1()

}
