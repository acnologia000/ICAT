package main

import "strings"

const (
	capitalLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	smallLetters   = "abcdefghijklmnopqrtwxyz"
)

var letterList = strings.Split(capitalLetters, "")

func encrypt(input string) string {
	finalString := strings.Builder{}

	letters := strings.Split(input, "")

	for _, letter := range letters {
		if strings.Contains(capitalLetters, letter) {
			letter = shiftLetter(letter, 4)
		}

		if strings.Contains(smallLetters, letter) {
			letter = strings.ToUpper(letter)
			letter = shiftLetter(letter, 4)
			letter = strings.ToLower(letter)
		}
		finalString.WriteString(letter)
	}
	return finalString.String()
}

func shiftLetter(letter string, magnitude int) string {
	index := strings.Index(capitalLetters, letter)

	if magnitude+index <= 25 {
		return letterList[magnitude+index]
	} else {
		index = (index + magnitude) - 25
		return letterList[index-1]
	}
}
