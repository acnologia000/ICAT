package main

import (
	"log"
	"strings"
)

func moveYoshita(instructions string, pos []int) []int {
	if len(pos) != 2 {
		log.Fatal("please give correct number of coordinates")
	}

	instructionList := strings.Split(strings.ToUpper(instructions), "")

	for _, instruction := range instructionList {
		switch instruction {
		case "L":
			pos[0] = pos[0] - 1
		case "R":
			pos[0] = pos[0] + 1
		case "U":
			pos[1] = pos[1] + 1
		case "D":
			pos[1] = pos[1] - 1
		default:
			log.Fatal("faulty instruction, please use L, R, D, U only")
		}
	}

	return pos
}
