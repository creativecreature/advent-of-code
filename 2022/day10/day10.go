package day10

import (
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

func parseInput(reader io.Reader) []string {
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(bytes), "\n")
}

func PartOneTwo(reader io.Reader) int {
	input := parseInput(reader)

	x, cycle, i, total := 1, 1, 0, 0
	commands := make([]int, 0)
	for i < len(input) || len(commands) > 0 {
		crtPosition := (cycle - 1) % 40
		if crtPosition == 0 {
			fmt.Printf("\n")
		}
		if crtPosition >= x-1 && crtPosition <= x+1 {
			fmt.Printf("#")
		} else {
			fmt.Printf(".")
		}
		switch cycle {
		case 20, 60, 100, 140, 180, 220:
			total += x * cycle
		}

		if len(commands) > 0 {
			x += commands[0]
			commands = commands[1:]
		} else {
			line := strings.Split(input[i], " ")
			if line[0] == "addx" {
				num, err := strconv.Atoi(line[1])
				if err != nil {
					panic(err)
				}
				commands = append(commands, num)
			}
			i++
		}
		cycle++
	}
	return total
}
