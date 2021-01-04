class Solution:
    """
    @param root: A Tree
    @return: Inorder in ArrayList which contains node values.
    """

    def inorderTraversal(self, root):
        if not root:
            return []

        stack = []
        while root:
            stack.append(root)
            root = root.left
        res = []
        while stack:
            node = stack.pop()
            res.append(node.val)

            node = node.right
            while node:
                stack.append(node)
                node = node.left

        return res

"""
Morris algorithm
"""

class Solution:
    """
    @param root: A Tree
    @return: Inorder in ArrayList which contains node values.
    """

    def inorderTraversal(self, root):
        nums = []
        cur = None

        while root:
            if root.left:
                cur = root.left
                while cur.right and cur.right != root:
                    cur = cur.right

                if cur.right == root:
                    nums.append(root.val)
                    cur.right = None
                    root = root.right
                else:
                    cur.right = root
                    root = root.left
            else:
                nums.append(root.val)
                root = root.right

        return nums



