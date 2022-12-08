package day08_test

import (
	"os"
	"testing"

	"day08"
)

func TestPartOneExampleInput(t *testing.T) {
	file, err := os.Open("./testdata/example.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 21
	got := day08.PartOne(file)

	if want != got {
		t.Errorf("want %d; got %d", want, got)
	}
}

func TestPartOneRealInput(t *testing.T) {
	file, err := os.Open("./testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 1693
	got := day08.PartOne(file)

	if want != got {
		t.Errorf("want %d; got %d", want, got)
	}
}

func TestPartTwoExampleInput(t *testing.T) {
	file, err := os.Open("./testdata/example.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 8
	got := day08.PartTwo(file)

	if want != got {
		t.Errorf("want %d; got %d", want, got)
	}
}

func TestPartTwoRealInput(t *testing.T) {
	file, err := os.Open("./testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 422059
	got := day08.PartTwo(file)

	if want != got {
		t.Errorf("want %d; got %d", want, got)
	}
}
