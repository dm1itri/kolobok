package main

import "fmt"

type Node struct {
	value string
	prev  *Node
	next  *Node
}

func addValue(value string, lastNode *Node) *Node {
	if lastNode != nil {
		lastNode.next = &Node{value, lastNode, nil}
		return lastNode.next
	}
	return &Node{value, lastNode, nil}
}

func pop(lastNode *Node) (*Node, string) {
	if lastNode != nil {
		if lastNode.prev != nil {
			lastNode.prev.next = nil
		}
		return lastNode.prev, lastNode.value
	}
	return nil, ""
}

func simplifyPath(path string) string {
	var newString string
	var lastNode *Node
	path += "/"
	firstNode := lastNode
	for i := 1; i < len(path); i++ {
		if path[i] == '/' {
			if newString == ".." {
				lastNode, _ = pop(lastNode)
			} else if newString != "." && newString != "" {
				lastNode = addValue(newString, lastNode)
			}
			newString = ""
		} else {
			newString += string(path[i])
		}
		if firstNode == nil || lastNode == nil {
			firstNode = lastNode
		}
	}
	var res string
	if firstNode == nil {
		return "/"
	}
	for firstNode != nil {
		res += "/"
		res += firstNode.value
		firstNode = firstNode.next
	}
	return res
}

func main() {
	var s string
	_, _ = fmt.Scan(&s)
	fmt.Println(simplifyPath(s))
}
