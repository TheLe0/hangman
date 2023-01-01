package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"

	"github.com/TheLe0/hangman/ui"
	"github.com/TheLe0/hangman/model"
)

func randomWord() model.Word {
	rand.Seed(time.Now().UnixNano())

	plan, _ := ioutil.ReadFile("./data/words.json")

	words := model.ParseJsonIntoWord(plan)

	return words[rand.Intn(30)]
}

func main() {

	var word string
	var clue string
	var input string
	var lifes int
	var play int
	var matches int

	lifes = 6
	wordObj := randomWord()
	word = strings.ToLower(wordObj.Value)
	clue = strings.ToLower(wordObj.Clue)

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

		fmt.Println("Clue: " + clue)
		fmt.Println("")
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
