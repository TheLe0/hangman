package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomWord() string {
	var words [5]string
	rand.Seed(time.Now().UnixNano())

	words[0] = "Car"
	words[1] = "Home"
	words[2] = "Hockey"
	words[3] = "Switch"
	words[4] = "Apple"

	return words[rand.Intn(5)]
}

func main() {

	var word string
	var input string

	word = randomWord()

	fmt.Println("Digit a letter: ")
	fmt.Scanln(&input)

	for pos, char := range word {
		fmt.Printf("character %c starts at byte position %d\n", char, pos)
	}
}
