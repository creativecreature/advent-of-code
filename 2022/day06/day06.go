package day06

import (
	"io"
	"io/ioutil"
	"log"
)

func findUniqueRunePattern(input io.Reader, numberOfUniqueCharacters int) int {
	bytes, err := ioutil.ReadAll(input)
	if err != nil {
		log.Fatal(err)
	}

	charsSeen := []rune{}
	for i, char := range string(bytes) {
		charsSeen = append(charsSeen, char)
		if len(charsSeen) == numberOfUniqueCharacters {
			charMap := make(map[rune]bool)
			for _, r := range charsSeen {
				if charMap[r] {
					break
				}
				charMap[r] = true
			}

			if len(charMap) == numberOfUniqueCharacters {
				return i + 1
			}
			charsSeen = charsSeen[1:]
		}
	}
	return -1
}

func PartOne(input io.Reader) int {
	return findUniqueRunePattern(input, 4)
}

func PartTwo(input io.Reader) int {
	return findUniqueRunePattern(input, 14)
}
