package day14

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

const separator = " -> "

type point struct {
	x, y int
}

func (p point) add(q point) point {
	return point{x: p.x + q.x, y: p.y + q.y}
}

func compare(i int) int {
	if i < 0 {
		return -1
	} else if i > 0 {
		return 1
	}
	return 0
}

func PartOneTwo(reader io.Reader) (int, int) {
	input, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	rock, maxy := map[point]struct{}{}, 0
	for _, s := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		s := strings.Split(s, separator)

		for i := 0; i < len(s)-1; i++ {
			var p, q point
			fmt.Sscanf(s[i], "%d,%d", &p.x, &p.y)
			fmt.Sscanf(s[i+1], "%d,%d", &q.x, &q.y)

			for d := (point{compare(q.x - p.x), compare(q.y - p.y)}); p != q.add(d); p = p.add(d) {
				rock[p] = struct{}{}
				if p.y > maxy {
					maxy = p.y
				}
			}
		}
	}

	d := []point{{0, 1}, {-1, 1}, {1, 1}}

	part1, part2 := (*int)(nil), 0
	for {
		p := point{500, 0}

		for i := 0; i < len(d); i++ {
			if _, ok := rock[p.add(d[i])]; !ok && p.add(d[i]).y < maxy+2 {
				p = p.add(d[i])
				if c := part2; part1 == nil && p.y >= maxy {
					part1 = &c
				}
				i = -1
			}
		}

		rock[p] = struct{}{}
		part2++
		if p.y == 0 {
			break
		}
	}
	return *part1, part2
}
