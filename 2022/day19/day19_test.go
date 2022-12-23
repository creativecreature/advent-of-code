package day19_test

import (
	"os"
	"testing"

	"day19"
)

func TestPartOneExampleInput(t *testing.T) {
	input, err := os.Open("./testdata/example.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 33
	got := day19.PartOne(input)
	if want != got {
		t.Errorf("wanted %d; got %d", want, got)
	}
}

func TestPartOneRealInput(t *testing.T) {
	input, err := os.Open("./testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 1962
	got := day19.PartOne(input)
	if want != got {
		t.Errorf("wanted %d; got %d", want, got)
	}
}

func TestPartTwoRealInput(t *testing.T) {
	input, err := os.Open("./testdata/input2.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 88160
	got := day19.PartTwo(input)
	if want != got {
		t.Errorf("wanted %d; got %d", want, got)
	}
}
