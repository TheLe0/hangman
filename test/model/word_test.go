package model

import (
	"github.com/TheLe0/hangman/model"
)

import "testing"

func TestParseLineIntoWord(t *testing.T) {

	word := model.ParseLineIntoWord("Penguin;Birds")

	if word.Value != "Penguin" {
		t.Errorf("Expected value to be 'Penguin', but got '%s'", word.Value)
	}

	if word.Clue != "Birds" {
		t.Errorf("Expected clue to be 'Birds', but got '%s'", word.Clue)
	}
}