/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

// 递归 时间复杂度O(n) 空间复杂度O(n)
var res [][]int

func levelOrder(root *Node) [][]int {
	res = [][]int{}
	if root != nil {
		recursive(root, 0)
	}
	return res
}

func recursive(node *Node, level int) {
	if len(res) == level {
		res = append(res, []int{})
	}
	res[level] = append(res[level], node.Val)
	for _, n := range node.Children {
		recursive(n, level+1)
	}
}

// 遍历 时间复杂度O(n) 空间复杂度O(n)
func levelOrder(root *Node) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	preLayer := []*Node{root}
	for len(preLayer) > 0 {
		preVals, currentLayer := []int{}, []*Node{}
		for _, node := range preLayer {
			preVals = append(preVals, node.Val)
			currentLayer = append(currentLayer, node.Children...)
		}
		res = append(res, preVals)
		preLayer = currentLayer
	}
	return res
}
