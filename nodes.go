package main

import "fmt"

type Node struct {
	Next  *Node
	Value interface{}
}

func NodesTest() {
	var root *Node = &Node{nil, 10}
	fillNode(root, 9).Next = root
	
	runNodesTest(root)
}

func runNodesTest(root *Node) {
	visited := map[*Node]bool{}
	for n := root; n != nil; n= n.Next {
		if visited[n] {
			fmt.Println("cycle detected")
			break
		}		
		visited[n] = true
		fmt.Println(n.Value)
	}
}

func fillNode(parent *Node, val int) *Node {
	if val == -1 {
		return parent
	}
	next := &Node{nil, val}
	parent.Next = next
	val--
	return fillNode(next, val)
}
