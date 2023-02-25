package day23_test

import (
	"os"
	"testing"

	"day23"
)

func TestPartOneTwoExampleInput(t *testing.T) {
	t.Parallel()
	input, err := os.Open("./testdata/example.txt")
	if err != nil {
		t.Fatal(err)
	}

	wantPartOne, wantPartTwo := 110, 20
	gotPartOne, gotPartTwo := day23.PartOneTwo(input)

	if wantPartOne != gotPartOne {
		t.Errorf("wanted %d; got %d", wantPartOne, gotPartOne)
	}

	if wantPartTwo != gotPartTwo {
		t.Errorf("wanted %d; got %d", wantPartTwo, gotPartTwo)
	}
}

func TestPartOneTwoRealInput(t *testing.T) {
	t.Parallel()
	input, err := os.Open("./testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}

	wantPartOne, wantPartTwo := 3871, 925
	gotPartOne, gotPartTwo := day23.PartOneTwo(input)

	if wantPartOne != gotPartOne {
		t.Errorf("wanted %d; got %d", wantPartOne, gotPartOne)
	}

	if wantPartTwo != gotPartTwo {
		t.Errorf("wanted %d; got %d", wantPartTwo, gotPartTwo)
	}
}
