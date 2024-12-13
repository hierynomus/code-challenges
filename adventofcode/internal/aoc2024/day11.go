package aoc2024

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type Stone struct {
	N     *Stone
	Value int
}

func (s *Stone) Blink() *Stone {
	nextStone := s.N
	nd := numDigits(s.Value)
	if s.Value == 0 {
		s.Value = 1
	} else if nd%2 == 0 {
		// Replace with 2 stones that have the first part of the digits and the last part of the digits
		fh, lh := s.Value/aoc.Pow(10, nd/2), s.Value%aoc.Pow(10, nd/2)
		s.Value = fh
		s.N = &Stone{N: s.N, Value: lh}
	} else {
		s.Value *= 2024
	}

	// Returning the original nextstone
	return nextStone
}

func numDigits(n int) int {
	if n > 0 {
		return 1 + numDigits(n/10)
	}
	return 0
}

func Day11(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	arr := aoc.ReadIntArrayLine(reader, " ")
	_ = createStones(arr) // dumb idea

	stones := map[int]int{}
	for _, s := range arr {
		stones[s]++
	}

	for i := 0; i < 25; i++ {
		stones = blink(stones)
	}

	for _, v := range stones {
		part1 += v
	}

	for i := 0; i < 50; i++ {
		stones = blink(stones)
	}

	for _, v := range stones {
		part2 += v
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func blink(stones map[int]int) map[int]int {
	newStones := map[int]int{}
	for s, v := range stones {
		digits := numDigits(s)
		if s == 0 {
			newStones[1] += v
		} else if digits%2 == 0 {
			fh, lh := s/aoc.Pow(10, digits/2), s%aoc.Pow(10, digits/2)
			newStones[fh] += v
			newStones[lh] += v
		} else {
			newStones[s*2024] += v
		}
	}
	return newStones
}

func createStones(arr []int) *Stone {
	head := &Stone{Value: arr[0]}
	curr := head
	for i := 1; i < len(arr); i++ {
		curr.N = &Stone{Value: arr[i]}
		curr = curr.N
	}
	return head
}
