package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//  tips

// Use recursion. If root.val <= V, you split root.right into two halves, then join it's left half back on root.right.

func splitBST(root *TreeNode, V int) []*TreeNode {

	if root == nil {
		return []*TreeNode{nil, nil}
	} else if root.Val <= V {
		newTrees := splitBST(root.Right, V)
		root.Right = newTrees[0]
		return []*TreeNode{root, newTrees[1]}
	}
	newTrees := splitBST(root.Left, V)
	root.Left = newTrees[1]
	return []*TreeNode{newTrees[0], root}
}

// 776. Split BST
// Given a Binary Search Tree (BST) with root node root, and a target value V, split the tree into two subtrees where one subtree has nodes that are all smaller or equal to the target value, while the other subtree has all nodes that are greater than the target value.  It's not necessarily the case that the tree contains a node with value V.
//
// Additionally, most of the structure of the original tree should remain.  Formally, for any child C with parent P in the original tree, if they are both in the same subtree after the split, then node C should still have the parent P.
//
// You should output the root TreeNode of both subtrees after splitting, in any order.
//
// Example 1:
//
// Input: root = [4,2,6,1,3,5,7], V = 2
// Output: [[2,1],[4,3,6,null,null,5,7]]
// Explanation:
// Note that root, output[0], and output[1] are TreeNode objects, not arrays.
//
// The given tree [4,2,6,1,3,5,7] is represented by the following diagram:
//
//          4
//        /   \
//      2      6
//     / \    / \
//    1   3  5   7
//
// while the diagrams for the outputs are:
//
//          4
//        /   \
//      3      6      and    2
//            / \           /
//           5   7         1
// Note:
//
// The size of the BST will not exceed 50.
// The BST is always valid and each node's value is different.

// 776. 拆分二叉搜索树
// 给你一棵二叉搜索树（BST）、它的根结点 root 以及目标值 V。
//
// 请将该树按要求拆分为两个子树：其中一个子树结点的值都必须小于等于给定的目标值 V；另一个子树结点的值都必须大于目标值 V；树中并非一定要存在值为 V 的结点。
//
// 除此之外，树中大部分结构都需要保留，也就是说原始树中父节点 P 的任意子节点 C，假如拆分后它们仍在同一个子树中，那么结点 P 应仍为 C 的子结点。
//
// 你需要返回拆分后两个子树的根结点 TreeNode，顺序随意。
//
//
//
// 示例：
//
// 输入：root = [4,2,6,1,3,5,7], V = 2
// 输出：[[2,1],[4,3,6,null,null,5,7]]
// 解释：
// 注意根结点 output[0] 和 output[1] 都是 TreeNode 对象，不是数组。
//
// 给定的树 [4,2,6,1,3,5,7] 可化为如下示意图：
//
//          4
//        /   \
//      2      6
//     / \    / \
//    1   3  5   7
//
// 输出的示意图如下：
//
//          4
//        /   \
//      3      6       和    2
//            / \           /
//           5   7         1
//
//
// 提示：
//
// 二叉搜索树节点个数不超过 50
// 二叉搜索树始终是有效的，并且每个节点的值都不相同
