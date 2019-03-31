package main

import "reflect"

type Node struct {
	State *State
	Children []Node
	Parent *Node
}

func (node *Node) getActionFor(nextNode *Node) Action {
	return Action{}
}

func (node *Node) getNextNodeAfter(action *Action) Node {
	nextState := node.State.getStateAfter(action)
	return nextState.toNode()
}

func (node *Node) addChildNode(child *Node) {
	node.Children = append(node.Children, *child)
	child.Parent = node
}

// game is over when the button player did an action but raise
func (node *Node) isTerminal() bool {
	return node.State.isTerminal()
}

func (node *Node) isExpanded() bool {
	// if the node is expanded, the node should have the parent
	return reflect.DeepEqual(node.Parent, Node{})
}

func (node *Node) getReward() int {
	// if the node is expanded, the node should have the parent
	return node.State.getReward()
}



