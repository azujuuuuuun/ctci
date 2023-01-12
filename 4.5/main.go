package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMaxVal(node *TreeNode) int {
	if node == nil {
		return math.MinInt
	}
	left := float64(getMaxVal(node.Left))
	right := float64(getMaxVal(node.Right))
	return int(math.Max(math.Max(left, right), float64(node.Val)))
}

func getMinVal(node *TreeNode) int {
	if node == nil {
		return math.MaxInt
	}
	left := float64(getMinVal(node.Left))
	right := float64(getMinVal(node.Right))
	return int(math.Min(math.Min(left, right), float64(node.Val)))
}

func isBinarySearchTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftMax := getMaxVal(root.Left)
	rightMin := getMinVal(root.Right)
	if leftMax <= root.Val && root.Val <= rightMin {
		return isBinarySearchTree(root.Left) && isBinarySearchTree(root.Right)
	} else {
		return false
	}
}

func main() {
	fmt.Println(isBinarySearchTree(nil) == true)
	fmt.Println(isBinarySearchTree(&TreeNode{Val: 1}) == true)
	fmt.Println(isBinarySearchTree(&TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}) == true)
	fmt.Println(isBinarySearchTree(&TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 3,
		},
		Right: &TreeNode{
			Val: 1,
		},
	}) == false)
}
