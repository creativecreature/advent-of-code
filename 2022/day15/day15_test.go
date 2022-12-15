package day15_test

import (
	"os"
	"testing"

	"day15"
)

func TestPartOneExampleData(t *testing.T) {
	input, err := os.Open("./testdata/example.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 26
	got := day15.PartOne(input)

	if got != want {
		t.Errorf("wanted %d; got %d", want, got)
	}
}

func TestPartOneRealData(t *testing.T) {
	input, err := os.Open("./testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 26
	got := day15.PartOne(input)

	if got != want {
		t.Errorf("wanted %d; got %d", want, got)
	}
}
