package model

import (
	"github.com/TheLe0/hangman/model"
)

import "testing"

func TestCreateGameStruct(t *testing.T) {

	game := model.Game {
		Word: model.Word {
			Value: "Penguin",
			Clue: "Birds",
		},
		IsFinished: false,
	}

	if game.IsFinished != false {
		t.Errorf("Expected value to be 'false', but got '%t'", game.IsFinished)
	}

	if game.Word.Value != "Penguin" {
		t.Errorf("Expected value to be 'Penguin', but got '%s'", game.Word.Value)
	}

	if game.Word.Clue != "Birds" {
		t.Errorf("Expected value to be 'Birds', but got '%s'", game.Word.Clue)
	}

}