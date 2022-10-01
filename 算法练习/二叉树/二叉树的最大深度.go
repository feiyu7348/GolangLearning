// author:zfy  date:2022/10/1

package main

//definition for a binary tree node.
type treenode struct {
	val   int
	left  *treenode
	right *treenode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 递归
func maxdepth1(root *treenode) int {
	if root == nil {
		return 0
	}
	return max(maxdepth(root.left), maxdepth(root.right)) + 1
}

// 遍历
func maxdepth(root *treenode) int {
	levl := 0
	queue := make([]*treenode, 0)
	if root != nil {
		queue = append(queue, root)
	}
	for l := len(queue); l > 0; {
		for ; l > 0; l-- {
			node := queue[0]
			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
			queue = queue[1:]
		}
		levl++
		l = len(queue)
	}
	return levl
}
