package day05

import (
	"bufio"
	"io"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var pattern = regexp.MustCompile("move ([0-9][0-9]?) from ([0-9]) to ([0-9])")

func moveContainers(stacks [][]string, crane func(stacks [][]string, amount, from, to int), input io.Reader) [][]string {
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		match := pattern.FindStringSubmatch(scanner.Text())
		amount, err := strconv.Atoi(match[1])
		if err != nil {
			log.Fatal(err)
		}

		from, err := strconv.Atoi(match[2])
		if err != nil {
			log.Fatal(err)
		}

		to, err := strconv.Atoi(match[3])
		if err != nil {
			log.Fatal(err)
		}

		// Make indexes start from 0
		from = from - 1
		to = to - 1
		crane(stacks, amount, from, to)
	}
	return stacks
}

func getTopContainers(stacks [][]string) string {
	topOfStacks := []string{}
	for i := range stacks {
		topOfStacks = append(topOfStacks, stacks[i][len(stacks[i])-1])
	}
	return strings.Join(topOfStacks, "")
}

func PartOne(stacks [][]string, input io.Reader) string {
	crane := func(stacks [][]string, amount, from, to int) {
		for i := 1; i <= amount; i++ {
			toMove := stacks[from][len(stacks[from])-1]
			stacks[to] = append(stacks[to], toMove)
			stacks[from] = stacks[from][:len(stacks[from])-1]
		}
	}
	moveContainers(stacks, crane, input)
	return getTopContainers(stacks)
}

func PartTwo(stacks [][]string, input io.Reader) string {
	crane := func(stacks [][]string, amount, from, to int) {
		stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-amount:]...)
		stacks[from] = stacks[from][0 : len(stacks[from])-amount]
	}
	moveContainers(stacks, crane, input)
	return getTopContainers(stacks)
}
