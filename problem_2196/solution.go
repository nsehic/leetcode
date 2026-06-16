package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createBinaryTree(descriptions [][]int) *TreeNode {
	hasParent := make(map[int]bool)
	nodes := make(map[int]*TreeNode)

	for _, description := range descriptions {
		parentValue := description[0]
		childValue := description[1]
		isLeft := description[2] == 1

		if _, ok := nodes[parentValue]; !ok {
			nodes[parentValue] = &TreeNode{Val: parentValue}
		}

		if _, ok := nodes[childValue]; !ok {
			nodes[childValue] = &TreeNode{Val: childValue}
		}

		if isLeft {
			nodes[parentValue].Left = nodes[childValue]
		} else {
			nodes[parentValue].Right = nodes[childValue]
		}

		hasParent[childValue] = true
	}

	for _, description := range descriptions {
		parent := description[0]

		if !hasParent[parent] {
			return nodes[parent]
		}
	}

	return nil
}
