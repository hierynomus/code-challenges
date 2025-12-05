package aoc2025

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type Range struct {
	Start int
	End   int
}

func NewRange(s string) Range {
	var r Range
	fmt.Sscanf(s, "%d-%d", &r.Start, &r.End)
	return r
}

func (r *Range) Includes(n int) bool {
	return n >= r.Start && n <= r.End
}

func (r *Range) Length() int {
	return r.End - r.Start + 1
}

func Day05(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	lines := aoc.ReadStringArray(reader)
	ranges := []Range{}
	for _, line := range lines {
		if line == "" {
			break
		}
		r := NewRange(line)
		ranges = append(ranges, r)
	}

	for _, line := range lines[len(ranges)+1:] {
		n, _ := strconv.Atoi(line)
		for _, r := range ranges {
			if r.Includes(n) {
				part1++
				break
			}
		}
	}

	sort.Slice(ranges, func(i int, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})

	merged := []Range{}
	for _, r := range ranges {
		if len(merged) == 0 {
			merged = append(merged, r)
			continue
		}
		last := &merged[len(merged)-1]
		if r.Start <= last.End+1 {
			last.End = aoc.MaxOf(last.End, r.End)
		} else {
			merged = append(merged, r)
		}
	}

	for _, r := range merged {
		part2 += r.Length()
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}
