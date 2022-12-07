package day07

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
)

var DirPattern = regexp.MustCompile("dir ([a-z]+)")
var FilePattern = regexp.MustCompile("([0-9]+) .+")
var CdPattern = regexp.MustCompile("cd (.+)")

type Node struct {
	Size     int
	Parent   *Node
	Children map[string]*Node
}

func BuildTree(input io.Reader) *Node {
	root := &Node{Children: make(map[string]*Node)}
	current := root
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		command := scanner.Text()
		if match := DirPattern.FindSubmatch([]byte(command)); match != nil {
			current.Children[string(match[1])] = &Node{
				Parent:   current,
				Children: make(map[string]*Node),
			}
		} else if match := FilePattern.FindSubmatch([]byte(command)); match != nil {
			fileSize, err := strconv.Atoi(string(match[1]))
			if err == nil {
				current.Size = current.Size + fileSize
			}
		} else if match := CdPattern.FindSubmatch([]byte(command)); match != nil {
			target := string(match[1])
			if target == ".." {
				current = current.Parent
			} else {
				current = current.Children[target]
			}
		}
	}
	return root
}

func calculateTreeSize(node *Node) int {
	childrenSize := 0
	for _, c := range node.Children {
		childrenSize += calculateTreeSize(c)
	}
	node.Size += childrenSize
	return node.Size
}

func getDirectoriesUnderThreshold(node *Node, count, threshold int) int {
	for _, c := range node.Children {
		count = getDirectoriesUnderThreshold(c, count, threshold)
		if c.Size <= threshold {
			count += c.Size
		}
	}
	return count
}

func FindNodeToDelete(node *Node, spaceToFreeUp, bestSoFar int) int {
	for _, c := range node.Children {
		if c.Size >= spaceToFreeUp {
			bestSoFar = FindNodeToDelete(c, spaceToFreeUp, bestSoFar)
		}
	}

	if node.Size >= spaceToFreeUp && node.Size < bestSoFar {
		bestSoFar = node.Size
	}
	return bestSoFar
}

func PartOne(reader io.Reader) int {
	root := BuildTree(reader)
	calculateTreeSize(root)
	return getDirectoriesUnderThreshold(root, 0, 100000)
}

func PartTwo(reader io.Reader) int {
	root := BuildTree(reader)
	rootSize := calculateTreeSize(root)
	totalDiskSpace, requiredMemory := 70000000, 30000000
	spaceToFreeUp := requiredMemory - (totalDiskSpace - rootSize)
	size := FindNodeToDelete(root, spaceToFreeUp, root.Size)
	return size
}
