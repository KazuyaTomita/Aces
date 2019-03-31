package main

type State struct {
	// position index from button. ex) SB is 1, BB is 2.
	Position int
	ActionHistory ActionHistory
	Hand Hand
	Board Board
}

type Board struct {
	Cards [5]Card
}

type Hand struct {
	Cards [2]Card
}

type Rank int

const (
	Two Rank = iota
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

type Suit int

const (
	Heart Suit = iota
	Spade
	Dia
	Club
)

type Card struct {
	Rank Rank
	Suit Suit
}

type ActualAction int

const (
	Raise ActualAction = iota
	Call // we treated call and check as same.
	Fold
)

type Action struct {
	// bet size if ActualAction is Raise
	BetSize int
	Action ActualAction
}

type ActionHistory struct {
	// Action History. ex) 'rrc' means raise, raise and call.
	PreflopActionHistory []Action
	FlopActionHistory []Action
	TurnActionHistory []Action
	RiverActionHistory []Action
}

// make Node from State
func (s *State) toNode() Node {
	return Node{
		State : s,
	}
}

// return a next state after a specific action
func (s *State) getStateAfter(a *Action) State {
	return State{}
}

func (s *State) getReward() int {
	return 1
}

func (s *State) isTerminal() bool {
	return true
}




