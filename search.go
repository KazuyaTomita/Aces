/*
There are some search methods.
We mainly implement Monte Carlo Tree Search(MCTS), following A Survey of Monte Carlo Tree Search Methods.
See http://mcts.ai/pubs/mcts-survey-master.pdf for details.
 */
package main

type Search struct {
	// position index from button. ex) SB is 1, BB is 2.
	RootNode Node
}

type Node struct {
	State State
	Children []Node
	Parent Node
}

func (node *Node) getActionFor(nextNode *Node) Action {
	return Action{}
}

func (s *Search) UCB1Search(state *State) Action {
	node := state.toNode()
	for i := 0; i < 1600; i++ {
		// iの値が0から9まで変化していく
		nextNode := treePolicy(&node)
		reward := defaultPolicy(&node)
		backup(reward, &nextNode)
	}
	bestChild := bestChild(&node, 0)
	return node.getActionFor(&bestChild)
}

func treePolicy(node *Node) Node {
	// go into tree following UCB1 formula
	// just go into already an expanded-tree or expand a new node.
	return Node{}
}

func defaultPolicy(node *Node) int {
	// do play out
	return 1
}

func backup(reward int, node *Node) {
	// backpropagation
}

func bestChild(node *Node, c int) Node {
	return Node{}
}
