package aoc2021

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day03(reader *bufio.Scanner) (string, string) {
	var part1, part2 int64

	nrs := aoc.ReadStringArray(reader)
	gamma, epsilon := "", ""
	for x := 0; x < len(nrs[0]); x++ {
		nr0, nr1 := 0, 0
		for y := 0; y < len(nrs); y++ {
			if nrs[y][x] == '0' {
				nr0++
			} else {
				nr1++
			}
		}
		if nr1 > nr0 {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}
	part1 = aoc.ParseBin(gamma) * aoc.ParseBin(epsilon)

	oxygen := filterBitByBit(nrs, 0, func(ones, zeros []string) bool { return len(ones) >= len(zeros) })
	scrubber := filterBitByBit(nrs, 0, func(ones, zeros []string) bool { return len(ones) < len(zeros) })

	part2 = aoc.ParseBin(oxygen) * aoc.ParseBin(scrubber)

	return strconv.FormatInt(part1, 10), strconv.FormatInt(part2, 10)
}

func filterBitByBit(nrs []string, pos int, cmp func([]string, []string) bool) string {
	if len(nrs) <= 1 {
		return nrs[0]
	}

	ones, zeros := partitionByBit(nrs, pos)

	if cmp(ones, zeros) {
		return filterBitByBit(ones, pos+1, cmp)
	} else {
		return filterBitByBit(zeros, pos+1, cmp)
	}
}

func partitionByBit(nrs []string, pos int) ([]string, []string) {
	ones, zeros := []string{}, []string{}
	for _, o := range nrs {
		if o[pos] == '1' {
			ones = append(ones, o)
		} else {
			zeros = append(zeros, o)
		}
	}

	return ones, zeros
}
