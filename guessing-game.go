package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func numInput() int {
	var input int
	for {
		fmt.Print("Guess a number: ")
		_, err := fmt.Scan(&input)
		if err == nil {
			return input
		}
		fmt.Println("That's not an integer. Try again")
	}
}

func playGame(goalNum, maxNum int) int {
	var input int
	counter := 0

	for goalNum != input {
		input = numInput()
		counter++
	}
	return counter
}

func playGameWithHints(goalNum, maxNum int) int {
	var input int
	counter := 0

	for goalNum != input {
		input = numInput()
		counter++

		if input > goalNum {
			fmt.Println("Nope, that's too big. Try again")
		} else if input < goalNum {
			fmt.Println("Nope, that's too small. Try again")
		}
	}
	return counter
}

// TODO


func main() {
	MAX_NUM := 10

	goalNum := rand.Intn(MAX_NUM) + 1
	// fmt.Println(goalNum)

	fmt.Printf("Welcome to Guessing Game!\nEnter numbers between 1 and %v\n\n", MAX_NUM)

	hints := flag.Bool("hints", false, "Enables hints while guessing")
	flag.Parse()

	var counter int
	if *hints {
		fmt.Println("Playing with hints")
		counter = playGameWithHints(goalNum, MAX_NUM)
	} else {
		counter = playGame(goalNum, MAX_NUM)

	}

	var guessString string
	if counter == 1 {
		guessString = "guess"
	} else {
		guessString = "guesses"
	}

	fmt.Printf("That's correct, good work! Took you %v %v\n", counter, guessString)
}
