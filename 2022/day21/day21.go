package day21

import (
	"io"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func parseInput(i io.Reader) map[string]string {
	input, _ := ioutil.ReadAll(i)
	monkeys := map[string]string{}
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		s := strings.Split(s, ": ")
		monkeys[s[0]] = s[1]
	}
	return monkeys
}

func PartOne(i io.Reader) int {
	monkeys := parseInput(i)
	return solve(monkeys, "root")
}

func PartTwo(i io.Reader) int {
	monkeys := parseInput(i)
	monkeys["humn"] = "0"
	s := strings.Fields(monkeys["root"])
	if solve(monkeys, s[0]) < solve(monkeys, s[2]) {
		s[0], s[2] = s[2], s[0]
	}

	part2, _ := sort.Find(1e16, func(v int) int {
		monkeys["humn"] = strconv.Itoa(v)
		return solve(monkeys, s[0]) - solve(monkeys, s[2])
	})
	return part2
}

func solve(monkeys map[string]string, expr string) int {
	if v, err := strconv.Atoi(monkeys[expr]); err == nil {
		return v
	}

	s := strings.Fields(monkeys[expr])
	return map[string]func(int, int) int{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
	}[s[1]](solve(monkeys, s[0]), solve(monkeys, s[2]))
}
