package day02_test

import (
	"os"
	"testing"

	"day02"
)

func TestPartTwoExampleInput(t *testing.T) {
	t.Parallel()

	input, err := os.Open("./testdata/example.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 12
	got := day02.PartTwo(input)

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

	want := 15457
	got := day02.PartTwo(input)

	if want != got {
		t.Errorf("want %d; got %d", want, got)
	}
}
