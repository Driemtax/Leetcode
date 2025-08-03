package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	result := []*TreeNode{}

	deleteSet := make(map[int]bool)
	for _, val := range to_delete {
		deleteSet[val] = true
	}

	var dfs func(*TreeNode, bool) *TreeNode
	dfs = func(node *TreeNode, isRoot bool) *TreeNode {
		if node == nil {
			return nil
		}

		// check if current node is to be deleted
		deleted := deleteSet[node.Val]

		// recursivly travel through the children (this implements reverse dfs)
		node.Left = dfs(node.Left, deleted)
		node.Right = dfs(node.Right, deleted)

		// case node is to be deleted
		if deleted {
			// add children to result as new roots
			if node.Left != nil {
				result = append(result, node.Left)
			}
			if node.Right != nil {
				result = append(result, node.Right)
			}
			// delete node by setting nil to the parent
			return nil
		}

		// if node is a new root and will not be deleted
		if isRoot {
			result = append(result, node)
		}

		// if we reach this we traverse backwards through the stack trace
		return node
	}

	dfs(root, true)
	return result
}
