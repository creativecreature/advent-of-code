package day17

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
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

func createKey(x, y int) string {
	return fmt.Sprintf("%d-%d", x, y)
}

func parseKey(s string) []int {
	keys := strings.Split(s, "-")
	x, err := strconv.Atoi(keys[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(keys[1])
	if err != nil {
		panic(err)
	}
	return []int{x, y}
}

func checkCollision(blocks [][]int, block [][]int) bool {
	blockPositions := make(map[string]bool)
	for _, b := range block {
		key := createKey(b[0], b[1])
		blockPositions[key] = true
	}

	stop := int(math.Max(float64(len(blocks)-300), 0))
	for i := len(blocks) - 1; i >= stop; i-- {
		key := createKey(blocks[i][0], blocks[i][1])
		if blockPositions[key] {
			return true
		}
	}

	return false
}

func getNewTop(blocks [][]int) int {
	max := 0
	for i := range blocks {
		if blocks[i][1] > max {
			max = blocks[i][1]
		}
	}
	return max
}

func PartOne(input io.Reader) int {
	var parsedInput string
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		parsedInput = scanner.Text()
	}

	// Push the floor into the "blocks" map
	blocks := [][]int{}

	// gamestateHashes := make(map[string]bool)
	for i := 0; i < 7; i++ {
		blocks = append(blocks, []int{i, 0})
	}

	cache := map[[2]int][2]int{}
	height, windIndex := 0, 0
	for t := 0; t < 1000000000000-t; t++ {
		// Part one:
		// if t == 2022 {
		// return heigh
		// }
		blockIndex := t % 5
		block := getBlock(blockIndex, height+4)
		k := [2]int{t % 5, windIndex}
		if c, ok := cache[k]; ok {
			if n, d := 1000000000000-t, t-c[0]; n%d == 0 {
				return height + n/d*(height-c[1])
			}
		}
		cache[k] = [2]int{t, height}

		for {
			wind := parsedInput[windIndex]
			// gamestateHashes[hash] = true
			if wind == '<' {
				moveLeft(block)
				if checkCollision(blocks, block) {
					moveRight(block)
				}
			} else {
				moveRight(block)
				if checkCollision(blocks, block) {
					moveLeft(block)
				}
			}
			windIndex = (windIndex + 1) % len(parsedInput)
			moveDown(block)
			if checkCollision(blocks, block) {
				moveUp(block)
				break
			}
		}
		blocks = append(blocks, block...)
		for _, b := range block {
			if b[1] > height {
				height = b[1]
			}
		}
	}

	return height
}
