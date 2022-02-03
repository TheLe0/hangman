package model

import (
	"github.com/TheLe0/hangman/model"
)

import "testing"

func TestCreatePlayerStruct(t *testing.T) {

	player := model.Player {
		Lifes: 6,
	}

	if player.Lifes != 6 {
		t.Errorf("Expected value to be '6', but got '%d'", player.Lifes)
	}

}