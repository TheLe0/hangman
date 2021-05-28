package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/common-nighthawk/go-figure"
	"github.com/inancgumus/screen"
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
	var lifes int
	logo := figure.NewColorFigure("Hangman", "", "red", true)
	win := figure.NewColorFigure("You Win!", "", "green", true)
	lost := figure.NewColorFigure("You Died!", "", "red", true)
	var stages [7]string

	stages[0] = `
			+---+
			|   |
			O   |
		   /|\  |
		   / \  |
				|
		=========
	`

	stages[1] = `
			+---+
			|   |
			O   |
		   /|\  |
		   /    |
				|
		=========
	`

	stages[2] = `
			+---+
			|   |
			O   |
		   /|\  |
			    |
				|
		=========
	`

	stages[3] = `
			+---+
			|   |
			O   |
		   /|  	|
				|
				|
		=========
	`

	stages[4] = `
			+---+
			|   |
			O   |
		   /  	|
				|
				|
		=========
	`

	stages[5] = `
			+---+
			|   |
			O   |
				|
				|
				|
		=========
	`

	stages[6] = `
			+---+
			|   |
				|
				|
				|
				|
		=========
	`

	lifes = 6
	word = strings.ToLower(randomWord())

	logo.Print()
	for lifes >= 0 {
		screen.Clear()
		screen.MoveTopLeft()

		fmt.Printf(stages[lifes])
		fmt.Println("Digit a letter: ")
		fmt.Scanln(&input)
		input = strings.ToLower(input)

		if len(input) == 1 {

			for pos, char := range word {
				fmt.Printf("character %c starts at byte position %d\n", char, pos)
			}

			lifes--

		} else {
			if input == word {
				win.Print()
				break
			} else {
				lifes--
			}
		}

	}

	if lifes < 0 {
		lost.Print()
	}

}
