// 递归 时间复杂度O(n+m) 空间复杂度O(n+m)
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val > l2.Val {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	} else {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
}

// 遍历 时间复杂度O(n+m) 空间复杂度O(1)
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	prehead := &ListNode{}
	pre := prehead

	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			pre.Next = l2
			l2 = l2.Next
		} else {
			pre.Next = l1
			l1 = l1.Next
		}
		pre = pre.Next
	}

	if l1 == nil {
		pre.Next = l2
	} else {
		pre.Next = l1
	}

	return prehead.Next
}