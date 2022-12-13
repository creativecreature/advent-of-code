package day04_test

import (
	"os"
	"testing"

	"day04"
)

func TestPartOneExampleInput(t *testing.T) {
	input, err := os.Open("./testdata/example.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 2
	got := day04.PartOne(input)
	if want != got {
		t.Errorf("wanted %d, got %d", want, got)
	}
}

func TestPartOneRealInput(t *testing.T) {
	input, err := os.Open("./testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 509
	got := day04.PartOne(input)
	if want != got {
		t.Errorf("wanted %d, got %d", want, got)
	}
}

func TestPartTwoExampleInput(t *testing.T) {
	input, err := os.Open("./testdata/example.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 4
	got := day04.PartTwo(input)
	if want != got {
		t.Errorf("wanted %d, got %d", want, got)
	}
}

func TestPartTwoRealInput(t *testing.T) {
	input, err := os.Open("./testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 870
	got := day04.PartTwo(input)
	if want != got {
		t.Errorf("wanted %d, got %d", want, got)
	}
}
