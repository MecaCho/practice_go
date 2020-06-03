package tree

import (
	"fmt"
	"math"
)

// 270. 最接近的二叉搜索树值
// 给定一个不为空的二叉搜索树和一个目标值 target，请在该二叉搜索树中找到最接近目标值 target 的数值。
//
// 注意：
//
// 给定的目标值 target 是一个浮点数
// 题目保证在该二叉搜索树中只会存在一个最接近目标值的数
// 示例：
//
// 输入: root = [4,2,5,1,3]，目标值 target = 3.714286
//
//    4
//   / \
//  2   5
// / \
// 1   3
//
// 输出: 4

// 270. Closest Binary Search Tree Value
// Given a non-empty binary search tree and a target value, find the value in the BST that is closest to the target.
//
// Note:
//
// Given target value is a floating point.
// You are guaranteed to have only one unique value in the BST that is closest to the target.
// Example:
//
// Input: root = [4,2,5,1,3], target = 3.714286
//
//    4
//   / \
//  2   5
// / \
// 1   3
//
// Output: 4

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func closestValue(root *TreeNode, target float64) int {

	vals := []int{}

	DFS(root, &vals)

	min_Sub := math.MaxFloat64

	res := 0

	for i := range vals {
		subVal := math.Abs(target - float64(vals[i]))
		if subVal < min_Sub {
			min_Sub = subVal
			res = vals[i]
		}
	}

	fmt.Println(vals)

	return res
}

func DFS(root *TreeNode, vals *[]int) {

	if root == nil {
		return
	}

	DFS(root.Left, vals)
	*vals = append(*vals, root.Val)
	DFS(root.Right, vals)

}
