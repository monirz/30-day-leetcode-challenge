package main

// Definition for a binary tree node.
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func bstFromPreorder(preorder []int) *TreeNode {

	val := preorder[0]

	node := &TreeNode{Val: val, Left: nil, Right: nil}

	index, left := bstHelper(preorder, 1, node)
	node.Left = left
	_, right := bstHelper(preorder, index, nil)
	node.Right = right

	return node
}

func bstHelper(preorder []int, start int, tree *TreeNode) (int, *TreeNode) {
	if start >= len(preorder) {
		return start, nil
	}
	v := preorder[start]
	if tree != nil && v > tree.Val {
		return start, nil
	}

	node := &TreeNode{v, nil, nil}
	index, left := bstHelper(preorder, start+1, node)
	node.Left = left
	next, right := bstHelper(preorder, index, tree)
	node.Right = right
	return next, node
}
