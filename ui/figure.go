package ui

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
)

func GameTitle() {
	logo := figure.NewColorFigure("Hangman", "", "red", true)
	logo.Print()
}

func GameWon() {
	win := figure.NewColorFigure("You Win!", "", "green", true)
	win.Print()
	fmt.Println("")
}

func GameLost(word string) {
	lost := figure.NewColorFigure("You Died!", "", "red", true)
	lost.Print()
	fmt.Println("")
	fmt.Println("The word was: " + word)
	fmt.Println("")
}