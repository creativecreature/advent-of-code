package day11_test

import (
	"os"
	"testing"

	"day11"
)

func TestPartOneExampleInput(t *testing.T) {
	input, err := os.Open("./testdata/example.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 10605
	got := day11.PartOne(input)

	if want != got {
		t.Errorf("want %d; got %d", want, got)
	}
}

func TestPartOneRealInput(t *testing.T) {
	input, err := os.Open("./testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 99852
	got := day11.PartOne(input)

	if want != got {
		t.Errorf("want %d; got %d", want, got)
	}
}

func TestPartTwoExampleInput(t *testing.T) {
	input, err := os.Open("./testdata/example.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 2713310158
	got := day11.PartTwo(input)

	if want != got {
		t.Errorf("want %d; got %d", want, got)
	}
}

func TestPartTwoRealInput(t *testing.T) {
	input, err := os.Open("./testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 25935263541
	got := day11.PartTwo(input)

	if want != got {
		t.Errorf("want %d; got %d", want, got)
	}
}
