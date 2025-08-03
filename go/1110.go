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
	// iterate backwards over to_delete, since higher nodes have a higher probability to be a node without children in a bineary tree
	for i := len(to_delete) - 1; i >= 0; i-- {
		to_del := to_delete[i]

		// find node to delete
		nodeToDelete := findNode(root, to_del) // nil if node not found -> skip to_delete[i]
		if nodeToDelete == nil {
			continue
		}

		// special case root == nodeToDelete
		if root == nodeToDelete {
			// add children to result
			if root.Left != nil {
				result = append(result, root.Left)
			}
			if root.Right != nil {
				result = append(result, root.Right)
			}

			root = nil
			continue
		}

		// find parent and delete connection
		parent := findParent(root, nodeToDelete)
		if parent != nil {
			// add children to result
			if nodeToDelete.Left != nil {
				result = append(result, nodeToDelete.Left)
			}
			if nodeToDelete.Right != nil {
				result = append(result, nodeToDelete.Right)
			}

			// now "delete" node through deleting the connection to parent
			if parent.Left == nodeToDelete {
				parent.Left = nil
			} else if parent.Right == nodeToDelete {
				parent.Right = nil
			}
		}

	}

	// add the remaining nodes from the original tree that still have a connection to root
	if root != nil {
		result = append(result, root)
	}

	return result
}

func findNode(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == target {
		return root
	}

	// Look through left half
	if leftResult := findNode(root.Left, target); leftResult != nil {
		return leftResult
	}

	// look trough right half
	if rightResult := findNode(root.Right, target); rightResult != nil {
		return rightResult
	}

	return nil
}

func findParent(root *TreeNode, target *TreeNode) *TreeNode {
	if root == nil || target == nil {
		return nil
	}

	// Prüfe ob target ein direktes Kind von root ist
	if root.Left == target || root.Right == target {
		return root
	}

	// Suche rekursiv im linken Teilbaum
	if leftResult := findParent(root.Left, target); leftResult != nil {
		return leftResult
	}

	// Suche rekursiv im rechten Teilbaum
	if rightResult := findParent(root.Right, target); rightResult != nil {
		return rightResult
	}

	return nil
}
