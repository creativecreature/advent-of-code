package day16

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type valve struct {
	name  string
	rate  int
	nodes []string
}
type volcano struct {
	paths     map[string]int
	valves    map[string]valve
	valveMaps map[string]map[string]int
}

type queueItem struct {
	valve valve
	count int
}

func parseValves(input io.Reader) map[string]valve {
	pattern := regexp.MustCompile("Valve (?P<name>[A-Z].) has flow rate=(?P<rate>[0-9]+); tunnels? leads? to valves? (?P<nodes>.*)")
	scanner := bufio.NewScanner(input)
	valves := make(map[string]valve)
	for scanner.Scan() {
		match := pattern.FindStringSubmatch(scanner.Text())
		name := match[pattern.SubexpIndex("name")]
		rate, _ := strconv.Atoi(match[pattern.SubexpIndex("rate")])
		vs := strings.Split(match[pattern.SubexpIndex("nodes")], ", ")
		v := valve{name, rate, vs}
		valves[name] = v
	}
	return valves
}

func createValveMaps(nodes map[string]valve) map[string]map[string]int {
	nodemap := map[string]map[string]int{}
	for _, n := range nodes {
		nodeMap := make(map[string]int)
		queue := []queueItem{{n, 0}}

		for len(queue) > 0 {
			currentItem := queue[0]
			queue = queue[1:]

			if _, ok := nodeMap[currentItem.valve.name]; ok {
				continue
			}

			nodeMap[currentItem.valve.name] = currentItem.count

			for _, c := range currentItem.valve.nodes {
				child := nodes[c]
				queue = append(queue, queueItem{valve: child, count: currentItem.count + 1})
			}
		}

		// Cleanup any nodes that gives us 0
		for n := range nodeMap {
			if nodes[n].rate == 0 && nodes[n].name != "AA" {
				delete(nodeMap, n)
			}
		}

		nodemap[n.name] = nodeMap
	}
	return nodemap
}

func newVolcano(input io.Reader) *volcano {
	valves := parseValves(input)
	valvesMaps := createValveMaps(valves)
	return &volcano{
		paths:     make(map[string]int),
		valves:    valves,
		valveMaps: valvesMaps,
	}
}

func (v *volcano) calculatePaths(current, path string, seen map[string]bool, score, stepsLeft int) {
	seen[current] = true
	score = score + (v.valves[current].rate * stepsLeft)
	v.paths[path+current] = score

	currentNodemap := v.valveMaps[current]
	for key, dist := range currentNodemap {
		if seen[key] {
			continue
		}
		if dist+1 > stepsLeft {
			continue
		}
		newSeen := make(map[string]bool)
		for k, v := range seen {
			newSeen[k] = v
		}
		v.calculatePaths(key, path+current, newSeen, score, stepsLeft-(dist+1))
	}
}

func overlaps(s1, s2 string) bool {
	keys := make(map[string]bool)
	for i := 2; i < len(s1)-1; i++ {
		key := fmt.Sprintf("%s%s", string(s1[i]), string(s1[i+1]))
		if keys[key] {
			return true
		}
		keys[key] = true
	}
	for i := 2; i < len(s2)-1; i++ {
		key := fmt.Sprintf("%s%s", string(s2[i]), string(s2[i+1]))
		if keys[key] {
			return true
		}
		keys[key] = true
	}
	return false
}

func (v *volcano) getBestPath() int {
	bestScore := 0
	for _, score := range v.paths {
		if score > bestScore {
			bestScore = score
		}
	}
	return bestScore
}

func (v *volcano) getBestUniquePaths() int {
	// This is a kinda hacky way to reduce the number of paths
	for s := range v.paths {
		if len(s) < 14 || v.paths[s] < 1000 {
			delete(v.paths, s)
		}
	}

	bestScore := 0
	for pathOne, scoreOne := range v.paths {
		for pathTwo, scoreTwo := range v.paths {
			if !overlaps(pathOne, pathTwo) {
				score := scoreOne + scoreTwo
				if score > bestScore {
					bestScore = score
				}
			}
		}
	}

	return bestScore
}

func PartOne(input io.Reader) int {
	v := newVolcano(input)
	visited := make(map[string]bool)
	v.calculatePaths("AA", "", visited, 0, 30)
	return v.getBestPath()
}

func PartTwo(input io.Reader) int {
	v := newVolcano(input)
	visited := make(map[string]bool)
	v.calculatePaths("AA", "", visited, 0, 26)
	return v.getBestUniquePaths()
}
