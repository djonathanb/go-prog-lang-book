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

	secret_number := rand.Intn(100)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Please input the guess: ")

		guess_string, _ := reader.ReadString('\n')
		trimmed_guess := strings.TrimSuffix(guess_string, "\n")
		guess, err := strconv.Atoi(trimmed_guess)

		if err != nil {
			fmt.Println("Invalid number!")
			continue
		}

		fmt.Printf("You guessed: %d\n", guess)

		if guess < secret_number {
			fmt.Println("Too small!")
		} else if guess > secret_number {
			fmt.Println("Too big!")
		} else {
			fmt.Println("You win!")
			break
		}
	}
}
