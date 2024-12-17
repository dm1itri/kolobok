package main

import (
	"fmt"
)

type Node struct {
	left  *Node
	right *Node
	value int
}

func GetMinimumNode(root *Node) *Node {
	for root.left != nil {
		root = root.left
	}
	return root
}

func Insert(root *Node, value int) *Node {
	if root == nil {
		return &Node{nil, nil, value}
	}
	if root.value > value {
		root.left = Insert(root.left, value)
	} else {
		root.right = Insert(root.right, value)
	}
	return root
}

func Delete(root *Node, value int) *Node {
	if root == nil {
		return nil
	}
	if root.value > value {
		root.left = Delete(root.left, value)
	} else if root.value < value {
		root.right = Delete(root.right, value)
	} else {
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		}
		node := GetMinimumNode(root.right)
		root.value = node.value
		root.right = Delete(root.right, node.value)
	}
	return root
}

func PreOrder(root *Node) {
	if root != nil {
		fmt.Println(root.value)
		PreOrder(root.left)
		PreOrder(root.right)
	}
}

func InOrder(root *Node) {
	if root != nil {
		InOrder(root.left)
		fmt.Println(root.value)
		InOrder(root.right)
	}
}

func PostOrder(root *Node) {
	if root != nil {
		PostOrder(root.left)
		PostOrder(root.right)
		fmt.Println(root.value)
	}
}

func GetMaxHeight(root *Node) int {
	if root == nil {
		return 0
	}
	return max(GetMaxHeight(root.left), GetMaxHeight(root.right)) + 1
}
