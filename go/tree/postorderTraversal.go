package tree

/**
 * @param root: A Tree
 * @return: Postorder in ArrayList which contains node values.
 */
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack := []*TreeNode{root}
	result := []int{}
	var prev *TreeNode

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		if prev == nil || prev.Left == curr || prev.Right == curr {
			if curr.Left != nil {
				stack = append(stack, curr.Left)
			} else if curr.Right != nil {
				stack = append(stack, curr.Right)
			}
		} else if curr.Left == prev {
			if curr.Right != nil {
				stack = append(stack, curr.Right)
			}
		} else {
			result = append(result, curr.Val)
			stack = stack[:len(stack)-1]
		}
		prev = curr
	}

	return result
}
