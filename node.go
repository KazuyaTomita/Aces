package main

import "reflect"

type Node struct {
	State State
	Children []Node
	Parent Node
}

func (node *Node) ==() bool {
	return true
}

func (node *Node) getActionFor(nextNode *Node) Action {
	return Action{}
}

// game is over when the button player did an action but raise
func (node *Node) isTerminal() bool {
	return false
}

func (node *Node) isExpanded() bool {
	// if the node is expanded, the node should have the parent
	return reflect.DeepEqual(node.Parent, Node{})
}



