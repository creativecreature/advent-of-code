package day09

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

type movement struct {
	direction string
	steps     int
}

func parseInput(input io.Reader) []movement {
	movements := []movement{}
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		m := strings.Split(scanner.Text(), " ")
		steps, err := strconv.Atoi(m[1])
		if err != nil {
			log.Fatal(err)
		}
		movements = append(movements, movement{
			direction: m[0],
			steps:     steps,
		})
	}

	return movements
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func areTouching(headX, headY, tailX, tailY int) bool {
	return diff(headX, tailX) <= 1 && diff(headY, tailY) <= 1
}

func moveKnot(headX, headY, tailX, tailY int) (int, int) {
	// First check if they are on the same row
	if headY == tailY {
		if headX > tailX {
			return tailX + 1, tailY
		} else {
			return tailX - 1, tailY
		}
	}
	// Check if they are in the same column
	if headX == tailX {
		if headY > tailY {
			return tailX, tailY + 1
		} else {
			return tailX, tailY - 1
		}
	}
	// If not we need to step diagonally
	var newX, newY int
	if headX > tailX {
		newX = tailX + 1
	} else {
		newX = tailX - 1
	}
	if headY > tailY {
		newY = tailY + 1
	} else {
		newY = tailY - 1
	}
	return newX, newY
}

func getTailMovements(numberOfKnots int, movements []movement) int {
	knots := make([][]int, numberOfKnots)
	for i := 0; i < numberOfKnots; i++ {
		knots[i] = make([]int, 2)
	}
	visistedPositions := map[string]bool{
		"0-0": true,
	}

	for _, m := range movements {
		for i := 0; i < m.steps; i++ {
			switch m.direction {
			case "R":
				knots[0][0]++
			case "L":
				knots[0][0]--
			case "U":
				knots[0][1]++
			case "D":
				knots[0][1]--
			default:
				panic("Unknown direction")
			}
			for i := 1; i < len(knots); i++ {
				if areTouching(knots[i-1][0], knots[i-1][1], knots[i][0], knots[i][1]) {
					continue
				}
				knots[i][0], knots[i][1] = moveKnot(knots[i-1][0], knots[i-1][1], knots[i][0], knots[i][1])
				if i == len(knots)-1 {
					key := fmt.Sprintf("%d-%d", knots[i][0], knots[i][1])
					visistedPositions[key] = true
				}
			}
		}
	}

	return len(visistedPositions)
}

func PartOne(input io.Reader) int {
	movements := parseInput(input)
	return getTailMovements(2, movements)
}

func PartTwo(input io.Reader) int {
	movements := parseInput(input)
	return getTailMovements(10, movements)
}
