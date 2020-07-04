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

func preorderTraversal(root *TreeNode) []int {
	res = make([]int, 0)
	recursive(root)
	return res
}

func recursive(node *TreeNode) {
	if node == nil {
		return
	}
	res = append(res, node.Val)
	recursive(node.Left)
	recursive(node.Right)
}

// 栈 时间复杂度O(n) 空间复杂度O(n)
func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return res
}