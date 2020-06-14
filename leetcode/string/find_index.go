package string

import (
	"fmt"
	"strings"
)

func FindIndex(str, subStr string) int {
	if len(subStr) > len(str) {
		return -1
	}
	for i := range str {
		if len(str)-i < len(subStr) {
			return -1
		}
		if subStr[0] == str[i] {
			j := 1
			for j < len(subStr) {
				if subStr[j] != str[i+j] {
					break
				}
				j++
			}
			if j == len(subStr) {
				return i
			}
		}
	}
	return -1
}

func main() {

	strings.Index("abc", "abc")

	res := FindIndex("abcdefg", "bcf")
	fmt.Println(res)

}
