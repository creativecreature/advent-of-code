package day21_test

import (
	"os"
	"testing"

	"day21"
)

func TestPartOneExampleInput(t *testing.T) {
	t.Parallel()
	input, err := os.Open("./testdata/example.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 152
	got := day21.PartOne(input)
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

	want := 51928383302238
	got := day21.PartOne(input)
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

	want := 301
	got := day21.PartTwo(input)
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

	want := 3305669217840
	got := day21.PartTwo(input)
	if want != got {
		t.Errorf("wanted %d; got %d", want, got)
	}
}
