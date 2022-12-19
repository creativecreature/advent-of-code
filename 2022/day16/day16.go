package day16

import (
	"bufio"
	"fmt"
	"io"
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

func generateNodeMaps(nodes map[string]node) map[string]map[string]int {
	nodemap := map[string]map[string]int{}
	for _, n := range nodes {
		nodeMap := make(map[string]int)
		queue := []workitem{{n, 0}}

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

func calculatePaths(current, path string, seen map[string]bool, score, stepsLeft int, scores map[string]int, nodemaps map[string]map[string]int, nodes map[string]node) {
	seen[current] = true
	score = score + (nodes[current].rate * stepsLeft)
	scores[path+current] = score

	currentNodemap := nodemaps[current]
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
		calculatePaths(key, path+current, newSeen, score, stepsLeft-(dist+1), scores, nodemaps, nodes)
	}
}

func PartOne(input io.Reader) int {
	nodes := parseInput(input)
	nodemaps := generateNodeMaps(nodes)
	scores := make(map[string]int)
	visited := make(map[string]bool)
	calculatePaths("AA", "", visited, 0, 30, scores, nodemaps, nodes)

	bestScore := 0
	for _, score := range scores {
		if score > bestScore {
			bestScore = score
		}
	}

	return bestScore
}

func PartTwo(input io.Reader) int {
	nodes := parseInput(input)
	nodemaps := generateNodeMaps(nodes)
	scores := make(map[string]int)
	visited := make(map[string]bool)
	calculatePaths("AA", "", visited, 0, 26, scores, nodemaps, nodes)

	// This is a kinda hacky way to reduce the number of paths
	for s := range scores {
		if len(s) < 8 || scores[s] < 700 {
			delete(scores, s)
		}
	}

	// These are the two paths we want for the example input:
	// "AADDHHEE" "AAJJBBCC"
	bestPathOne, bestPathTwo := "", ""
	bestScore := 0
	for pathOne, scoreOne := range scores {
		for pathTwo, scoreTwo := range scores {
			if !strings.ContainsAny(pathOne, pathTwo[2:]) && !strings.ContainsAny(pathTwo, pathOne[2:]) {
				score := scoreOne + scoreTwo
				if score > bestScore {
					bestPathOne = pathOne
					bestPathTwo = pathTwo
					bestScore = score
				}
			}
		}
	}

	fmt.Println(bestPathOne)
	fmt.Println(bestPathTwo)
	fmt.Println("COMBINED")

	return bestScore
}
