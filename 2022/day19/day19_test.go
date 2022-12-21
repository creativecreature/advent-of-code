package day19_test

import (
	"os"
	"testing"

	"day19"
)

func TestPartOneExampleInput(t *testing.T) {
	input, err := os.Open("./testdata/example.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := 33
	got := day19.PartOne(input)
	if want != got {
		t.Errorf("wanted %d; got %d", want, got)
	}
}
