package day16_test

import (
	"os"
	"testing"

	"day16"
)

func TestPartOneExampleInput(t *testing.T) {
	t.Parallel()
	input, err := os.Open("./testdata/example.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 1651
	got := day16.PartOne(input)
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

	want := 1754
	got := day16.PartOne(input)
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

	want := 2474
	got := day16.PartTwo(input)
	if want != got {
		t.Errorf("wanted %d; got %d", want, got)
	}
}
