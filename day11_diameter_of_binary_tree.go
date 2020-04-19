package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lHeight := height(root.Left)
	rHeight := height(root.Right)

	lDiameter := diameterOfBinaryTree(root.Left)
	rDiameter := diameterOfBinaryTree(root.Right)

	return max(lHeight+rHeight, max(lDiameter, rDiameter))

}

func height(tree *TreeNode) int {
	if tree == nil {
		return 0
	}

	return 1 + max(height(tree.Left), height(tree.Right))
}

func max(a, b int) int {

	if a > b {
		return a
	}

	return b
}
