package tree

/**
 * Definition for Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Parent *Node
 * }
 */

type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
}

func inorderSuccessor(node *Node) *Node {
	if node == nil {
		return nil
	}

	if node.Right == nil {
		for node.Parent != nil && node.Parent.Right == node {
			node = node.Parent
		}
		return node.Parent
	}
	node = node.Right
	res := node
	for node != nil {
		res = node
		node = node.Left
	}
	return res

}

// 510. 二叉搜索树中的中序后继 II
// 给定一棵二叉搜索树和其中的一个节点 node ，找到该节点在树中的中序后继。
//
// 如果节点没有中序后继，请返回 null 。
//
// 一个结点 node 的中序后继是键值比 node.val大所有的结点中键值最小的那个。
//
// 你可以直接访问结点，但无法直接访问树。每个节点都会有其父节点的引用。节点定义如下：
//
// class Node {
//    public int val;
//    public Node left;
//    public Node right;
//    public Node parent;
// }
//
//
// 进阶：
//
// 你能否在不访问任何结点的值的情况下解决问题?
//
//
//
// 示例 1:
//
//
//
// 输入: tree = [2,1,3], node = 1
// 输出: 2
// 解析: 1 的中序后继结点是 2 。注意节点和返回值都是 Node 类型的。
// 示例 2:
//
//
//
// 输入: tree = [5,3,6,2,4,null,null,1], node = 6
// 输出: null
// 解析: 该结点没有中序后继，因此返回 null 。
// 示例 3:
//
//
//
// 输入: tree = [15,6,18,3,7,17,20,2,4,null,13,null,null,null,null,null,null,null,null,9], node = 15
// 输出: 17
// 示例 4:
//
//
//
// 输入: tree = [15,6,18,3,7,17,20,2,4,null,13,null,null,null,null,null,null,null,null,9], node = 13
// 输出: 15
//
//
// 提示：
//
// -10^5 <= Node.val <= 10^5
// 1 <= Number of Nodes <= 10^4
// 树中各结点的值均保证唯一。

// 510. Inorder Successor in BST II
// Given a node in a binary search tree, find the in-order successor of that node in the BST.
//
// If that node has no in-order successor, return null.
//
// The successor of a node is the node with the smallest key greater than node.val.
//
// You will have direct access to the node but not to the root of the tree. Each node will have a reference to its parent node. Below is the definition for Node:
//
// class Node {
//    public int val;
//    public Node left;
//    public Node right;
//    public Node parent;
// }
//
//
// Follow up:
//
// Could you solve it without looking up any of the node's values?
//
//
//
// Example 1:
//
//
// Input: tree = [2,1,3], node = 1
// Output: 2
// Explanation: 1's in-order successor node is 2. Note that both the node and the return value is of Node type.
// Example 2:
//
//
// Input: tree = [5,3,6,2,4,null,null,1], node = 6
// Output: null
// Explanation: There is no in-order successor of the current node, so the answer is null.
// Example 3:
//
//
// Input: tree = [15,6,18,3,7,17,20,2,4,null,13,null,null,null,null,null,null,null,null,9], node = 15
// Output: 17
// Example 4:
//
//
// Input: tree = [15,6,18,3,7,17,20,2,4,null,13,null,null,null,null,null,null,null,null,9], node = 13
// Output: 15
// Example 5:
//
// Input: tree = [0], node = 0
// Output: null
//
//
// Constraints:
//
// -10^5 <= Node.val <= 10^5
// 1 <= Number of Nodes <= 10^4
// All Nodes will have unique values.

//  solutions

// 前驱和后继
// 前驱：指的是中序遍历的上一个结点，或者说是比当前结点小的最大结点。
// 后继：指的是中序遍历的下一个结点，或者说是比当前结点大的最小结点。
//
//
// 方法：迭代
// 这里可能有两种情况：
//
// node 结点有右孩子，且它的后继在树中相对较低的位置，为了找到后继，我们应该向右走一次，再尽可能的向左走。
//
//
// node 结点没有右孩子，则它的后继在树中相对较高的位置，为了找到后继，我们向上走到直到结点 tmp 的左孩子是 node 的父节点时，则 node 的后继为 tmp。若没有找到符合条件的结点说明该结点没有后继。
//
//
//
//
// 算法：
//
// 若 node 结点有右孩子，则它的后继在树中相对较低的位置。我们向右走一次，再尽可能的向左走，返回最后所在的结点。
// 若 node 结点没有右孩子，则它的后继在树中相对较高的位置。我们向上走到直到结点 tmp 的左孩子是 node 的父节点时，则 node 的后继为 tmp。
// PythonJavaC++
//
// class Solution:
//    def inorderSuccessor(self, node: 'Node') -> 'Node':
//        # the successor is somewhere lower in the right subtree
//        if node.right:
//            node = node.right
//            while node.left:
//                node = node.left
//            return node
//
//        # the successor is somewhere upper in the tree
//        while node.parent and node == node.parent.right:
//            node = node.parent
//        return node.parent
// 复杂度分析
//
// 时间复杂度：\mathcal{O}(H)O(H)。其中 HH 为数的高度。平均时间复杂度为 \mathcal{O}(\log N)O(logN)，最坏的事件复杂度为 \mathcal{O}(N)O(N)，其中 NN 为树的结点数。
// 空间复杂度：\mathcal{O}(1)O(1)，在计算的过程中没有使用额外空间。
//
// 作者：LeetCode
// 链接：https://leetcode-cn.com/problems/inorder-successor-in-bst-ii/solution/er-cha-sou-suo-shu-zhong-de-zhong-xu-hou-ji-ii-by-/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
