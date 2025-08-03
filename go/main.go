package main

import "fmt"

func main() {
	// Example 1
	tree1 := buildTree([]interface{}{1, 2, 3, 4, 5, 6, 7})
	toDelete1 := []int{3, 5}
	forest1 := delNodes(tree1, toDelete1)
	fmt.Println("Example 1 output:")
	for _, tree := range forest1 {
		fmt.Println(treeToList(tree))
	}
	fmt.Println()
	// Example 2
	tree2 := buildTree([]interface{}{1, 2, 4, nil, 3})
	toDelete2 := []int{3}
	forest2 := delNodes(tree2, toDelete2)
	fmt.Println("Example 2 output:")
	for _, tree := range forest2 {
		fmt.Println(treeToList(tree))
	}
}

// Auxilary function to build a tree from a level-order-list
func buildTree(vals []interface{}) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}
	nodes := make([]*TreeNode, len(vals))
	for i, v := range vals {
		if v != nil {
			nodes[i] = &TreeNode{Val: v.(int)}
		}
	}
	kids := nodes[1:]
	for _, node := range nodes {
		if node != nil {
			if len(kids) > 0 {
				node.Left = kids[0]
				kids = kids[1:]
			}
			if len(kids) > 0 {
				node.Right = kids[0]
				kids = kids[1:]
			}
		}
	}
	return nodes[0]
}

// Hilfsfunktion, um einen Binary Tree in Level-Order-Liste umzuwandeln
func treeToList(root *TreeNode) []interface{} {
	var result []interface{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			result = append(result, nil)
			continue
		}
		result = append(result, node.Val)
		queue = append(queue, node.Left, node.Right)
	}
	// Entferne abschließende nils
	for len(result) > 0 && result[len(result)-1] == nil {
		result = result[:len(result)-1]
	}
	return result
}
