package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("I picked a number between 1 and 100")
	fmt.Println("Could you guess it?")
	fmt.Println("But you have 7 guesses ")
	rand.Seed(time.Now().UnixNano())
	pickedNum := rand.Intn(100) + 1
	lives := 0
	guessedNum := 0

	for lives != 7 {
		fmt.Println("Choose a number between 1 and 100: ")
		fmt.Scan(&guessedNum)
		if guessedNum > 100 || guessedNum < 0 {
			fmt.Println("Invalid interval you can just guess between 1 and 100")
		} else {
			if pickedNum == guessedNum {
				fmt.Println("You won the game!")
				break
			} else {
				lives++
				if pickedNum > guessedNum {
					fmt.Printf("You have %d guesses left\n", 7-lives)
					fmt.Println("Wrong , guessed Higher")
				}
				if pickedNum < guessedNum {
					fmt.Printf("You have %d guesses left\n", 7-lives)
					fmt.Println("Wrong , guessed Lower")
				}
			}
			if lives == 7 {
				fmt.Println("You've run out of guessing rights")
				fmt.Printf("Picked Number was %d", pickedNum)
			}

		}

	}

}
