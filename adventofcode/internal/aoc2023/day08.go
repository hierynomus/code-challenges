package aoc2023

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type DesertNode struct {
	Name  string
	Left  *DesertNode
	Right *DesertNode
}

func (d *DesertNode) String() string {
	return fmt.Sprintf("%s (%s, %s)", d.Name, d.Left.Name, d.Right.Name)
}

func Day08(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	lines := aoc.ReadStringArray(reader)
	instructions := lines[0]

	nodes := make(map[string]*DesertNode)
	for _, node := range lines[2:] {
		parse := strings.Fields(node)
		n := parse[0]
		if _, ok := nodes[n]; !ok {
			nodes[n] = &DesertNode{Name: n}
		}

		l := parse[2][1 : len(parse[2])-1]
		if _, ok := nodes[l]; !ok {
			nodes[l] = &DesertNode{Name: l}
		}

		r := parse[3][:len(parse[3])-1]
		if _, ok := nodes[r]; !ok {
			nodes[r] = &DesertNode{Name: r}
		}

		nodes[n].Left = nodes[l]
		nodes[n].Right = nodes[r]
	}

	// fmt.Printf("%v", nodes)

	curr := nodes["AAA"]
	for curr.Name != "ZZZ" {
		for _, i := range instructions {
			if curr.Name == "ZZZ" {
				break
			}
			switch i {
			case 'L':
				curr = curr.Left
			case 'R':
				curr = curr.Right
			}
			part1++
		}
	}
	return strconv.Itoa(part1), strconv.Itoa(part2)
}
