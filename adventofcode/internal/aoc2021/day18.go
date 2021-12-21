package aoc2021

import (
	"bufio"
	"fmt"
	"math"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type SnailFishNumber struct {
	parent, left, right *SnailFishNumber
	value               int
}

func (p *SnailFishNumber) String() string {
	if p.left == nil && p.right == nil {
		return strconv.Itoa(p.value)
	}

	return fmt.Sprintf("[%v,%v]", p.left, p.right)
}

func (p *SnailFishNumber) IsPair() bool {
	return p.left != nil && p.right != nil
}

func (p *SnailFishNumber) Explode() bool {
	// Can only explode pairs
	if !p.IsPair() {
		return false
	}

	// Can only explode if left and right are regular numbers
	regulars := !p.left.IsPair() && !p.right.IsPair()
	fourLevels := p.parent != nil && p.parent.parent != nil && p.parent.parent.parent != nil && p.parent.parent.parent.parent != nil

	if regulars && fourLevels {
		p.addLeftValue(p.left.value)
		p.addRightValue(p.right.value)
		p.left, p.right, p.value = nil, nil, 0
		return true
	} else {
		if p.left.Explode() {
			return true
		} else {
			return p.right.Explode()
		}
	}
}

func (p *SnailFishNumber) addLeftValue(v int) {
	parent, x := p.parent, p
	for parent.parent != nil && parent.left == x {
		parent, x = parent.parent, parent
	}

	// Left-most leaf node gets lost
	if parent.parent == nil && parent.left == x {
		return
	}

	x = parent.left
	for x.IsPair() {
		x = x.right
	}

	x.value += v
}

func (p *SnailFishNumber) addRightValue(v int) {
	parent, x := p.parent, p
	for parent.parent != nil && parent.right == x {
		parent, x = parent.parent, parent
	}

	// Right-most leaf node gets lost
	if parent.parent == nil && parent.right == x {
		return
	}

	x = parent.right
	for x.IsPair() {
		x = x.left
	}

	x.value += v
}

func (p *SnailFishNumber) Split() bool {
	if p.IsPair() {
		if p.left.Split() {
			return true
		} else {
			return p.right.Split()
		}
	}

	if p.value >= 10 {
		p.left = &SnailFishNumber{value: int(math.Floor(float64(p.value) / 2)), parent: p}
		p.right = &SnailFishNumber{value: int(math.Ceil(float64(p.value) / 2)), parent: p}
		p.value = 0
		return true
	}

	return false
}

func (p *SnailFishNumber) Reduce() {
	for p.Explode() || p.Split() {
	}
}

func (p *SnailFishNumber) Add(n *SnailFishNumber) *SnailFishNumber {
	x := &SnailFishNumber{left: p, right: n, parent: nil}
	p.parent, n.parent = x, x
	x.Reduce()
	return x
}

func (p *SnailFishNumber) Magnitude() int64 {
	if p.IsPair() {
		return 3*p.left.Magnitude() + 2*p.right.Magnitude()
	} else {
		return int64(p.value)
	}
}

func ParseSnailFish(inp string, idx int, parent *SnailFishNumber) (*SnailFishNumber, int) {
	if inp[idx] == '[' {
		p := &SnailFishNumber{parent: parent}
		p.left, idx = ParseSnailFish(inp, idx+1, p)

		if inp[idx] != ',' {
			panic("expected ','")
		}

		p.right, idx = ParseSnailFish(inp, idx+1, p)

		if inp[idx] == ']' {
			return p, idx + 1
		}

		panic("Expected ']' as closing pair")
	} else {
		s := []byte{inp[idx]}

		for inp[idx+1] != ',' && inp[idx+1] != ']' {
			s = append(s, inp[idx+1])
			idx++
		}

		return &SnailFishNumber{value: aoc.ToInt(string(s)), parent: parent}, idx + 1
	}
}

func Day18(reader *bufio.Scanner) (string, string) {
	var part1, part2 int64

	lines := aoc.ReadStringArray(reader)
	var p *SnailFishNumber
	for _, line := range lines {
		x, _ := ParseSnailFish(line, 0, nil)
		if p != nil {
			p = p.Add(x)
		} else {
			p = x
		}
	}

	part1 = p.Magnitude()

	for _, l := range lines {
		for _, l2 := range lines {
			x, _ := ParseSnailFish(l, 0, nil)
			y, _ := ParseSnailFish(l2, 0, nil)
			p = x.Add(y)
			m := p.Magnitude()
			if m > part2 {
				part2 = m
			}
		}
	}

	return aoc.Int64ToString(part1), aoc.Int64ToString(part2)
}
