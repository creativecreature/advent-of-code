package day17

import (
	"bufio"
	"fmt"
	"io"
	"math"
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

func checkCollision(blocks [][]int, block [][]int) bool {
	blockPositions := make(map[string]bool)
	for _, b := range block {
		key := fmt.Sprintf("%d-%d", b[0], b[1])
		blockPositions[key] = true
	}

	// Compare against last 300 blocks
	stop := int(math.Max(float64(len(blocks)-300), 0))
	for i := len(blocks) - 1; i >= stop; i-- {
		key := fmt.Sprintf("%d-%d", blocks[i][0], blocks[i][1])
		if blockPositions[key] {
			return true
		}
	}

	return false
}

func getHeight(input io.Reader, numOfBlocks int) int {
	var parsedInput string
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		parsedInput = scanner.Text()
	}

	// Push the floor into the "blocks" map
	blocks := [][]int{}
	for i := 0; i < 7; i++ {
		blocks = append(blocks, []int{i, 0})
	}

	cache := map[[2]int][2]int{}
	height, windIndex := 0, 0
	for t := 0; t < numOfBlocks; t++ {
		blockIndex := t % 5
		block := getBlock(blockIndex, height+4)
		k := [2]int{blockIndex, windIndex}
		if c, ok := cache[k]; ok {
			if blocksLeft, d := numOfBlocks-t, t-c[0]; blocksLeft%d == 0 {
				return height + blocksLeft/d*(height-c[1])
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

func PartOne(input io.Reader) int {
	return getHeight(input, 2022)
}

func PartTwo(input io.Reader) int {
	return getHeight(input, 1000000000000)
}
