package day17_test

import (
	"os"
	"testing"

	"day17"
)

func TestPartOneExampleInput(t *testing.T) {
	input, err := os.Open("./testdata/example.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 3068
	got := day17.PartOne(input)
	if want != got {
		t.Errorf("wanted %d; got %d", want, got)
	}
}

func TestPartOneRealInput(t *testing.T) {
	input, err := os.Open("./testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 3193
	got := day17.PartOne(input)
	if want != got {
		t.Errorf("wanted %d; got %d", want, got)
	}
}

func TestPartTwoExampleInput(t *testing.T) {
	input, err := os.Open("./testdata/example.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 1514285714288
	got := day17.PartTwo(input)
	if want != got {
		t.Errorf("wanted %d; got %d", want, got)
	}
}

func TestPartTwoRealInput(t *testing.T) {
	input, err := os.Open("./testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 1577650429835
	got := day17.PartTwo(input)
	if want != got {
		t.Errorf("wanted %d; got %d", want, got)
	}
}
