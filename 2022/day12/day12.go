package day12

import (
	"bufio"
	"fmt"
	"io"
)

const (
	start = 83
	end   = 69
)

func parseInput(input io.Reader) [][]rune {
	runes := [][]rune{}
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		currentRunes := []rune{}
		for _, r := range scanner.Text() {
			currentRunes = append(currentRunes, r)
		}
		runes = append(runes, currentRunes)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return runes
}

func getStartPosition(matrix [][]rune) []int {
	for i := 0; i < len(matrix); i++ {
		for y := 0; y < len(matrix[i]); y++ {
			if matrix[i][y] == start {
				matrix[i][y] = 'a' // After we've found it we can treat it as any other a
				return []int{i, y}
			}
		}
	}
	panic("Could not locate start position")
}

func getEndPosition(matrix [][]rune) []int {
	for i := 0; i < len(matrix); i++ {
		for y := 0; y < len(matrix[i]); y++ {
			if matrix[i][y] == end {
				matrix[i][y] = 'z' // After we've found it we can treat it as z
				return []int{i, y}
			}
		}
	}
	panic("Could not locate end position")
}

func canMoveAsc(x, y rune) bool {
	if y == end {
		y = 'z'
	}
	return y <= x+1
}

func canMoveDesc(x, y rune) bool {
	if y == start {
		y = 'a'
	}
	return x <= y+1
}

func createKey(x, y int) string {
	return fmt.Sprintf("%d-%d", x, y)
}

func getNextNodes(matrix [][]rune, currentPosition []int, canMove func(x, y rune) bool) [][]int {
	x, y := currentPosition[0], currentPosition[1]
	validSiblings := [][]int{}
	// Add left
	if y > 0 && canMove(matrix[x][y], matrix[x][y-1]) {
		validSiblings = append(validSiblings, []int{x, y - 1})
	}
	// Add right
	if y < len(matrix[x])-1 && canMove(matrix[x][y], matrix[x][y+1]) {
		validSiblings = append(validSiblings, []int{x, y + 1})
	}
	// Add top
	if x > 0 && canMove(matrix[x][y], matrix[x-1][y]) {
		validSiblings = append(validSiblings, []int{x - 1, y})
	}
	// Add bottom
	if x < len(matrix)-1 && canMove(matrix[x][y], matrix[x+1][y]) {
		validSiblings = append(validSiblings, []int{x + 1, y})
	}

	return validSiblings
}

func bfs(matrix [][]rune, startPos []int, canMove func(x, y rune) bool, goal rune) int {
	queue := [][]int{startPos}
	visisted := make(map[string]bool)
	steps := make(map[string]int)
	steps[createKey(startPos[0], startPos[1])] = 0

	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]
		key := createKey(pos[0], pos[1])
		if visisted[key] {
			continue
		}

		visisted[key] = true
		count := steps[key]
		if matrix[pos[0]][pos[1]] == goal {
			return count
		}

		nextNodes := getNextNodes(matrix, pos, canMove)
		for _, n := range nextNodes {
			steps[createKey(n[0], n[1])] = count + 1
		}
		queue = append(queue, nextNodes...)
	}

	return -1
}

func PartOne(input io.Reader) int {
	matrix := parseInput(input)
	startPos := getStartPosition(matrix)
	n := bfs(matrix, startPos, canMoveAsc, end)
	return n
}

func PartTwo(input io.Reader) int {
	matrix := parseInput(input)
	startPos := getEndPosition(matrix)
	n := bfs(matrix, startPos, canMoveDesc, 'a')
	return n
}
