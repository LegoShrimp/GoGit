package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

const (
	ROCK = iota
	SCISSORS
	PAPER
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Chose either (r)ock, (p)aper or (s)cissors: ")
		input, _ := reader.ReadString('\n')
		if human := Convert(string(input[0])); human >= 0 {
			computer := rand.Intn(3)
			switch Wins(human, computer) {
			case -1:
				fmt.Println("Sorry, the computer wins")
			case 0:
				fmt.Println("Draw, try again")
			case 1:
				fmt.Println("YES! You win")
			}
		}
	}
}

// Takes a string and converts it to one of ROCK, PAPER or SCISSORS
// Valid strings are both capital and lowercase:
// * r - ROCK
// * s - SCISSORS
// * p - PAPER
// Return -1 for non-valid string
func Convert(s string) int {
	if s == "R" || s == "r" {
		return 0
	}
	if s == "S" || s == "s" {
		return 1
	}
	if s == "P" || s == "p" {
		return 2
	}

	return -1
}

// Takes two turns and returns
// * -1 for win of player2
// * 0 for draw
// * 1 for win of player 1
// Only valid turns are given in the argument
func Wins(player1, player2 int) int {
	if player1 == player2 {
		return 0
	}
	if (player1 == 0 && player2 == 1) || (player1 == 1 && player2 == 2) || (player1 == 2 && player2 == 0) {
		return 1
	}
	return -1
	return 0
}
