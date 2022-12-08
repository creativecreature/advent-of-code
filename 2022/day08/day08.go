package day08

import (
	"bufio"
	"io"
	"strconv"
)

func parseInput(input io.Reader) [][]int {
	grid := [][]int{}
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		row := []int{}
		line := scanner.Text()
		for _, rune := range line {
			val, _ := strconv.Atoi(string(rune))
			row = append(row, val)
		}
		grid = append(grid, row)
	}
	return grid
}

func isHiddenLeft(i, y int, grid [][]int) bool {
	for l := y - 1; l >= 0; l-- {
		if grid[i][l] >= grid[i][y] {
			return true
		}
	}
	return false
}

func isHiddenRight(i, y int, grid [][]int) bool {
	for r := y + 1; r < len(grid[i]); r++ {
		if grid[i][r] >= grid[i][y] {
			return true
		}
	}
	return false
}

func isHiddenUp(i, y int, grid [][]int) bool {
	for u := i - 1; u >= 0; u-- {
		if grid[u][y] >= grid[i][y] {
			return true
		}
	}
	return false
}

func isHiddenDown(i, y int, grid [][]int) bool {
	for d := i + 1; d < len(grid); d++ {
		if grid[d][y] >= grid[i][y] {
			return true
		}
	}
	return false
}

func isHiddenInAllDirections(i, y int, grid [][]int) bool {
	if !isHiddenLeft(i, y, grid) {
		return false
	}
	if !isHiddenRight(i, y, grid) {
		return false
	}
	if !isHiddenUp(i, y, grid) {
		return false
	}
	if !isHiddenDown(i, y, grid) {
		return false
	}
	return true
}

func PartOne(input io.Reader) int {
	grid := parseInput(input)
	visibleTrees := len(grid) * len(grid[0])
	for i := 1; i < len(grid)-1; i++ {
		for y := 1; y < len(grid[i])-1; y++ {
			if isHiddenInAllDirections(i, y, grid) {
				visibleTrees--
			}
		}
	}
	return visibleTrees
}

func getScoreLeft(i, y int, grid [][]int) int {
	score := 0
	for l := y - 1; l >= 0; l-- {
		score++
		if grid[i][l] >= grid[i][y] {
			return score
		}
	}
	return score
}

func getScoreRight(i, y int, grid [][]int) int {
	score := 0
	for r := y + 1; r < len(grid[i]); r++ {
		score++
		if grid[i][r] >= grid[i][y] {
			return score
		}
	}
	return score
}

func getScoreUp(i, y int, grid [][]int) int {
	score := 0
	for u := i - 1; u >= 0; u-- {
		score++
		if grid[u][y] >= grid[i][y] {
			return score
		}
	}
	return score
}

func getScoreDown(i, y int, grid [][]int) int {
	score := 0
	for d := i + 1; d < len(grid); d++ {
		score++
		if grid[d][y] >= grid[i][y] {
			return score
		}
	}
	return score
}

func PartTwo(input io.Reader) int {
	grid := parseInput(input)
	bestScore := 0

	for i := 1; i < len(grid)-1; i++ {
		for y := 1; y < len(grid[i])-1; y++ {
			total := getScoreLeft(i, y, grid) *
				getScoreRight(i, y, grid) *
				getScoreUp(i, y, grid) *
				getScoreDown(i, y, grid)

			if total > bestScore {
				bestScore = total
			}
		}
	}
	return bestScore
}
