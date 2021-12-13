package aoc2021

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type Path []*aoc.Node

func Day12(reader *bufio.Scanner) (string, string) {
	var part1, part2 int
	nodeMap := make(map[string]*aoc.Node)
	for reader.Scan() {
		nodes := strings.Split(reader.Text(), "-")
		for _, node := range nodes {
			if _, ok := nodeMap[node]; !ok {
				nodeMap[node] = aoc.NewNode(node)
			}
		}
		nodeMap[nodes[0]].Connect(nodeMap[nodes[1]])
	}

	start := nodeMap["start"]

	part1 = len(AllPaths(start, []*aoc.Node{start}, false))
	part2 = len(AllPaths(start, []*aoc.Node{start}, true))
	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func AllPaths(node *aoc.Node, curPath []*aoc.Node, canVisitSmall bool) [][]*aoc.Node {
	paths := [][]*aoc.Node{}
	for _, n := range node.Connections {
		if n.Id == "start" {
			// Can only visit start once
			continue
		} else if n.Id == "end" {
			// Path is done
			p := append(curPath, n)
			paths = append(paths, p)
			continue
		} else if strings.ToLower(n.Id) == n.Id && InPath(n, curPath) {
			if canVisitSmall {
				p := append(curPath, n)
				paths = append(paths, AllPaths(n, p, false)...)
			} else {
				// Lowercase nodes can be visited only once, unless canVisitSmall is true
				continue
			}
		} else {
			// Recurse
			p := append(curPath, n)
			paths = append(paths, AllPaths(n, p, canVisitSmall)...)
			continue
		}
	}

	return paths
}

func InPath(n *aoc.Node, path []*aoc.Node) bool {
	for _, node := range path {
		if node.Id == n.Id {
			return true
		}
	}
	return false
}
