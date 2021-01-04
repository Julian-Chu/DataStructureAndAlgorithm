package tree

func preordertraversalMorris(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	for root != nil {
		if root.Left != nil {
			cur := root.Left
			for cur.Right != nil && cur.Right != root {
				cur = cur.Right
			}

			if cur.Right == root {
				cur.Right = nil
				root = root.Right
			} else {
				res = append(res, root.Val)
				cur.Right = root
				root = root.Left
			}
		} else {
			res = append(res, root.Val)
			root = root.Right
		}
	}

	return res
}
