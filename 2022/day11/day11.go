package day11

import (
	"io"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	items        []int
	operation    func(num int) int
	operationNum int
	divideBy     int
	nextIfTrue   int
	nextIfFalse  int
	score        int
}

var (
	itemPattern       = regexp.MustCompile("Starting items: (.*)$")
	operationPattern  = regexp.MustCompile("new = old ([*|+]) ([0-9]+|old)")
	divideByPattern   = regexp.MustCompile("divisible by ([0-9]+)")
	nextMonkeyPattern = regexp.MustCompile("throw to monkey ([0-9]+)")
)

func parseItems(itemString string) []int {
	match := itemPattern.FindSubmatch([]byte(itemString))
	rawItems := strings.Split(string(match[1]), ",")
	items := []int{}
	for _, i := range rawItems {
		item, err := strconv.Atoi(strings.Trim(i, " "))
		if err != nil {
			panic(err)
		}
		items = append(items, item)
	}
	return items
}

func parseOperation(rawOperation string) func(int) int {
	match := operationPattern.FindSubmatch([]byte(rawOperation))
	return func(num int) int {
		var val int
		if string(match[2]) == "old" {
			val = num
		} else {
			v, err := strconv.Atoi(string(match[2]))
			if err != nil {
				panic(err)
			}
			val = v
		}

		if string(match[1]) == "*" {
			return num * val
		}
		return num + val
	}
}

func extractIntFromMatch(str string, r *regexp.Regexp) int {
	match := r.FindSubmatch([]byte(str))
	i, err := strconv.Atoi(string(match[1]))
	if err != nil {
		panic(err)
	}
	return i
}

func parseInput(input io.Reader) []*monkey {
	bytes, err := ioutil.ReadAll(input)
	if err != nil {
		panic(err)
	}
	monkeys := []*monkey{}
	inputSlice := strings.Split(string(bytes), "\n")
	for i := 1; i < len(inputSlice); i += 7 {
		items := parseItems(inputSlice[i])
		operation := parseOperation(inputSlice[i+1])
		divideBy := extractIntFromMatch(inputSlice[i+2], divideByPattern)
		nextIfTrue := extractIntFromMatch(inputSlice[i+3], nextMonkeyPattern)
		nextIfFalse := extractIntFromMatch(inputSlice[i+4], nextMonkeyPattern)
		monkeys = append(monkeys, &monkey{
			items:       items,
			operation:   operation,
			divideBy:    divideBy,
			nextIfTrue:  nextIfTrue,
			nextIfFalse: nextIfFalse,
			score:       0,
		})
	}
	return monkeys
}

func getScore(monkeys []*monkey, iterations int, reduceWorry func(int) int) int {
	for i := 0; i < iterations; i++ {
		for _, m := range monkeys {
			for len(m.items) != 0 {
				m.score++
				o := reduceWorry(m.operation(m.items[0]))
				m.items = m.items[1:]
				if o%m.divideBy == 0 {
					monkeys[m.nextIfTrue].items = append(monkeys[m.nextIfTrue].items, o)
				} else {
					monkeys[m.nextIfFalse].items = append(monkeys[m.nextIfFalse].items, o)
				}
			}
		}
	}
	scores := []int{}
	for _, m := range monkeys {
		scores = append(scores, m.score)
	}
	sort.Ints(scores)
	return scores[len(scores)-1] * scores[len(scores)-2]
}

func PartOne(input io.Reader) int {
	monkeys := parseInput(input)
	reduceWorry := func(num int) int {
		return num / 3
	}
	return getScore(monkeys, 20, reduceWorry)
}

func PartTwo(input io.Reader) int {
	monkeys := parseInput(input)
	prime := 1
	for _, m := range monkeys {
		prime = prime * m.divideBy
	}
	reduceWorry := func(num int) int {
		return num % prime
	}
	return getScore(monkeys, 10000, reduceWorry)
}
