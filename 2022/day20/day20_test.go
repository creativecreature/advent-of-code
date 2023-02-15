package day20_test

import (
	"os"
	"testing"

	"day20"
)

func TestPartOneExampleInput(t *testing.T) {
	t.Parallel()
	input, err := os.Open("./testdata/example.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 3
	got := day20.PartOne(input)
	if want != got {
		t.Errorf("wanted %d; got %d", want, got)
	}
}

func TestPartOneRealInput(t *testing.T) {
	t.Parallel()
	input, err := os.Open("./testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 6640
	got := day20.PartOne(input)
	if want != got {
		t.Errorf("wanted %d; got %d", want, got)
	}
}

func TestPartTwoExampleInput(t *testing.T) {
	t.Parallel()
	input, err := os.Open("./testdata/example.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 1623178306
	got := day20.PartTwo(input)
	if want != got {
		t.Errorf("wanted %d; got %d", want, got)
	}
}

func TestPartTwoRealInput(t *testing.T) {
	t.Parallel()
	input, err := os.Open("./testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 11893839037215
	got := day20.PartTwo(input)
	if want != got {
		t.Errorf("wanted %d; got %d", want, got)
	}
}
