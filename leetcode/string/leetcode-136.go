package string

import (
	"fmt"
	"sort"
)

func WordBreak(s string, wordDict []string) bool {

	if len(s) < 1 {
		return false
	}

	sort.Slice(wordDict, func(i, j int) bool {

		return len(wordDict[i]) > len(wordDict[j])

	})

	fmt.Println(wordDict)

	dp := []int{}

	for i := range s {

		dp = append(dp, 0)

		// start := 0
		for _, word := range wordDict {
			wordLength := len(word)
			if i >= wordLength-1 {
				fmt.Println(i, s[i-wordLength+1:i+1], word)
			}
			if s[:i+1] == word || (i >= wordLength && s[i-wordLength+1:i+1] == word && dp[i-wordLength] == 1) {
				dp[i] = 1
				break
			}

		}
	}

	fmt.Println(dp)

	return dp[len(s)-1] == 1
}
