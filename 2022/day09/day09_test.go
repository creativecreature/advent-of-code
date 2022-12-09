package day09_test

import (
	"os"
	"testing"

	"day09"
)

func TestPartOneExampleInput(t *testing.T) {
	t.Parallel()

	input, err := os.Open("./testdata/example.txt")
	if err != nil {
		t.Fatal(err)
	}

	got := day09.PartOne(input)
	want := 13

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

	got := day09.PartOne(input)
	want := 5858

	if want != got {
		t.Errorf("want %d; got %d", want, got)
	}
}

func TestPartTwoExampleInput(t *testing.T) {
	t.Parallel()

	input, err := os.Open("./testdata/example2.txt")
	if err != nil {
		t.Fatal(err)
	}

	got := day09.PartTwo(input)
	want := 36

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

	got := day09.PartTwo(input)
	want := 2602

	if want != got {
		t.Errorf("want %d; got %d", want, got)
	}
}
