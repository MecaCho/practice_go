package slide_window

// 567. Permutation in String
// Given two strings s1 and s2, write a function to return true if s2 contains the permutation of s1. In other words, one of the first string's permutations is the substring of the second string.
//
//
//
// Example 1:
//
// Input: s1 = "ab" s2 = "eidbaooo"
// Output: True
// Explanation: s2 contains one permutation of s1 ("ba").
// Example 2:
//
// Input:s1= "ab" s2 = "eidboaoo"
// Output: False
//
//
// Constraints:
//
// The input strings only contain lower case letters.
// The length of both given strings is in range [1, 10,000].

// 567. 字符串的排列
// 给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。
//
// 换句话说，第一个字符串的排列之一是第二个字符串的子串。
//
// 示例1:
//
// 输入: s1 = "ab" s2 = "eidbaooo"
// 输出: True
// 解释: s2 包含 s1 的排列之一 ("ba").
//
//
// 示例2:
//
// 输入: s1= "ab" s2 = "eidboaoo"
// 输出: False
//
//
// 注意：
//
// 输入的字符串只包含小写字母
// 两个字符串的长度都在 [1, 10,000] 之间

func checkInclusion1(s1, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}
	var cnt1, cnt2 [26]int
	for i, ch := range s1 {
		cnt1[ch-'a']++
		cnt2[s2[i]-'a']++
	}
	if cnt1 == cnt2 {
		return true
	}
	for i := n; i < m; i++ {
		cnt2[s2[i]-'a']++
		cnt2[s2[i-n]-'a']--
		if cnt1 == cnt2 {
			return true
		}
	}
	return false
}

func checkInclusion(s1 string, s2 string) bool {
	counter1, counter2 := [26]int{}, [26]int{}
	i := 0
	length1, length2 := len(s1), len(s2)
	if length1 > length2 {
		return false
	}

	for i, ch := range s1 {
		counter1[ch-'a'] += 1
		counter2[s2[i]-'a'] += 1
	}

	if counter2 == counter1 {
		return true
	}

	for i = length1; i < length2; i++ {
		counter2[s2[i]-'a'] += 1
		counter2[s2[i-length1]-'a'] -= 1
		if counter2 == counter1 {
			return true
		}
	}

	return false
}

func CheckInclusion(s1 string, s2 string) bool {
	return checkInclusion(s1, s2)
}
