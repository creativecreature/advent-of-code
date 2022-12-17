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
	node  *node
	count int
}

var pattern = regexp.MustCompile("Valve (?P<name>[A-Z].) has flow rate=(?P<rate>[0-9]+); tunnels? leads? to valves? (?P<nodes>.*)")

func parseInput(input io.Reader) map[string]*node {
	scanner := bufio.NewScanner(input)
	nodeMap := make(map[string]*node)
	for scanner.Scan() {
		match := pattern.FindStringSubmatch(scanner.Text())
		name := match[pattern.SubexpIndex("name")]
		rate, _ := strconv.Atoi(match[pattern.SubexpIndex("rate")])
		nodes := strings.Split(match[pattern.SubexpIndex("nodes")], ", ")
		node := &node{name, rate, nodes}
		nodeMap[name] = node
	}
	return nodeMap
}

func generateHeatmap(nodes map[string]*node) map[string]map[string]int {
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

func rec(currentNode string, nodes map[string]*node, heatmap map[string]map[string]int, visited map[string]bool, steps int) int {
	if _, ok := visited[currentNode]; ok {
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
		heatmapScore := i + 2
		childCount := float64(rec(n, nodes, heatmap, newMap, steps-(heatmapScore-1)))
		bestChild = int(math.Max(float64(bestChild), childCount))
	}

	return bestChild + nodes[currentNode].rate*steps
}

func recTwo(nodeOne, nodeTwo string, nodes map[string]*node, heatmap map[string]map[string]int, visited map[string]bool, stepsOne, stepsTwo int) int {
	if _, ok := visited[nodeOne]; ok {
		return 0
	}
	if _, ok := visited[nodeTwo]; ok {
		return 0
	}

	if stepsOne <= 2 && stepsTwo <= 2 {
		return 0
	}

	visited[nodeOne] = true
	visited[nodeTwo] = true

	bestChild := 0
	nodeOneHeatMap := heatmap[nodeOne]
	nodeTwoHeatMap := heatmap[nodeTwo]
	for n := range nodeOneHeatMap {
		if nodeOneHeatMap[n] > stepsOne {
			continue
		}
		for m := range nodeTwoHeatMap {
			if nodeTwoHeatMap[m] > stepsTwo {
				continue
			}
			if m == n {
				continue
			}

			nodeOneHeatmapScore := nodeOneHeatMap[n]
			heatmapScoreNodeOne := nodeOneHeatmapScore + 2
			nodeTwoHeatmapScore := nodeTwoHeatMap[m]
			heatmapScoreNodeTwo := nodeTwoHeatmapScore + 2

			newMap := make(map[string]bool)
			for k, v := range visited {
				newMap[k] = v
			}

			potentialBestChild := float64(recTwo(n, m, nodes, heatmap, newMap, stepsOne-(heatmapScoreNodeOne-1), stepsTwo-(heatmapScoreNodeTwo-1)))
			bestChild = int(math.Max(float64(bestChild), potentialBestChild))
		}
	}

	total := 0
	resultOne := (nodes[nodeOne].rate) * stepsOne
	resultTwo := (nodes[nodeTwo].rate) * stepsTwo

	if stepsOne > 1 {
		total = total + resultOne
	}

	if stepsTwo > 1 {
		total = total + resultTwo
	}
	return bestChild + total
}

func PartOne(input io.Reader) int {
	nodes := parseInput(input)
	heatmap := generateHeatmap(nodes)

	wuu := rec("AA", nodes, heatmap, map[string]bool{}, 30)

	// fmt.Println(order)
	return wuu
}

func PartTwo(input io.Reader) int {
	nodes := parseInput(input)
	heatmap := generateHeatmap(nodes)
	wuu := recTwo("AA", "AA", nodes, heatmap, map[string]bool{}, 26, 26)

	// fmt.Println(order)
	return wuu
}
