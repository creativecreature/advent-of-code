package day13

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strings"
)

func cmp(a, b any) int {
	as, aok := a.([]any)
	bs, bok := b.([]any)

	switch {
	case !aok && !bok:
		return int(a.(float64) - b.(float64))
	case !aok:
		as = []any{a}
	case !bok:
		bs = []any{b}
	}

	for i := 0; i < len(as) && i < len(bs); i++ {
		if c := cmp(as[i], bs[i]); c != 0 {
			return c
		}
	}
	return len(as) - len(bs)
}

func parseInput(input io.Reader) ([]any, int) {
	rawInput, _ := ioutil.ReadAll(input)
	packages, total := []any{}, 0
	for i, s := range strings.Split(strings.TrimSpace(string(rawInput)), "\n\n") {
		s := strings.Split(s, "\n")
		var a, b any
		json.Unmarshal([]byte(s[0]), &a)
		json.Unmarshal([]byte(s[1]), &b)
		packages = append(packages, a, b)
		if cmp(a, b) <= 0 {
			total += i + 1
		}

	}
	return packages, total
}

func PartOne(input io.Reader) int {
	_, total := parseInput(input)
	return total
}

func PartTwo(input io.Reader) int {
	packages, _ := parseInput(input)
	packages = append(packages, []any{[]any{2.}}, []any{[]any{6.}})
	sort.Slice(packages, func(i, j int) bool { return cmp(packages[i], packages[j]) < 0 })
	total := 1
	for i, p := range packages {
		if fmt.Sprint(p) == "[[2]]" || fmt.Sprint(p) == "[[6]]" {
			total *= i + 1
		}
	}
	return total
}
