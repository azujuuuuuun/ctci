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
	if node.Left == nil && node.Right == nil {
		return node.Val
	}
	if node.Left != nil && node.Right == nil {
		if node.Val >= node.Left.Val {
			return node.Val
		} else {
			return math.MaxInt
		}
	}
	if node.Left == nil && node.Right != nil {
		if node.Val <= node.Right.Val {
			return node.Right.Val
		} else {
			return math.MaxInt
		}
	}
	left := getMaxVal(node.Left)
	right := getMaxVal(node.Right)
	if left == math.MaxInt || right == math.MaxInt {
		return math.MaxInt
	}
	if left > node.Val || node.Val > right {
		return math.MaxInt
	}
	return int(math.Max(math.Max(float64(left), float64(right)), float64(node.Val)))
}

func getMinVal(node *TreeNode) int {
	if node.Left == nil && node.Right == nil {
		return node.Val
	}
	if node.Left != nil && node.Right == nil {
		if node.Val >= node.Left.Val {
			return node.Left.Val
		} else {
			return math.MinInt
		}
	}
	if node.Left == nil && node.Right != nil {
		if node.Val <= node.Right.Val {
			return node.Val
		} else {
			return math.MinInt
		}
	}
	left := getMinVal(node.Left)
	right := getMinVal(node.Right)
	if left == math.MinInt || right == math.MinInt {
		return math.MinInt
	}
	if left > node.Val || node.Val > right {
		return math.MinInt
	}
	return int(math.Min(math.Min(float64(left), float64(right)), float64(node.Val)))
}

func isBinarySearchTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftMax := root.Val - 1
	if root.Left != nil {
		leftMax = getMaxVal(root.Left)
	}
	if leftMax == math.MaxInt {
		return false
	}
	rightMin := root.Val + 1
	if root.Right != nil {
		rightMin = getMinVal(root.Right)
	}
	if rightMin == math.MinInt {
		return false
	}
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
