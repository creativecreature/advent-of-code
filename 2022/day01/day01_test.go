package day01_test

import (
	"os"
	"testing"

	"day01"
)

func TestPartOneExampleInput(t *testing.T) {
	t.Parallel()

	input, err := os.Open("./testdata/example.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 24000
	got := day01.PartOne(input)

	if want != got {
		t.Errorf("want %d; got %d", want, got)
	}
}

func TestPartOneRealInput(t *testing.T) {
	t.Parallel()

	input, err := os.Open("./testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 69206
	got := day01.PartOne(input)

	if want != got {
		t.Errorf("want %d; got %d", want, got)
	}
}

func TestPartTwoExampleInput(t *testing.T) {
	t.Parallel()

	input, err := os.Open("./testdata/example.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 45000
	got := day01.PartTwo(input)

	if want != got {
		t.Errorf("want %d; got %d", want, got)
	}
}

func TestPartTwoRealInput(t *testing.T) {
	t.Parallel()

	input, err := os.Open("./testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 197400
	got := day01.PartTwo(input)

	if want != got {
		t.Errorf("want %d; got %d", want, got)
	}
}
