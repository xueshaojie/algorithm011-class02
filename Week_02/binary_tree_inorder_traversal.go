/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 递归 时间复杂度O(n) 空间复杂度O(n)
var res []int

func inorderTraversal(root *TreeNode) []int {
	res = make([]int, 0)
	recursive(root)
	return res
}

func recursive(node *TreeNode) {
	if node == nil {
		return
	}
	recursive(node.Left)
	res = append(res, node.Val)
	recursive(node.Right)
}

// 栈 时间复杂度O(n) 空间复杂度O(n)
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	stack := []*TreeNode{}

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return res
}
