package main

import (
	"flag"
	"fmt"
	"math/rand"
)

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

func playGame(goalNum, minNum, maxNum int, hints *bool) int {
	var input, counter int

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
	return counter
}

func main() {
	MIN_NUM := 1
	MAX_NUM := 10

	goalNum := rand.Intn(MAX_NUM) + 1
	// fmt.Println(goalNum)

	fmt.Printf("Welcome to Guessing Game!\nEnter numbers between 1 and %v\n\n", MAX_NUM)

	hints := flag.Bool("hints", false, "Enables hints while guessing")
	flag.Parse()

	counter := playGame(goalNum, MIN_NUM, MAX_NUM, hints)

	var guessString string
	if counter == 1 {
		guessString = "guess"
	} else {
		guessString = "guesses"
	}

	fmt.Printf("That's correct, good work! Took you %v %v\n", counter, guessString)
}
