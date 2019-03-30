/*
There are some search methods.
We mainly implement Monte Carlo Tree Search(MCTS), following A Survey of Monte Carlo Tree Search Methods.
See http://mcts.ai/pubs/mcts-survey-master.pdf for details.
 */
package main

const Ucb1Coeficient = 0.5

type Search struct {
	// position index from button. ex) SB is 1, BB is 2.
	RootNode Node
}

func (s *Search) UCB1Search(state *State) Action {
	node := state.toNode()
	// how many searches are executed is option
	// it will be different in both learning and playing case
	for i := 0; i < 1600; i++ {
		nextNode := treePolicy(&node)
		reward := defaultPolicy(&node)
		backup(reward, &nextNode)
	}
	bestChild := bestChild(&node, 0)
	return node.getActionFor(&bestChild)
}

func treePolicy(node *Node) Node {
	for !node.isTerminal() {
		if !node.isExpanded() {
			return expand(node)
		}
		newNode := bestChild(node, Ucb1Coeficient)
		node = &newNode
	}
	// go into tree following UCB1 formula
	// just go into already an expanded-tree or expand a new node.
	return *node
}

func expand(node *Node) Node {
	// return next node
	return Node{}
}


func defaultPolicy(node *Node) int {
	// do play out
	return 1
}

func backup(reward int, node *Node) {
	// backpropagation reward and something like that
}

func bestChild(node *Node, c int) Node {
	return Node{}
}
