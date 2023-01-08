package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

type QueueNode struct {
	*Node
	depth int
}

func createSameDepthNodeLinkedList(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}

	var queue []*QueueNode
	// TODO: change []int to LinkedList
	list := [][]int{}

	queue = append(queue, &QueueNode{Node: root, depth: 1})

	for {
		if len(queue) == 0 {
			break
		}

		qNode := queue[0]
		queue = queue[1:]
		node := qNode.Node
		depth := qNode.depth

		if len(list) < depth {
			list = append(list, []int{node.Val})
		} else {
			list[depth-1] = append(list[depth-1], node.Val)
		}

		if qNode.Left != nil {
			queue = append(queue, &QueueNode{Node: node.Left, depth: depth + 1})
		}
		if qNode.Right != nil {
			queue = append(queue, &QueueNode{Node: node.Right, depth: depth + 1})
		}
	}

	return list
}

func main() {
	fmt.Println(createSameDepthNodeLinkedList(nil))
	fmt.Println(createSameDepthNodeLinkedList(&Node{Val: 1}))
	fmt.Println(createSameDepthNodeLinkedList(&Node{Val: 1, Left: &Node{Val: 2}, Right: &Node{Val: 3}}))
	fmt.Println(createSameDepthNodeLinkedList(&Node{
		Val: 1,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val: 4,
			},
			Right: &Node{
				Val: 5,
			},
		},
		Right: &Node{
			Val: 3,
			Left: &Node{
				Val: 6,
			},
			Right: &Node{
				Val: 7,
			},
		},
	}))
}
