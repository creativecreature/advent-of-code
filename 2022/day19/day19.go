package day19

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"regexp"
	"strconv"
)

type ResourceType uint8

const (
	Ore ResourceType = iota
	Clay
	Obsidian
	Geode
)

type cost struct {
	resourceType ResourceType
	value        int
}

type robot struct {
	cost     []cost
	produces ResourceType
}

type blueprint struct {
	id     int
	robots []robot
}

type balance struct {
	ores     int
	clay     int
	obsidian int
	geode    int
}

var (
	blueprintPattern     = regexp.MustCompile("Blueprint (?P<id>[0-9]+):")
	oreRobotPattern      = regexp.MustCompile("ore robot costs (?P<ore>[0-9]+) ore")
	clayRobotPattern     = regexp.MustCompile("clay robot costs (?P<ore>[0-9]+) ore")
	obsidianRobotPattern = regexp.MustCompile("obsidian robot costs (?P<ore>[0-9]+) ore and (?P<clay>[0-9]+) clay")
	geodeRobotPattern    = regexp.MustCompile("Each geode robot costs (?P<ore>[0-9]+) ore and (?P<obsidian>[0-9]+) obsidian")
)

func parseInput(input io.Reader) []blueprint {
	scanner := bufio.NewScanner(input)
	blueprints := make([]blueprint, 0)
	for scanner.Scan() {
		text := scanner.Text()
		oreRobotMatch := oreRobotPattern.FindStringSubmatch(text)
		oreRobotCost, _ := strconv.Atoi(oreRobotMatch[oreRobotPattern.SubexpIndex("ore")])
		oreRobot := robot{cost: []cost{{Ore, oreRobotCost}}, produces: Ore}
		clayRobotMatch := clayRobotPattern.FindStringSubmatch(text)
		clayRobotCost, _ := strconv.Atoi(clayRobotMatch[clayRobotPattern.SubexpIndex("ore")])
		clayRobot := robot{cost: []cost{{Ore, clayRobotCost}}, produces: Clay}
		obsidianRobotMatch := obsidianRobotPattern.FindStringSubmatch(text)
		obsidianRobotOreCost, _ := strconv.Atoi(obsidianRobotMatch[obsidianRobotPattern.SubexpIndex("ore")])
		obsidianRobotClayCost, _ := strconv.Atoi(obsidianRobotMatch[obsidianRobotPattern.SubexpIndex("ore")])
		obsidianRobot := robot{cost: []cost{{Ore, obsidianRobotOreCost}, {Clay, obsidianRobotClayCost}}, produces: Obsidian}
		geodeRobotMatch := geodeRobotPattern.FindStringSubmatch(text)
		geodeRobotCost, _ := strconv.Atoi(geodeRobotMatch[geodeRobotPattern.SubexpIndex("ore")])
		geodeRobot := robot{cost: []cost{{Obsidian, geodeRobotCost}}, produces: Geode}
		blueprintMatch := blueprintPattern.FindStringSubmatch(text)
		id, _ := strconv.Atoi(blueprintMatch[blueprintPattern.SubexpIndex("id")])
		blueprints = append(blueprints, blueprint{id, []robot{oreRobot, clayRobot, obsidianRobot, geodeRobot}})
	}
	return blueprints
}

func findBestPath(bp blueprint, minutes int, bal balance, robots, queue []robot) int {
	if minutes < 1 {
		return bal.geode
	}

	// Add whatever the robots produced to the balance
	for _, robot := range robots {
		switch robot.produces {
		case Ore:
			bal.ores++
		case Clay:
			bal.clay++
		case Obsidian:
			bal.obsidian++
		case Geode:
			bal.geode++
		default:
			panic("Produces unknown resource type")
		}
	}

	if len(queue) > 0 {
		robotToProduce := queue[0]
		canBuildRobot := true
		for _, cost := range robotToProduce.cost {
			switch cost.resourceType {
			case Ore:
				canBuildRobot = bal.ores >= cost.value
			case Clay:
				canBuildRobot = bal.clay >= cost.value
			case Obsidian:
				canBuildRobot = bal.obsidian >= cost.value
			default:
				panic("Unknown resource type")
			}
			if !canBuildRobot {
				break
			}
		}

		if canBuildRobot {
			for _, cost := range robotToProduce.cost {
				switch cost.resourceType {
				case Ore:
					bal.ores = bal.ores - cost.value
				case Clay:
					bal.clay = bal.clay - cost.value
				case Obsidian:
					bal.clay = bal.clay - cost.value
				}
			}
			robots = append(robots, robotToProduce)
			queue = queue[1:]
		}
	}

	bestChild := 0
	for _, r := range bp.robots {
		newRobots := []robot{}
		newRobots = append(newRobots, robots...)
		newQueue := []robot{}
		newQueue = append(newQueue, queue...)
		newQueue = append(newQueue, r)
		res := findBestPath(bp, minutes-1, bal, newRobots, newQueue)
		bestChild = int(math.Max(float64(bestChild), float64(res)))
	}

	return bal.geode + bestChild
}

func PartOne(input io.Reader) int {
	blueprints := parseInput(input)
	firstOreRobot := robot{
		produces: Ore,
	}
	res := findBestPath(blueprints[0], 24, balance{}, []robot{firstOreRobot}, []robot{})
	fmt.Println(res)
	return -1
}
