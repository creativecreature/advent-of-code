package day14_test

import (
	"os"
	"testing"

	"day14"
)

func TestExampleInput(t *testing.T) {
	t.Parallel()
	input, err := os.Open("./testdata/example.txt")
	if err != nil {
		t.Fatal(err)
	}

	partOneWant, partTwoWant := 24, 93
	partOneGot, partTwoGot := day14.PartOneTwo(input)

	if partOneGot != partOneWant {
		t.Errorf("wanted %d; got %d", partOneWant, partOneGot)
	}

	if partTwoGot != partTwoWant {
		t.Errorf("wanted %d; got %d", partTwoWant, partTwoGot)
	}
}

func TestRealInput(t *testing.T) {
	t.Parallel()
	input, err := os.Open("./testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}

	partOneWant, partTwoWant := 618, 26358
	partOneGot, partTwoGot := day14.PartOneTwo(input)

	if partOneGot != partOneWant {
		t.Errorf("wanted %d; got %d", partOneWant, partOneGot)
	}

	if partTwoGot != partTwoWant {
		t.Errorf("wanted %d; got %d", partTwoWant, partTwoGot)
	}
}
