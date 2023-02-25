package day24

import (
	"image"
	"io"
	"io/ioutil"
	"strings"
)

func getValleys(i io.Reader) map[image.Point]rune {
	input, _ := ioutil.ReadAll(i)
	valleys := map[image.Point]rune{}
	for y, s := range strings.Fields(string(input)) {
		for x, r := range s {
			valleys[image.Point{x, y}] = r
		}
	}
	return valleys
}

func getBlizzard(valleys map[image.Point]rune) image.Rectangle {
	var blizzard image.Rectangle
	for p := range valleys {
		blizzard = blizzard.Union(image.Rectangle{p, p.Add(image.Point{1, 1})})
	}
	blizzard.Min, blizzard.Max = blizzard.Min.Add(image.Point{1, 1}), blizzard.Max.Sub(image.Point{1, 1})
	return blizzard
}

type State struct {
	P image.Point
	T int
}

func bfs(valleys map[image.Point]rune, blizzard image.Rectangle, start image.Point, end image.Point, time int) int {
	delta := map[rune]image.Point{
		'#': {0, 0}, '^': {0, -1}, '>': {1, 0}, 'v': {0, 1}, '<': {-1, 0},
	}

	queue := []State{{start, time}}
	seen := map[State]struct{}{queue[0]: {}}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

	loop:
		for _, d := range delta {
			next := State{cur.P.Add(d), cur.T + 1}
			if next.P == end {
				return next.T
			}

			if _, ok := seen[next]; ok {
				continue
			}
			if r, ok := valleys[next.P]; !ok || r == '#' {
				continue
			}

			if next.P.In(blizzard) {
				for r, d := range delta {
					if valleys[next.P.Sub(d.Mul(next.T)).Mod(blizzard)] == r {
						continue loop
					}
				}
			}

			seen[next] = struct{}{}
			queue = append(queue, next)
		}
	}
	return -1
}

func PartOne(i io.Reader) int {
	valleys := getValleys(i)
	blizzard := getBlizzard(valleys)
	start, end := blizzard.Min.Sub(image.Point{0, 1}), blizzard.Max.Sub(image.Point{1, 0})
	return bfs(valleys, blizzard, start, end, 0)
}

func PartTwo(i io.Reader) int {
	valleys := getValleys(i)
	blizzard := getBlizzard(valleys)
	start, end := blizzard.Min.Sub(image.Point{0, 1}), blizzard.Max.Sub(image.Point{1, 0})
	return bfs(valleys, blizzard, start, end, bfs(valleys, blizzard, end, start, bfs(valleys, blizzard, start, end, 0)))
}
