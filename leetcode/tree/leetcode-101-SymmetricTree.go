package tree

import "math"

// 101. Symmetric Tree
// Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).
//
// For example, this binary tree [1,2,2,3,4,4,3] is symmetric:
//
//    1
//   / \
//  2   2
// / \ / \
// 3  4 4  3
//
//
// But the following [1,2,2,null,3,null,3] is not:
//
//    1
//   / \
//  2   2
//   \   \
//   3    3
//
//
// Follow up: Solve it both recursively and iteratively.

// 101. 对称二叉树
// 给定一个二叉树，检查它是否是镜像对称的。
//
//
//
// 例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
//
//    1
//   / \
//  2   2
// / \ / \
// 3  4 4  3
//
//
// 但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
//
//    1
//   / \
//  2   2
//   \   \
//   3    3
//
//
// 进阶：
//
// 你可以运用递归和迭代两种方法解决这个问题吗？

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isMiorr(p, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && isMiorr(q.Right, p.Left) && isMiorr(q.Left, p.Right)

}

func isSymmetric(root *TreeNode) bool {
	//
	// if root.Left == nil && root.Right == nil {
	// 	return true
	// }
	//
	// if root.Left == nil || root.Right == nil {
	// 	return false
	// }
	//
	// if root.Left.Val != root.Right.Val {
	// 	return false
	// }

	return isMiorr(root, root)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isMiorr(p, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && isMiorr(q.Right, p.Left) && isMiorr(q.Left, p.Right)

}

// 迭代

func isSymmetric1(root *TreeNode) bool {

	// return isMiorr(root, root)
	queue := []*TreeNode{}

	queue = append(queue, root)

	for len(queue) > 0 {

		vals := []int{}
		for _, node := range queue {
			if node != nil {
				vals = append(vals, node.Val)
			} else {
				vals = append(vals, math.MinInt16)
			}

		}

		for i := range vals {
			if vals[i] != vals[len(vals)-1-i] {
				return false
			}
		}

		newQueue := []*TreeNode{}

		for i := range queue {
			if queue[i] != nil {
				newQueue = append(newQueue, queue[i].Left)
				newQueue = append(newQueue, queue[i].Right)
			}

		}

		queue = newQueue

	}

	return true

}
