package day19

import (
	"bufio"
	"io"
	"math"
	"regexp"
	"strconv"
)

type blueprint struct {
	id        int
	robots    [4][][2]int
	maxNeeded [4]int
}

type balance [4]int

const geode = 3

var (
	blueprintPattern = regexp.MustCompile("Blueprint (?P<id>[0-9]+):")
	orePattern       = regexp.MustCompile("ore robot costs (?P<ore>[0-9]+) ore")
	clayPattern      = regexp.MustCompile("clay robot costs (?P<ore>[0-9]+) ore")
	obsidianPattern  = regexp.MustCompile("obsidian robot costs (?P<ore>[0-9]+) ore and (?P<clay>[0-9]+) clay")
	geodePattern     = regexp.MustCompile("Each geode robot costs (?P<ore>[0-9]+) ore and (?P<obsidian>[0-9]+) obsidian")
)

func max(val ...int) int {
	m := 0
	for _, v := range val {
		if v > m {
			m = v
		}
	}
	return m
}

func parseInput(input io.Reader) []blueprint {
	scanner := bufio.NewScanner(input)
	blueprints := make([]blueprint, 0)
	for scanner.Scan() {
		text := scanner.Text()
		oreCost, _ := strconv.Atoi(orePattern.FindStringSubmatch(text)[orePattern.SubexpIndex("ore")])
		clayCost, _ := strconv.Atoi(clayPattern.FindStringSubmatch(text)[clayPattern.SubexpIndex("ore")])
		obsidianOreCost, _ := strconv.Atoi(obsidianPattern.FindStringSubmatch(text)[obsidianPattern.SubexpIndex("ore")])
		obsidianClayCost, _ := strconv.Atoi(obsidianPattern.FindStringSubmatch(text)[obsidianPattern.SubexpIndex("clay")])
		geodeOreCost, _ := strconv.Atoi(geodePattern.FindStringSubmatch(text)[geodePattern.SubexpIndex("ore")])
		geodeObsidianCost, _ := strconv.Atoi(geodePattern.FindStringSubmatch(text)[geodePattern.SubexpIndex("obsidian")])
		blueprintMatch := blueprintPattern.FindStringSubmatch(text)
		id, _ := strconv.Atoi(blueprintMatch[blueprintPattern.SubexpIndex("id")])
		robots := [4][][2]int{
			{{oreCost, 0}},
			{{clayCost, 0}},
			{{obsidianOreCost, 0}, {obsidianClayCost, 1}},
			{{geodeOreCost, 0}, {geodeObsidianCost, 2}},
		}
		maxOre := max(oreCost, clayCost, obsidianOreCost)
		maxNeeded := [4]int{maxOre, obsidianClayCost, geodeObsidianCost, math.MaxInt}
		blueprints = append(blueprints, blueprint{id, robots, maxNeeded})
	}
	return blueprints
}

func createCacheKey(bal balance, robots [4]int, minutes int) [9]int {
	key := [9]int{}
	copy(key[:], bal[:])
	copy(key[2:], robots[:])
	key[8] = minutes
	return key
}

func dfs(bp blueprint, bal balance, robots [4]int, cache map[[9]int]int, minutes int) int {
	if minutes == 0 {
		return bal[geode]
	}

	key := createCacheKey(bal, robots, minutes)
	if v, ok := cache[key]; ok {
		return v
	}

	// Default max to doing nothing
	max := bal[geode] + (robots[geode] * minutes)
	for robotType, robotResources := range bp.robots {
		// Check if we want any more of these robots.
		// We always want more geode robots
		if robotType != geode && robots[robotType] >= bp.maxNeeded[robotType] {
			continue
		}

		wait, canBuildBot := 0, true
		for _, resourceNeeded := range robotResources {
			resourceCost, resourceType := resourceNeeded[0], resourceNeeded[1]
			if robots[resourceType] == 0 {
				canBuildBot = false
				break
			}
			if bal[resourceType] < resourceCost {
				value, timetoWait := bal[resourceType], 0
				for resourceCost > value {
					timetoWait++
					value += robots[resourceType]
				}
				wait = int(math.Max(float64(wait), float64(timetoWait)))
			}
		}

		if canBuildBot {
			remTime := minutes - wait - 1
			if remTime <= 0 {
				continue
			}
			newBal, newRobots := balance{}, [4]int{}
			copy(newBal[:], bal[:])
			copy(newRobots[:], robots[:])
			for i := range newBal {
				newBal[i] = newBal[i] + (newRobots[i] * (wait + 1))
			}
			for _, resourceNeeded := range robotResources {
				resourceCost, resourceType := resourceNeeded[0], resourceNeeded[1]
				newBal[resourceType] -= resourceCost
			}
			newRobots[robotType]++
			res := dfs(bp, newBal, newRobots, cache, remTime)
			if res > max {
				max = res
			}
		}
	}
	cache[key] = max
	return max
}

func PartOne(input io.Reader) int {
	blueprints := parseInput(input)
	total := 0
	for _, bp := range blueprints {
		total = total + bp.id*dfs(bp, balance{}, [4]int{1, 0, 0, 0}, map[[9]int]int{}, 24)
	}
	return total
}

func PartTwo(input io.Reader) int {
	blueprints := parseInput(input)
	total := 1
	for _, bp := range blueprints {
		total *= dfs(bp, balance{}, [4]int{1, 0, 0, 0}, map[[9]int]int{}, 32)
	}
	return total
}
