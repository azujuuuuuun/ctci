package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countLevel(node *TreeNode, level int) (balanced bool, subTreeLevel int) {
	if node == nil {
		return true, level
	}
	if node.Left == nil && node.Right == nil {
		return true, level
	}
	leftResult := true
	leftLevel := level
	rightResult := true
	rightLevel := level
	maxLevel := level
	if node.Left != nil {
		leftResult, leftLevel = countLevel(node.Left, level+1)
	}
	if node.Right != nil {
		rightResult, rightLevel = countLevel(node.Right, level+1)
	}
	if leftLevel > rightLevel {
		maxLevel = leftLevel
		diff := leftLevel - rightLevel
		if diff > 1 {
			return false, maxLevel
		}
	} else {
		maxLevel = rightLevel
		diff := rightLevel - leftLevel
		if diff > 1 {
			return false, maxLevel
		}
	}
	if !leftResult || !rightResult {
		return false, maxLevel
	}
	return true, maxLevel
}

func isBalanced(root *TreeNode) bool {
	b, _ := countLevel(root, 1)
	return b
}

func main() {
	fmt.Println(isBalanced(nil) == true)
	fmt.Println(isBalanced(&TreeNode{Val: 1}) == true)
	fmt.Println(isBalanced(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}) == true)
	fmt.Println(isBalanced(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}) == false)
	fmt.Println(isBalanced(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}) == true)
}
