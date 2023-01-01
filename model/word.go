package model

import (
	"encoding/json"
	"strings"
)

type Word struct {
	Value string `json:"word"`
	Clue string `json:"clue"`
}

func ParseJsonIntoWord(jsonFile []byte) []Word {
	var wordList []Word

	json.Unmarshal([]byte(jsonFile), &wordList)

	return wordList
}

func ParseLineIntoWord(line string) Word {
	var word Word

	var list = strings.Split(line, ";")

	word.Value = list[0]
	word.Clue = list[1]

	return word
}