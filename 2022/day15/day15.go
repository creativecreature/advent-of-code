package day15

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var pattern = regexp.MustCompile("Sensor at x=(?P<sX>-?[0-9]+), y=(?P<sY>-?[0-9]+): closest beacon is at x=(?P<bX>-?[0-9]+), y=(?P<bY>-?[0-9]+)")

func minOf(vars ...int) int {
	min := vars[0]
	for _, i := range vars {
		if min > i {
			min = i
		}
	}
	return min
}

func maxOf(vars ...int) int {
	max := vars[0]
	for _, i := range vars {
		if i > max {
			max = i
		}
	}
	return max
}

func printGrid(grid [][]string) {
	var str string
	for i := range grid {
		str += strings.Join(grid[i], "")
		str += "\n"
	}
	fmt.Println(str)
}

func drawGrid(input io.Reader) [][]string {
	scanner := bufio.NewScanner(input)
	positions := make([][][]int, 0)
	for scanner.Scan() {
		matches := pattern.FindStringSubmatch(scanner.Text())
		sX, sY := matches[pattern.SubexpIndex("sX")], matches[pattern.SubexpIndex("sY")]
		bX, bY := matches[pattern.SubexpIndex("bX")], matches[pattern.SubexpIndex("bY")]
		sensorX, _ := strconv.Atoi(sX)
		sensorY, _ := strconv.Atoi(sY)
		beaconX, _ := strconv.Atoi(bX)
		beaconY, _ := strconv.Atoi(bY)

		sensorPos := []int{sensorX, sensorY}
		beaconPos := []int{beaconX, beaconY}
		pos := [][]int{sensorPos, beaconPos}
		positions = append(positions, pos)
	}

	// We need to find the min/max to normalize the values for the grid
	var minX, minY, maxX, maxY int
	for _, row := range positions {
		sensorX, sensorY := row[0][0], row[0][1]
		beaconX, beaconY := row[1][0], row[1][1]
		minX, minY = minOf(minX, sensorX, beaconX), minOf(minY, sensorY, beaconY)
		maxX, maxY = maxOf(maxX, sensorX, beaconX), maxOf(maxY, sensorY, beaconY)
	}

	// Create a grid with the correct height and width
	// and fill it with "." to start with.
	width := (maxX - minX) + 1
	height := (maxY - minY) + 1
	grid := make([][]string, height)
	for i := 0; i < height; i++ {
		row := make([]string, width)
		for i := range row {
			row[i] = "."
		}
		grid[i] = row
	}

	// Normalize the values for our grid size
	for _, row := range positions {
		row[0][0] = int(float64(row[0][0]) + math.Abs(float64(minX)))
		row[0][1] = row[0][1] - minY
		row[1][0] = int(float64(row[1][0]) + math.Abs(float64(minX)))
		row[1][1] = row[1][1] - minY
	}

	// Mark out all the sensors and beacons
	for _, row := range positions {
		sensorX, sensorY := row[0][0], row[0][1]
		beaconX, beaconY := row[1][0], row[1][1]
		grid[sensorY-minY][sensorX] = "S"
		grid[beaconY-minY][beaconX] = "B"
	}

	for _, row := range positions {
		sensorX, sensorY := row[0][0], row[0][1]
		beaconX, beaconY := row[1][0], row[1][1]
		minX, minY = minOf(sensorX, beaconX), minOf(sensorY, beaconY)
		maxX, maxY = maxOf(sensorX, beaconX), maxOf(sensorY, beaconY)
		positionsToCover := (maxX - minX) + (maxY - minY)

		may := maxOf(sensorY-positionsToCover, 0)
		miy := minOf(sensorY+positionsToCover, len(grid)-1)
		padding := 0
		if sensorY-positionsToCover < 0 {
			padding = int(math.Abs(float64(sensorY) - float64(positionsToCover)))
		}
		// Draw from top down
		for i := may; i <= sensorY; i++ {
			if grid[i][sensorX] != "S" && grid[i][sensorX] != "B" {
				grid[i][sensorX] = "#"
			}
			// Draw left
			minX = maxOf(positionsToCover-(i+padding), 0)
			for n := sensorX; n >= minX; n-- {
				if grid[i][n] != "S" && grid[i][n] != "B" {
					grid[i][n] = "#"
				}
			}

			// Draw right
			minX = minOf(sensorX+i+padding, len(grid[i])-1)
			for n := sensorX; n <= minX; n++ {
				if grid[i][n] != "S" && grid[i][n] != "B" {
					grid[i][n] = "#"
				}
			}
		}

		// Draw bottom up
		for i := miy; i >= sensorY; i-- {
			wuu := positionsToCover - (i - sensorY)
			if grid[i][sensorX] != "S" && grid[i][sensorX] != "B" {
				grid[i][sensorX] = "#"
			}
			// Draw left
			minX = maxOf(sensorX-(wuu), 0)
			for n := sensorX; n >= minX; n-- {
				if grid[i][n] != "S" && grid[i][n] != "B" {
					grid[i][n] = "#"
				}
			}
			// Draw right
			minX = minOf(sensorX+wuu, len(grid[i])-1)
			for n := sensorX; n <= minX; n++ {
				if grid[i][n] != "S" && grid[i][n] != "B" {
					grid[i][n] = "#"
				}
			}
		}
	}
	return grid
}

func PartOne(input io.Reader) int {
	grid := drawGrid(input)
	count := 0
	for i := range grid[20] {
		if grid[20][i] == "#" {
			count++
		}
	}
	return count
}
