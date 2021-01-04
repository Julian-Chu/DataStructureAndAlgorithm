package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/**
 * @param root: A Tree
 * @return: Inorder in ArrayList which contains node values.
 */
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	stack := make([]*TreeNode, 0)
	for root != nil {
		stack = append(stack, root)
		root = root.Left
	}
	res := []int{}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)

		node = node.Right
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
	}
	return res
}
