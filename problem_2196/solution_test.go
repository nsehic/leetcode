package main

import (
	"slices"
	"testing"
)

func levelOrderTree(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var result []int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		result = append(result, node.Val)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return result
}

func TestCreateBinaryTree(t *testing.T) {
	tests := []struct {
		input [][]int
		want  []int
	}{
		{
			input: [][]int{},
			want:  []int{},
		},
		{
			input: [][]int{},
			want:  []int{},
		},
	}

	for _, test := range tests {
		got := levelOrderTree(createBinaryTree(test.input))
		if !slices.Equal(got, test.want) {
			t.Errorf("createBinaryTree(%v) expected %v, actual %v", test.input, test.want, got)
		}
	}
}
