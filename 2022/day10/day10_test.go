package day10_test

import (
	"os"
	"testing"

	"day10"
)

func TestPartOneExampleInput(t *testing.T) {
	input, err := os.Open("./testdata/example.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 13140
	got := day10.PartOneTwo(input)

	if want != got {
		t.Errorf("want %d; got %d", want, got)
	}
}

func TestPartOneRealInput(t *testing.T) {
	input, err := os.Open("./testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 14240
	got := day10.PartOneTwo(input)

	if want != got {
		t.Errorf("want %d; got %d", want, got)
	}
}
