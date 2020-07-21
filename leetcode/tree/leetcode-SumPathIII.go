package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) int {

	return countSum(root, []int{}, sum)

}

func countSum(root *TreeNode, sumList []int, sum int) int {
	if root == nil {
		return 0
	}

	res := 0
	sumList = append(sumList, 0)
	for k := range sumList {
		sumList[k] += root.Val
		if sum == sumList[k] {
			res++
		}
	}

	// if sum == root.Val{
	// 	// fmt.Println(sum)
	// 	return 1
	// }

	// return pathSum(root.Left, sum-root.Val) + pathSum(root.Right, sum-root.Val) + pathSum(root.Left, sum) + pathSum(root.Right, sum)

	return res + countSum(root.Left, sumList, sum-root.Val) + countSum(root.Right, sumList, sum-root.Val)
}
