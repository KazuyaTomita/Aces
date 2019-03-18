package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
This is poker AI engine and implements UPI (Universal Poker Interface)[https://cdn.shopify.com/s/files/1/0769/9693/files/UPI-documentation_e9050fc5-e6e6-4f37-8611-04819b636333.pdf]
 */
func main() {
	// first we wait command from GUI(CUI) server.
	// After that, we execute something following the command.
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		switch stdin.Text() {
		case "is_ready":
			fmt.Println("is_ready ok!")
		case "exit":
			panic("exit")
		default:
			fmt.Println("other")
		}
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

