package main

import "fmt"

func main() {
	fmt.Printf("start\n")
	for {
		search()
		choose_action()
		end_game()
	}

}

func search() {
	fmt.Printf("search starts\n")
	// execute search and set next action to global variables
}

func choose_action() {
	fmt.Printf("choose next action from \n")
}

func end_game() {
	fmt.Printf("Dose this game end? \n")
}

