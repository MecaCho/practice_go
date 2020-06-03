package tree

// 156. 上下翻转二叉树
// 给定一个二叉树，其中所有的右节点要么是具有兄弟节点（拥有相同父节点的左节点）的叶节点，要么为空，将此二叉树上下翻转并将它变成一棵树， 原来的右节点将转换成左叶节点。返回新的根。
//
// 例子:
//
// 输入: [1,2,3,4,5]
//
//    1
//   / \
//  2   3
// / \
// 4   5
//
// 输出: 返回二叉树的根 [4,5,2,#,#,3,1]
//
//   4
//  / \
// 5   2
//    / \
//   3   1
// 说明:
//
// 对 [4,5,2,#,#,3,1] 感到困惑? 下面详细介绍请查看 二叉树是如何被序列化的。
//
// 二叉树的序列化遵循层次遍历规则，当没有节点存在时，'#' 表示路径终止符。
//
// 这里有一个例子:
//
//   1
//  / \
// 2   3
//    /
//   4
//    \
//     5
// 上面的二叉树则被序列化为 [1,2,3,#,#,4,#,#,5].

// 156. Binary Tree Upside Down
// Given a binary tree where all the right nodes are either leaf nodes with a sibling (a left node that shares the same parent node) or empty, flip it upside down and turn it into a tree where the original right nodes turned into left leaf nodes. Return the new root.
//
// Example:
//
// Input: [1,2,3,4,5]
//
//    1
//   / \
//  2   3
// / \
// 4   5
//
// Output: return the root of the binary tree [4,5,2,#,#,3,1]
//
//   4
//  / \
// 5   2
//    / \
//   3   1
// Clarification:
//
// Confused what [4,5,2,#,#,3,1] means? Read more below on how binary tree is serialized on OJ.
//
// The serialization of a binary tree follows a level order traversal, where '#' signifies a path terminator where no node exists below.
//
// Here's an example:
//
//   1
//  / \
// 2   3
//    /
//   4
//    \
//     5
// The above binary tree is serialized as [1,2,3,#,#,4,#,#,5].

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func upsideDownBinaryTree(root *TreeNode) *TreeNode {

	if root == nil || root.Left == nil {
		return root
	}

	root_ := root
	for root != nil && root.Left != nil && root.Left.Left != nil {
		root = root.Left
	}

	l := root.Left
	r := root.Right

	root.Left = nil
	root.Right = nil

	l.Left = r
	l.Right = upsideDownBinaryTree(root_)

	return l

}
