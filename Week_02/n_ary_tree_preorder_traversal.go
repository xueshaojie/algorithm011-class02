// 递归 时间复杂度O(M) 空间复杂度O(M)

var res []int

func preorder(root *Node) []int {
	res = make([]int, 0)
	recursive(root)
	return res
}

func recursive(root *Node) {
	if root == nil {
		return
	}
	res = append(res, root.Val)
	for _, node := range root.Children {
		recursive(node)
	}
}

// 迭代 时间复杂度O(M) 空间复杂度O(M)
func preorder(root *Node) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	stack := []*Node{root}
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, top.Val)
		for i := len(top.Children) - 1; i >= 0; i-- {
			stack = append(stack, top.Children[i])
		}
	}
	return res
}