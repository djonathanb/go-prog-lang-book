package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Guess the number!")

	secretNumber := rand.Intn(100)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Please input the guess: ")

		guessString, _ := reader.ReadString('\n')
		trimmedGuess := strings.TrimSuffix(guessString, "\n")
		guess, err := strconv.Atoi(trimmedGuess)

		if err != nil {
			fmt.Println("Invalid number!")
			continue
		}

		fmt.Printf("You guessed: %d\n", guess)

		if guess < secretNumber {
			fmt.Println("Too small!")
		} else if guess > secretNumber {
			fmt.Println("Too big!")
		} else {
			fmt.Println("You win!")
			break
		}
	}
}
