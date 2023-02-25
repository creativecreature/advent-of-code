package day15_test

import (
	"os"
	"testing"

	"day15"
)

func TestPartOneTwoRealData(t *testing.T) {
	input, err := os.Open("./testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}

	wantPartOne, wantPartTwo := 4724228, 13622251246513
	gotPartOne, gotPartTwo := day15.PartOneTwo(input)

	if gotPartOne != wantPartOne {
		t.Errorf("wanted %d; got %d", wantPartOne, gotPartOne)
	}

	if gotPartTwo != wantPartTwo {
		t.Errorf("wanted %d; got %d", wantPartTwo, gotPartTwo)
	}
}
