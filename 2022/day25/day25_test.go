package day25_test

import (
	"os"
	"testing"

	"day25"
)

func TestPartOneExampleInput(t *testing.T) {
	t.Parallel()
	input, err := os.Open("./testdata/example.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := "2=-1=0"
	got := day25.PartOne(input)
	if want != got {
		t.Errorf("wanted %s; got %s", want, got)
	}
}

func TestPartOneRealInput(t *testing.T) {
	t.Parallel()
	input, err := os.Open("./testdata/input.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := "2-20=01--0=0=0=2-120"
	got := day25.PartOne(input)
	if want != got {
		t.Errorf("wanted %s; got %s", want, got)
	}
}
