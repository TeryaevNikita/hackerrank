package main

//https://www.geeksforgeeks.org/count-smaller-elements-on-right-side/
import "fmt"

type Node struct {
	key    int
	left   *Node
	right  *Node
	height int
	size   int
}

func newNodeF(key int) *Node {
	node := &Node{}
	node.key = key
	node.height = 1
	node.size = 1
	return node
}

func max(a, b int) int {
	maxVal := a
	if b > maxVal {
		maxVal = b
	}
	return maxVal
}

func height(node *Node) int {
	if node == nil {
		return 0
	}
	return node.height
}

func size(node *Node) int {
	if node == nil {
		return 0
	}
	return node.size
}

func getBalance(node *Node) int {
	if node == nil {
		return 0
	}

	return height(node.left) - height(node.right)
}

func insertNode(node *Node, key int, counters []int) *Node {

	if node == nil {
		return newNodeF(key)
	}

	if key < node.key {
		node.left = insertNode(node.left, key, counters)
	} else {
		node.right = insertNode(node.right, key, counters)
	}

	node.height = max(height(node.left), height(node.right)) + 1
	node.size = size(node.left) + size(node.right) + 1
	return node
}

//
//length := len(arr)
//
//counters := make([]int32, length)
//
//fmt.Println(arr)
//fmt.Println(counters)
//
//var root *Node
//
//for i := length - 1; i >= 0; i-- {
//root = insertNode(root, arr[i], counters)
//}
//
//return []int32{}

func main() {
	var arr []int
	arr = append(arr, 5, 7)

	length := len(arr)

	counters := make([]int, length)

	fmt.Println(arr)
	fmt.Println(counters)

	var root *Node

	for i := length - 1; i >= 0; i-- {
		root = insertNode(root, arr[i], counters)
		fmt.Println(root)
	}

	fmt.Println(arr)
}
