package day04

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func unsafeParse(str string) int {
	num, _ := strconv.Atoi(str)
	return num
}

func intersects(leftStart, leftEnd, rightStart, rightEnd int) bool {
	return (leftStart <= rightStart && rightEnd <= leftEnd) ||
		(rightStart <= leftStart && leftEnd <= rightEnd)
}

func overlaps(leftStart, leftEnd, rightStart, rightEnd int) bool {
	return (leftStart <= rightStart && (leftEnd) >= rightStart) ||
		(rightStart <= leftStart && (rightEnd) >= leftStart)
}

type compareFunc func(leftStart, leftEnd, rightStart, rightEnd int) bool

func compareSchedules(input io.Reader, compare compareFunc) int {
	scanner := bufio.NewScanner(input)
	count := 0
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), ",")
		left, right := strings.Split(text[0], "-"), strings.Split(text[1], "-")
		leftStart, leftEnd := unsafeParse(left[0]), unsafeParse(left[1])
		rightStart, rightEnd := unsafeParse(right[0]), unsafeParse(right[1])
		if compare(leftStart, leftEnd, rightStart, rightEnd) {
			count++
		}
	}
	return count
}

func PartOne(input io.Reader) int {
	return compareSchedules(input, intersects)
}

func PartTwo(input io.Reader) int {
	return compareSchedules(input, overlaps)
}
