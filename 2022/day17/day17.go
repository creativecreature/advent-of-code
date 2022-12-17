package day17

import (
	"fmt"
	"io"
)

func getBlock(blockType, y int) [][]int {
	switch blockType {
	case 0:
		return [][]int{{2, y}, {3, y}, {4, y}, {5, y}}
	case 1:
		return [][]int{{3, y + 2}, {2, y + 1}, {3, y + 1}, {4, y + 1}, {3, y}}
	case 2:
		return [][]int{{2, y}, {3, y}, {4, y}, {4, y + 1}, {4, y + 2}}
	case 3:
		return [][]int{{2, y}, {2, y + 1}, {2, y + 2}, {2, y + 3}}
	case 4:
		return [][]int{{2, y + 1}, {2, y}, {3, y + 1}, {3, y}}
	default:
		panic("Unknown block type")
	}
}

func moveLeft(block [][]int) [][]int {
	canMove := true
	for _, pos := range block {
		xPos := pos[0]
		if xPos == 0 {
			canMove = false
			break
		}
	}

	if !canMove {
		return block
	}

	for _, pos := range block {
		pos[0] = pos[0] - 1
	}
	return block
}

func moveRight(block [][]int) [][]int {
	canMove := true
	for _, pos := range block {
		xPos := pos[0]
		if xPos == 6 {
			canMove = false
			break
		}
	}

	if !canMove {
		return block
	}

	for _, pos := range block {
		pos[0] = pos[0] + 1
	}
	return block
}

func moveDown(block [][]int) [][]int {
	for _, pos := range block {
		pos[1] = pos[1] - 1
	}
	return block
}

func moveUp(block [][]int) [][]int {
	for _, pos := range block {
		pos[1] = pos[1] + 1
	}
	return block
}

var parsedInput = ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"

func createKey(x, y int) string {
	return fmt.Sprintf("%d-%d", x, y)
}

func PartOne(input io.Reader) int {
	// Push the floor into the "seen" map
	seen := make(map[string]bool)
	for i := 0; i < 7; i++ {
		key := createKey(i, 0)
		seen[key] = true
	}

	for t := 1; t <= 2022; t++ {
		block := getBlock(t%5, 1)
	}
	return -1
}
