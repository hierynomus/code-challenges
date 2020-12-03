package aoc2018

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type LicenseNode struct {
	Children []*LicenseNode
	Metadata []int
}

func NewNode() *LicenseNode {
	return &LicenseNode{
		Children: []*LicenseNode{},
		Metadata: []int{},
	}
}

func (l *LicenseNode) MetaSum() int {
	s := 0
	for _, i := range l.Metadata {
		s += i
	}

	for _, c := range l.Children {
		s += c.MetaSum()
	}

	return s
}

func (l *LicenseNode) RootValue() int {
	v := 0
	if len(l.Children) == 0 {
		for _, i := range l.Metadata {
			v += i
		}
	} else {
		for _, i := range l.Metadata {
			if i > 0 && i <= len(l.Children) {
				v += l.Children[i-1].RootValue()
			}
		}
	}

	return v
}

func Day08(reader *bufio.Scanner) (string, string) {
	if !reader.Scan() {
		panic("no inputß")
	}

	ints := aoc.ToIntArray(strings.Split(reader.Text(), " "))

	node, _ := parseLicenseNode(ints, 0)

	part1 := node.MetaSum()
	part2 := node.RootValue()

	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func parseLicenseNode(ints []int, idx int) (*LicenseNode, int) {
	nrChildren, nrMeta := ints[idx], ints[idx+1]
	node := NewNode()
	idx += 2
	if nrChildren > 0 {
		for x := 0; x < nrChildren; x++ {
			child, indx := parseLicenseNode(ints, idx)
			idx = indx
			node.Children = append(node.Children, child)
		}
	}
	node.Metadata = ints[idx : idx+nrMeta]
	return node, idx + nrMeta
}
