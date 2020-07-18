# 学习笔记

## 知识点

bfs广度优先搜索，代码模板：

```go
func bfs(graph [][]string, start end string) {
  queue := []string{start}
  visited := map[string]bool
  for len(queue) > 0 {
    node := queue[0]
    queue = queue[1:]
    visited[node] = true
    process(node)
    nodes := generateRelatedNodes(node)
    queue = append(queue, nodes...)
    # other processing work
  }
}
```

dfs深度优先搜索，代码模板：

```go
visited := map[*TreeNode]bool{}
var dfs func(*TreeNode)
dfs = func(node *TreeNode) {
  if _, ok := visited[node]; ok {
    return 
  }
 	visited[node] = true
  dfs(node.Left)
  dfs(node.Right)
}
dfs(node)
```

贪心算法。在当前情况下使用最优解，通过局部最优解，得到全局最优解。

二分查找，代码模板

```go
left, right := 0, len(array)-1
for left <= right {
	mid := (left+right)/2
	if array[mid] == target {
			return 
	} else if array[mid] < target {
		  left = mid + 1
	} else {
			right = mid -1 
	}
}
```

## 练习题

使用二分查找，寻找一个半有序数组 [4, 5, 6, 7, 0, 1, 2] 中间无序的地方。

```go
func findIndex(nums []int) int {
  l, r, mid := 0, len(nums)-1, 0
  for l <= r {
    mid = (l+r)/2
    if mid == 0 {
	    return -1
    }
    if nums[mid] < nums[mid-1] && nums[mid] < nums[mid+1] {
      return mid
    } else if nums[0] < nums[mid] && nums[mid] > nums[r] {
      l = mid + 1
    } else {
      r = mid - 1
    }
  }
  return -1
}
```