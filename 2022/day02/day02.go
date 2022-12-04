package day02

import (
	"bufio"
	"io"
	"log"
	"strings"
)

type GameInput = struct {
	Win  int
	Draw int
	Lose int
}

func PartTwo(reader io.Reader) int {
	score := 0
	gameInputMap := map[string]GameInput{
		"A": {8, 4, 3},
		"B": {9, 5, 1},
		"C": {7, 6, 2},
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		opponentChoice := input[0]
		winLoseDraw := input[1]

		switch winLoseDraw {
		case "Z":
			score = score + gameInputMap[opponentChoice].Win
		case "Y":
			score = score + gameInputMap[opponentChoice].Draw
		default:
			score = score + gameInputMap[opponentChoice].Lose
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return score
}
