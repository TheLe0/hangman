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
	var words [16]string
	rand.Seed(time.Now().UnixNano())

	words[0] = "Car"
	words[1] = "Home"
	words[2] = "Hockey"
	words[3] = "Switch"
	words[4] = "Apple"
	words[5] = "Book"
	words[6] = "Wild"
	words[7] = "Canada"
	words[8] = "King"
	words[9] = "Penguin"
	words[10] = "Hurricane"
	words[11] = "Pirate"
	words[12] = "Ray"
	words[13] = "Dolphin"
	words[14] = "Twin"
	words[15] = "Squirrel"

	return words[rand.Intn(5)]
}

func main() {

	var word string
	var input string
	var lifes int
	var play int
	var matches int
	logo := figure.NewColorFigure("Hangman", "", "red", true)
	win := figure.NewColorFigure("You Win!", "", "green", true)
	lost := figure.NewColorFigure("You Died!", "", "red", true)
	var stages [7]string

	stages[0] = `
			+---+
			|   |
			|   0
			|  /|\
			|  / \
			|
		=========
	`

	stages[1] = `
			+---+
			|   |
			|   0
			|  /|\
			|  /
			|
		=========
	`

	stages[2] = `
			+---+
			|   |
			|   0
			|  /|\
			|
			|
		=========
	`

	stages[3] = `
			+---+
			|   |
			|   0
			|  /|
			|
			|
		=========
	`

	stages[4] = `
			+---+
			|   |
			|   0
			|  /
			|
			|
		=========
	`

	stages[5] = `
			+---+
			|   |
			|   0
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

	positions := make(map[int]string)
	letters := make(map[int]string)
	play = 0
	matches = 0

	for pos := range word {

		positions[pos] = ""
	}

	lettersHistoric := ""

	for lifes >= 0 {
		screen.Clear()
		screen.MoveTopLeft()
		logo.Print()

		fmt.Printf(stages[lifes])

		hiddenWord := ""

		for pos, char := range word {

			if positions[pos] != "" {
				hiddenWord += " " + string(char)
			} else {
				hiddenWord += " _"
			}
		}

		fmt.Println(hiddenWord)
		fmt.Println("")

		if matches == len(word) {
			win.Print()
			fmt.Println("")
			break
		}

		fmt.Println("Letters played: " + lettersHistoric)
		fmt.Println("")
		fmt.Println("Digit a letter: ")
		fmt.Println("")
		fmt.Scanln(&input)
		input = strings.ToLower(input)
		play++

		if len(input) == 1 {

			counter := 0

			letters[play] = input

			if lettersHistoric == "" {
				lettersHistoric += input
			} else {
				lettersHistoric += " - " + input
			}

			for pos, char := range word {

				if string(char) == input && positions[pos] == "" {
					positions[pos] = string(char)
					counter++
					matches++
				}
			}

			if counter == 0 {
				lifes--
			}

		} else {
			if input == word {
				win.Print()
				fmt.Println("")
				break
			} else {
				lifes--
			}
		}

	}

	if lifes < 0 {
		lost.Print()
		fmt.Println("")
		fmt.Println("The word was: " + word)
		fmt.Println("")
	}

}
