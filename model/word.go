package model

import (
	"strings"
)

type Word struct {
	Value string
	Clue string
}

func ParseLineIntoWord(line string) Word {
	var word Word

	var list = strings.Split(line, ";")

	word.Value = list[0]
	word.Clue = list[1]

	return word
}