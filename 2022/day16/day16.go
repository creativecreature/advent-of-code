package day16

import (
	"bufio"
	"io"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type node struct {
	name  string
	rate  int
	nodes []string
}

type workitem struct {
	node  node
	count int
}

var pattern = regexp.MustCompile("Valve (?P<name>[A-Z].) has flow rate=(?P<rate>[0-9]+); tunnels? leads? to valves? (?P<nodes>.*)")

func parseInput(input io.Reader) map[string]node {
	scanner := bufio.NewScanner(input)
	nodeMap := make(map[string]node)
	for scanner.Scan() {
		match := pattern.FindStringSubmatch(scanner.Text())
		name := match[pattern.SubexpIndex("name")]
		rate, _ := strconv.Atoi(match[pattern.SubexpIndex("rate")])
		nodes := strings.Split(match[pattern.SubexpIndex("nodes")], ", ")
		node := node{name, rate, nodes}
		nodeMap[name] = node
	}
	return nodeMap
}

func generateHeatmap(nodes map[string]node) map[string]map[string]int {
	heatmap := map[string]map[string]int{}
	for _, n := range nodes {
		nodeMap := make(map[string]int)
		queue := []workitem{{n, 0}}

		if n.rate == 0 && n.name != "AA" {
			continue
		}

		for len(queue) > 0 {
			currentItem := queue[0]
			queue = queue[1:]

			if _, ok := nodeMap[currentItem.node.name]; ok {
				continue
			}

			nodeMap[currentItem.node.name] = currentItem.count

			for _, c := range currentItem.node.nodes {
				child := nodes[c]
				queue = append(queue, workitem{node: child, count: currentItem.count + 1})
			}
		}

		heatmap[n.name] = nodeMap
	}
	return heatmap
}

func rec(currentNode string, nodes map[string]node, heatmap map[string]map[string]int, visited map[string]bool, steps int) int {
	if visited[currentNode] {
		return 0
	}

	if steps <= 1 {
		return 0
	}

	visited[currentNode] = true
	m := heatmap[currentNode]
	bestChild := 0
	for n := range m {
		i := m[n]
		newMap := make(map[string]bool)
		for k, v := range visited {
			newMap[k] = v
		}
		childCount := float64(rec(n, nodes, heatmap, newMap, steps-(i+1)))
		bestChild = int(math.Max(float64(bestChild), childCount))
	}

	return bestChild + nodes[currentNode].rate*steps
}

func recTwo(currentNode string, nodes map[string]node, heatmap map[string]map[string]int, visited map[string]bool, steps int) int {
	if visited[currentNode] {
		return 0
	}

	if steps <= 1 {
		return 0
	}

	visited[currentNode] = true
	m := heatmap[currentNode]
	bestChild := 0
	for n := range m {
		i := m[n]
		newMap := make(map[string]bool)
		for k, v := range visited {
			newMap[k] = v
		}
		heatmapScore := i + 1
		childCount := float64(rec(n, nodes, heatmap, newMap, steps-(heatmapScore)))
		bestChild = int(math.Max(float64(bestChild), childCount))
	}

	return bestChild + nodes[currentNode].rate*steps
}

func PartOne(input io.Reader) int {
	nodes := parseInput(input)
	heatmap := generateHeatmap(nodes)
	scores := make(map[string]int)
	visited := make(map[string]bool)
	calculatePaths("AA", "", visited, 0, 30, scores, heatmap, nodes)

	bestScore := 0
	for _, score := range scores {
		if score > bestScore {
			bestScore = score
		}
	}

	return bestScore
}

type path struct {
	currentNode string
	path        string
	score       int
}

func calculatePaths(current, parent string, seen map[string]bool, score, steps int, scores map[string]int, heatmap map[string]map[string]int, nodes map[string]node) {
	score = score + (nodes[current].rate * steps)
	seen[current] = true

	currentHeatmap := heatmap[current]
	for key := range currentHeatmap {
		if seen[key] || (currentHeatmap[key])+1 > steps {
			continue
		}
		dist := currentHeatmap[key]
		newSeen := make(map[string]bool)
		for k, v := range seen {
			newSeen[k] = v
		}

		calculatePaths(key, parent+current, newSeen, score, steps-(dist+1), scores, heatmap, nodes)
	}

	scores[parent] = score
}

func PartTwo(input io.Reader) int {
	nodes := parseInput(input)
	heatmap := generateHeatmap(nodes)

	scores := make(map[string]int)
	visited := make(map[string]bool)
	calculatePaths("AA", "", visited, 0, 26, scores, heatmap, nodes)

	bestScore := 0
	for pathOne, scoreOne := range scores {
		for pathTwo, scoreTwo := range scores {
			if pathOne == pathTwo {
				continue
			}
			if len(pathTwo) < 2 || len(pathOne) < 2 {
				continue
			}

			if !strings.ContainsAny(pathOne, pathTwo[2:]) && !strings.ContainsAny(pathTwo, pathOne[2:]) {
				score := scoreOne + scoreTwo
				if score > bestScore {
					bestScore = score
				}
			}
		}
	}

	return bestScore
}
