package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
This is a poker AI engine and implements modified UPI (Universal Poker Interface).
See UPI.md for the protocol's detail.
 */
func main() {
	/*
	First we wait command from GUI(CUI) server.
	After that, we execute something corresponding to the sent command.
	Finally we have to respond with some command.
	Now we expect GUI does exec(2) this engine and reads this process's standard out through pipe.
	So, we just write something to standard out and that means out response.
	 */
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		switch stdin.Text() {
		case "upi":
			fmt.Println("upiok")
		case "isready":
			// do some initialization if necessary
			fmt.Println("is_readyok")
		case "upinewgame":
			preparaForGame()
		case "state":
			setState()
		case "go":
			search()
		case "exit":
			panic("exit")
		case "gameover":
			// do something
		case "stop":
			choose_best_action()
		default:
			fmt.Println("other")
		}
	}
}

func preparaForGame() {
	// do some preparation
}

func setState() {
	// set state
}

func search() {
	// execute search and set next action to global variables
}

func choose_best_action() {
	// choose best action at the time
}
