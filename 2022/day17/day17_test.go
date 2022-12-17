package day17_test

import (
	"os"
	"testing"

	"day17"
)

func TestPartOneExampleInput(t *testing.T) {
	input, err := os.Open("./testdata/example.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 3068
	got := day17.PartOne(input)
	if want != got {
		t.Errorf("wanted %d; got %d", want, got)
	}
}
