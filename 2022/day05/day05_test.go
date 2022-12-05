package day05_test

import (
	"os"
	"testing"

	"day05"
)

func TestPartOneExampleInput(t *testing.T) {
	t.Parallel()

	input, err := os.Open("./testdata/example.txt")
	if err != nil {
		t.Fatal(err)
	}

	stackOne := []string{"Z", "N"}
	stackTwo := []string{"M", "C", "D"}
	stackThree := []string{"P"}
	stacks := [][]string{stackOne, stackTwo, stackThree}
	want := "CMZ"
	got := day05.PartOne(stacks, input)

	if want != got {
		t.Errorf("want %s; got %s", want, got)
	}
}

func TestPartOneRealInput(t *testing.T) {
	t.Parallel()

	input, err := os.Open("./testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}

	stacks := [][]string{
		{"R", "G", "H", "Q", "S", "B", "T", "N"},
		{"H", "S", "F", "D", "P", "Z", "J"},
		{"Z", "H", "V"},
		{"M", "Z", "J", "F", "G", "H"},
		{"T", "Z", "C", "D", "L", "M", "S", "R"},
		{"M", "T", "W", "V", "H", "Z", "J"},
		{"T", "F", "P", "L", "Z"},
		{"Q", "V", "W", "S"},
		{"W", "H", "L", "M", "T", "D", "N", "C"},
	}
	want := "PTWLTDSJV"
	got := day05.PartOne(stacks, input)

	if want != got {
		t.Errorf("want %s; got %s", want, got)
	}
}

func TestPartTwoExampleInput(t *testing.T) {
	t.Parallel()

	input, err := os.Open("./testdata/example.txt")
	if err != nil {
		t.Fatal(err)
	}

	stackOne := []string{"Z", "N"}
	stackTwo := []string{"M", "C", "D"}
	stackThree := []string{"P"}
	stacks := [][]string{stackOne, stackTwo, stackThree}
	want := "MCD"
	got := day05.PartTwo(stacks, input)

	if want != got {
		t.Errorf("want %s; got %s", want, got)
	}
}

func TestPartTwoRealInput(t *testing.T) {
	t.Parallel()

	input, err := os.Open("./testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}

	stacks := [][]string{
		{"R", "G", "H", "Q", "S", "B", "T", "N"},
		{"H", "S", "F", "D", "P", "Z", "J"},
		{"Z", "H", "V"},
		{"M", "Z", "J", "F", "G", "H"},
		{"T", "Z", "C", "D", "L", "M", "S", "R"},
		{"M", "T", "W", "V", "H", "Z", "J"},
		{"T", "F", "P", "L", "Z"},
		{"Q", "V", "W", "S"},
		{"W", "H", "L", "M", "T", "D", "N", "C"},
	}
	want := "WZMFVGGZP"
	got := day05.PartTwo(stacks, input)

	if want != got {
		t.Errorf("want %s; got %s", want, got)
	}
}
