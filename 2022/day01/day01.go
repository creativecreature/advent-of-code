package day01

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"sort"
	"strconv"
)

func getCaloriesDescending(input io.Reader) []int {
	calories := []int{}
	currentCount := 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		if len(bytes.TrimSpace(scanner.Bytes())) < 1 {
			calories = append(calories, currentCount)
			currentCount = 0
			continue
		}

		calories, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		currentCount += calories
	}

	// The input does not end with an empty line so we need to append the last one
	calories = append(calories, currentCount)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Sort the calories in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(calories)))
	return calories
}

func PartOne(input io.Reader) int {
	return getCaloriesDescending(input)[0]
}

func PartTwo(input io.Reader) int {
	sortedSliceOfCalories := getCaloriesDescending(input)
	return sortedSliceOfCalories[0] + sortedSliceOfCalories[1] + sortedSliceOfCalories[2]
}
