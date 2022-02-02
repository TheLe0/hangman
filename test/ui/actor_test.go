package ui

import (
	"github.com/TheLe0/hangman/ui"
)

import "testing"

func Test0LifesRemaining(t *testing.T){

	want :=  `
			+---+
			|   |
			|   0
			|  /|\
			|  / \
			|
		=========
	`

	got := ui.RemainLifes(0)

	if got != want {
        t.Errorf("0 lifes do not match")
    }
}

func Test1LifesRemaining(t *testing.T){

	want := `
			+---+
			|   |
			|   0
			|  /|\
			|  /
			|
		=========
	`

	got := ui.RemainLifes(1)

	if got != want {
        t.Errorf("1 lifes do not match")
    }
}

func Test2LifesRemaining(t *testing.T){


	want := `
			+---+
			|   |
			|   0
			|  /|\
			|
			|
		=========
	`

	got := ui.RemainLifes(2)

	if got != want {
        t.Errorf("2 lifes do not match")
    }
}

func Test3LifesRemaining(t *testing.T){

	want :=  `
			+---+
			|   |
			|   0
			|  /|
			|
			|
		=========
	`

	got := ui.RemainLifes(3)

	if got != want {
        t.Errorf("3 lifes do not match")
    }
}

func Test4LifesRemaining(t *testing.T){

	want :=   `
			+---+
			|   |
			|   0
			|  /
			|
			|
		=========
	`

	got := ui.RemainLifes(4)

	if got != want {
        t.Errorf("4 lifes do not match")
    }
}

func Test5LifesRemaining(t *testing.T){

	want :=  `
			+---+
			|   |
			|   0
			|
			|
			|
		=========
	`

	got := ui.RemainLifes(5)

	if got != want {
        t.Errorf("5 lifes do not match")
    }
}

func Test6LifesRemaining(t *testing.T){

	want :=  `
			+---+
			|   |
			|    
			|
			|
			|
		=========
	`

	got := ui.RemainLifes(6)

	if got != want {
        t.Errorf("6 lifes do not match")
    }
}
