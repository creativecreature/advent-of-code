package day03

import (
	"bufio"
	"io"
	"io/ioutil"
	"strings"
	"unicode"
)

// Lowercase item types a through z have priorities 1 through 26.
// Uppercase item types A through Z have priorities 27 through 52.
func getRuneScore(r rune) int {
	if unicode.IsLower(r) {
		return int(r) - 96
	} else {
		return int(r) - 38
	}
}

func getRuneMap(s string) map[rune]bool {
	entries := make(map[rune]bool)
	for _, r := range s {
		entries[r] = true
	}
	return entries
}

func PartOne(input io.Reader) int {
	scanner := bufio.NewScanner(input)
	count := 0
	for scanner.Scan() {
		text := scanner.Text()
		left, right := text[:len(text)/2], text[len(text)/2:]
		entries := getRuneMap(left)
		for _, r := range right {
			if entries[r] {
				count += getRuneScore(r)
				break
			}
		}
	}

	return count
}

func PartTwo(input io.Reader) int {
	b, err := ioutil.ReadAll(input)
	if err != nil {
		panic(err)
	}

	inputSlice := strings.Split(string(b), "\n")
	count := 0
	for i := 0; i < len(inputSlice)-1; i += 3 {
		first, second, third := inputSlice[i], inputSlice[i+1], inputSlice[i+2]
		firstEntries, secondEntries := getRuneMap(first), getRuneMap(second)
		for _, r := range third {
			if firstEntries[r] && secondEntries[r] {
				count += getRuneScore(r)
				break
			}
		}
	}
	return count
}
