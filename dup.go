package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	DupLinesInFiles()
}

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
