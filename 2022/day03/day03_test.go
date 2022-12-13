package day03_test

import (
	"os"
	"testing"

	"day03"
)

func TestPartOneExampleInput(t *testing.T) {
	input, err := os.Open("./testdata/example.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 157
	got := day03.PartOne(input)

	if want != got {
		t.Errorf("wanted %d; got %d", want, got)
	}
}

func TestPartOneRealInput(t *testing.T) {
	input, err := os.Open("./testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 8401
	got := day03.PartOne(input)

	if want != got {
		t.Errorf("wanted %d; got %d", want, got)
	}
}

func TestPartTwoExampleInput(t *testing.T) {
	input, err := os.Open("./testdata/example.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 70
	got := day03.PartTwo(input)

	if want != got {
		t.Errorf("wanted %d; got %d", want, got)
	}
}

func TestPartTwoRealInput(t *testing.T) {
	input, err := os.Open("./testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 2641
	got := day03.PartTwo(input)

	if want != got {
		t.Errorf("wanted %d; got %d", want, got)
	}
}
