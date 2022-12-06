package day06_test

import (
	"os"
	"testing"

	"day06"
)

func TestPartOneExampleInputOne(t *testing.T) {
	t.Parallel()

	input, err := os.Open("./testdata/example1.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 5
	got := day06.PartOne(input)

	if want != got {
		t.Errorf("want %d; got %d", want, got)
	}
}

func TestPartOneExampleInputTwo(t *testing.T) {
	t.Parallel()

	input, err := os.Open("./testdata/example2.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 6
	got := day06.PartOne(input)

	if want != got {
		t.Errorf("want %d; got %d", want, got)
	}
}

func TestPartOneExampleInputThree(t *testing.T) {
	t.Parallel()

	input, err := os.Open("./testdata/example3.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 10
	got := day06.PartOne(input)

	if want != got {
		t.Errorf("want %d; got %d", want, got)
	}
}

func TestPartOneExampleInputFour(t *testing.T) {
	t.Parallel()

	input, err := os.Open("./testdata/example4.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 11
	got := day06.PartOne(input)

	if want != got {
		t.Errorf("want %d; got %d", want, got)
	}
}

func TestPartOneRealInput(t *testing.T) {
	t.Parallel()

	input, err := os.Open("./testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 1912
	got := day06.PartOne(input)

	if want != got {
		t.Errorf("want %d; got %d", want, got)
	}
}

func TestPartTwoExampleInputTwo(t *testing.T) {
	t.Parallel()

	input, err := os.Open("./testdata/example2.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 23
	got := day06.PartTwo(input)

	if want != got {
		t.Errorf("want %d; got %d", want, got)
	}
}

func TestPartTwoExampleInputThree(t *testing.T) {
	t.Parallel()

	input, err := os.Open("./testdata/example3.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 29
	got := day06.PartTwo(input)

	if want != got {
		t.Errorf("want %d; got %d", want, got)
	}
}

func TestPartTwoExampleInputFour(t *testing.T) {
	t.Parallel()

	input, err := os.Open("./testdata/example4.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 26
	got := day06.PartTwo(input)

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

	want := 2122
	got := day06.PartTwo(input)

	if want != got {
		t.Errorf("want %d; got %d", want, got)
	}
}
