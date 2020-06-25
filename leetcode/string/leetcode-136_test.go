package string

import (
	"fmt"
	"testing"
)

func TestWordBreak(t *testing.T) {
	s := "applepenapple"
	wordDict := []string{"leet", "code", "apple", "pen"}

	res := WordBreak(s, wordDict)
	fmt.Println(res)
}
