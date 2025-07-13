package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// Generate a random number between 1 and 10
	target := random.Intn(100) + 1

	// welcome message
	fmt.Println("Welcome to the guessing game!")
	fmt.Println("I have chosen a number between 1 and 100")
	fmt.Println("Can you guess what it is?")

	var guess int

	for {
		fmt.Println("Enter your number: ")
		fmt.Scanln(&guess) // providing a pointer updates the actual variable instead of a copy

		// check if the guess if correct
		if guess == target {
			fmt.Println("Congrats! You guessed the number.")
			break
		} else if guess < target {
			fmt.Println("Too low! Try guessing a higher number.")
		} else {
			fmt.Println("Too high! Try guessing a lower number.")
		}

	}

}
