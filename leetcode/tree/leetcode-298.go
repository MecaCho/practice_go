package tree


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestConsecutive(root *TreeNode) int {
	maxLength := 0
	if root == nil{
		return 0
	}

	Dfs(root, root.Val,1,  &maxLength)
	return maxLength
}

func Dfs(root *TreeNode, val,length int,maxLength  *int) {

	if root == nil{
		return
	}
	// newLength := length
	if root.Val == val + 1{
		length ++
	}else {
		length = 1
	}
	if *maxLength < length{
		*maxLength = length
	}
	Dfs(root.Left, root.Val, length, maxLength)
	Dfs(root.Right, root.Val, length, maxLength)
}

// 298. 二叉树最长连续序列
// 给你一棵指定的二叉树，请你计算它最长连续序列路径的长度。
//
// 该路径，可以是从某个初始结点到树中任意结点，通过「父 - 子」关系连接而产生的任意路径。
//
// 这个最长连续的路径，必须从父结点到子结点，反过来是不可以的。
//
// 示例 1：
//
// 输入:
//
//   1
//    \
//     3
//    / \
//   2   4
//        \
//         5
//
// 输出: 3
//
// 解析: 当中，最长连续序列是 3-4-5，所以返回结果为 3
// 示例 2：
//
// 输入:
//
//   2
//    \
//     3
//    /
//   2
//  /
// 1
//
// 输出: 2
//
// 解析: 当中，最长连续序列是 2-3。注意，不是 3-2-1，所以返回 2。

// 298. Binary Tree Longest Consecutive Sequence
// Given a binary tree, find the length of the longest consecutive sequence path.
//
// The path refers to any sequence of nodes from some starting node to any node in the tree along the parent-child connections. The longest consecutive path need to be from parent to child (cannot be the reverse).
//
// Example 1:
//
// Input:
//
//   1
//    \
//     3
//    / \
//   2   4
//        \
//         5
//
// Output: 3
//
// Explanation: Longest consecutive sequence path is 3-4-5, so return 3.
// Example 2:
//
// Input:
//
//   2
//    \
//     3
//    /
//   2
//  /
// 1
//
// Output: 2
//
// Explanation: Longest consecutive sequence path is 2-3, not 3-2-1, so return 2.