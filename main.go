package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/TheLe0/hangman/ui"
)

func randomWord() string {
	var words [31]string
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
	words[16] = "Barbecue"
	words[17] = "Squirrel"
	words[18] = "Tomato"
	words[19] = "Aesthetic"
	words[20] = "Anarchy"
	words[21] = "Neverland"
	words[22] = "Blasphemous"
	words[23] = "Cookbook"
	words[24] = "Soccer"
	words[25] = "Language"
	words[26] = "Payroll"
	words[27] = "Settlement"
	words[28] = "Retail"
	words[29] = "Acknowledgement"
	words[30] = "Freight"

	return words[rand.Intn(30)]
}

func main() {

	var word string
	var input string
	var lifes int
	var play int
	var matches int

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
		ui.ClearScreen()
		ui.GameTitle()

		fmt.Printf(ui.RemainLifes(lifes))

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
			ui.GameWon()
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
				ui.GameWon()
				break
			} else {
				lifes--
			}
		}

	}

	if lifes < 0 {
		ui.GameLost(word)
	}

}
