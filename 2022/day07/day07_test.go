package day07_test

import (
	"os"
	"testing"

	"day07"
)

func TestPartOneExampleInput(t *testing.T) {
	input, err := os.Open("./testdata/example.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer input.Close()

	want := 95437
	got := day07.PartOne(input)

	if got != want {
		t.Errorf("want %d; got %d", want, got)
	}
}

func TestPartOneRealInput(t *testing.T) {
	input, err := os.Open("./testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer input.Close()

	want := 1723892
	got := day07.PartOne(input)

	if got != want {
		t.Errorf("want %d; got %d", want, got)
	}
}

func TestPartTwoExampleInput(t *testing.T) {
	input, err := os.Open("./testdata/example.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer input.Close()

	want := 24933642
	got := day07.PartTwo(input)

	if got != want {
		t.Errorf("want %d; got %d", want, got)
	}
}

func TestPartTwoRealInput(t *testing.T) {
	input, err := os.Open("./testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer input.Close()

	want := 8474158
	got := day07.PartTwo(input)

	if got != want {
		t.Errorf("want %d; got %d", want, got)
	}
}
