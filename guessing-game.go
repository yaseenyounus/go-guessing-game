package main

import (
	"flag"
	"fmt"
	"math/rand"
)

// TODO update level flag with logic

func numInput(minNum, maxNum int) int {
	var input int
	for {
		fmt.Print("Guess a number: ")
		_, err := fmt.Scan(&input)
		if err == nil {
			if input < minNum || input > maxNum {
				fmt.Printf("Number doesn't fall in the range of %v - %v. Try again\n", minNum, maxNum)
			} else {
				return input
			}
		} else {
			fmt.Println("That's not an integer. Try again")
		}
	}
}

func playGame(hints *bool, level *int) string {
	var input, counter int

	minNum, maxNum := 1, 10
	goalNum := rand.Intn(maxNum) + 1

	fmt.Printf("Welcome to Go Guessing Game (G3)!\nEnter numbers between 1 and %v\n\n", maxNum)

	for goalNum != input {
		input = numInput(minNum, maxNum)
		counter++

		if *hints {
			if input > goalNum {
				fmt.Println("Nope, that's too big. Try again")
			} else if input < goalNum {
				fmt.Println("Nope, that's too small. Try again")
			}
		}
	}

	var guessString string
	if counter == 1 {
		guessString = "guess"
	} else {
		guessString = "guesses"
	}
	return fmt.Sprintf("That's correct, good work! Took you %v %v\n", counter, guessString)

}

func main() {
	hints := flag.Bool("hints", false, "Enables hints while guessing")
	level := flag.Int("level", 1, "Chooses a level - 1 (Easy), 2 (Medium), 3 (Hard)")
	flag.Parse()

	fmt.Println(playGame(hints, level))
}
