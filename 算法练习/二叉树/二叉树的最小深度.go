// author:zfy  date:2022/10/2

package main

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 递归
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right != nil {
		return 1 + minDepth(root.Right)
	}
	if root.Right == nil && root.Left != nil {
		return 1 + minDepth(root.Left)
	}
	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}

// 迭代

func minDepth1(root *TreeNode) int {
	dep := 0
	queue := make([]*TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}
	for l := len(queue); l > 0; {
		dep++
		for ; l > 0; l-- {
			node := queue[0]
			if node.Left == nil && node.Right == nil {
				return dep
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = queue[1:]
		}
		l = len(queue)
	}
	return dep
}
