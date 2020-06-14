package hash

import (
	"sort"
	"strings"
)

// 49. 字母异位词分组
// 给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。
//
// 示例:
//
// 输入: ["eat", "tea", "tan", "ate", "nat", "bat"]
// 输出:
// [
//  ["ate","eat","tea"],
//  ["nat","tan"],
//  ["bat"]
// ]
// 说明：
//
// 所有输入均为小写字母。
// 不考虑答案输出的顺序。

// 49. Group Anagrams
// Given an array of strings, group anagrams together.
//
// Example:
//
// Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
// Output:
// [
//  ["ate","eat","tea"],
//  ["nat","tan"],
//  ["bat"]
// ]
// Note:
//
// All inputs will be in lowercase.
// The order of your output does not matter.

func groupAnagrams(strs []string) [][]string {
	hashMap := make(map[string][]string, 0)
	for _, val := range strs {
		vals := strings.Split(val, "")
		sort.Strings(vals)
		key := strings.Join(vals, "")
		v, ok := hashMap[key]
		if ok {
			v = append(v, val)
			hashMap[key] = v
		} else {
			hashMap[key] = []string{val}
		}
	}

	res := [][]string{}

	for _, v := range hashMap {
		res = append(res, v)
	}
	return res
}
