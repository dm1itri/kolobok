package main

import "fmt"

type Node struct {
	symbol uint8
	prev   *Node
}

func addValue(value uint8, lastNode *Node) *Node {
	return &Node{value, lastNode}
}

func popValue(lastNode *Node) (*Node, uint8) {
	if lastNode != nil {
		return lastNode.prev, lastNode.symbol
	}
	return nil, ' '
}

func isValid(s string) bool {
	var lastNode *Node = nil
	var lastSymbol uint8
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			lastNode = addValue(s[i], lastNode)
		} else {
			lastNode, lastSymbol = popValue(lastNode)
			if s[i] == ')' && lastSymbol != '(' || s[i] == '}' && lastSymbol != '{' || s[i] == ']' && lastSymbol != '[' {
				return false
			}
		}
	}
	return lastNode == nil
}

func main() {
	var s string
	_, _ = fmt.Scan(&s)
	fmt.Println(isValid(s))
}
