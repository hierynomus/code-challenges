package aoc2023

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type Range struct {
	Start  int
	Length int
}

type Mapping struct {
	Source Range
	Dest   Range
}

type MappingMap []Mapping

func (r Range) InRange(i int) bool {
	return i >= r.Start && i < r.Start+r.Length
}

func (m MappingMap) Map(i int) int {
	for _, mm := range m {
		if mm.AppliesTo(i) {
			return mm.Map(i)
		}
	}
	return i
}

func (m MappingMap) Reverse() MappingMap {
	rev := make(MappingMap, len(m))
	for i, mm := range m {
		rev[i] = Mapping{
			Source: mm.Dest,
			Dest:   mm.Source,
		}
	}
	return rev
}

func (m Mapping) AppliesTo(i int) bool {
	return m.Source.InRange(i)
}

func (m Mapping) Map(i int) int {
	return m.Dest.Start + (i - m.Source.Start)
}

func Day05(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	lines := aoc.ReadStringArray(reader)

	seeds := aoc.ToIntArray(strings.Fields(lines[0])[1:])
	mappings := readMappings(lines[2:])

	locs := make([]int, len(seeds))
	for i, s := range seeds {
		locs[i] = s
		for _, mm := range mappings {
			locs[i] = mm.Map(locs[i])
		}
	}

	part1 = aoc.Min(locs)

	part2 = -1
	reverseMappings := make([]MappingMap, len(mappings))
	for i, mm := range mappings {
		reverseMappings[len(reverseMappings)-i-1] = mm.Reverse()
	}

	seedRanges := make([]Range, len(seeds)/2)
	for i := 0; i < len(seeds); i += 2 {
		seedRanges[i/2] = Range{seeds[i], seeds[i+1]}
	}

	for l := 1; l < 100000000; l++ {
		loc := l
		for _, mm := range reverseMappings {
			loc = mm.Map(loc)
		}

		for _, sr := range seedRanges {
			if sr.InRange(loc) {
				part2 = l
				break
			}
		}

		if part2 != -1 {
			break
		}
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func readMappings(lines []string) []MappingMap {
	mappings := []MappingMap{}

	currentMappingMap := MappingMap{}
	for _, l := range lines {
		if l == "" {
			mappings = append(mappings, currentMappingMap)
			currentMappingMap = MappingMap{}
			continue
		}

		if strings.HasSuffix(l, "map:") {
			continue
		}

		nrs := aoc.AsIntArraySpace(l)
		currentMappingMap = append(currentMappingMap, Mapping{
			Dest:   Range{nrs[0], nrs[2]},
			Source: Range{nrs[1], nrs[2]},
		})
	}

	return append(mappings, currentMappingMap)
}
