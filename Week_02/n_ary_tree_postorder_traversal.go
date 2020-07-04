// 递归 时间复杂度O(M) 空间复杂度O(M)
var res []int

func postorder(root *Node) []int {
	res = make([]int, 0)
	recursive(root)
	return res
}

func recursive(node *Node) {
	if node == nil {
		return
	}
	for _, n := range node.Children {
		recursive(n)
	}
	res = append(res, node.Val)
}

// 迭代 时间复杂度O(M) 空间复杂度O(M)

func postorder(root *Node) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	stack := []*Node{root}
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append([]int{top.Val}, res...)
		for _, n := range top.Children {
			stack = append(stack, n)
		}
	}
	return res
}