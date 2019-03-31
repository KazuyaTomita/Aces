/*
There are some search methods.
We mainly implement Monte Carlo Tree Search(MCTS), following A Survey of Monte Carlo Tree Search Methods.
See http://mcts.ai/pubs/mcts-survey-master.pdf for details.
 */
package main

const Ucb1Coeficient = 0.5

type Search struct {
}

func (s *Search) UCB1Search(state *State) Action {
	node := state.toNode()
	// how many searches are executed is option
	// it will be different in both learning and playing case
	for i := 0; i < 1600; i++ {
		nextNode := treePolicy(&node)
		reward := defaultPolicy(node.State)
		backup(reward, &nextNode)
	}
	bestChild := bestChild(&node, 0)
	return node.getActionFor(&bestChild)
}

// go into tree, following UCB1 formula
// just go into already an expanded-tree or expand a new node.
func treePolicy(node *Node) Node {
	for !node.isTerminal() {
		if !node.isExpanded() {
			return expand(node)
		}
		newNode := bestChild(node, Ucb1Coeficient)
		node = &newNode
	}
	return *node
}

/*
expand the game tree.
Namely, a new node is added to the game tree, following a certain action distribution when choosing an action.
 */
func expand(node *Node) Node {
	action := getAction(node.State)
	nextNode := node.getNextNodeAfter(&action)
	return nextNode
}

/*
do play out and get a reward
 */
func defaultPolicy(state *State) int {
	// do play out
	for !state.isTerminal() {
		action := getAction(state)
		nextState := state.getStateAfter(&action)
		state = &nextState
	}
	return state.getReward()
}

func backup(reward int, node *Node) {
	// backpropagation reward and something like that
}

func bestChild(node *Node, c int) Node {
	// argmax{(Q/N) + c√{(2lnN)/N}}
	// N is the node's visited count and Q is a total of rewards.
	return Node{}
}

func getAction(state *State) Action {
	/*
	Original MCTS chooses an action randomly.
	However, the randomness is not so important, we can limit actions in fact.
	AlphaZero and its derivatives uses a trained neural network for choosing an action.
	*/
	return Action{}
}
