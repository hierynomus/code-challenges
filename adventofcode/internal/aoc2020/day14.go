package aoc2020

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day14(reader *bufio.Scanner) (string, string) {
	lines := aoc.ReadStringArray(reader)
	part1 := emulateV1Chip(lines)
	return fmt.Sprintf("%d", part1), ""
}

func parseMask(l string) (int64, int64) {
	return aoc.ParseBin(strings.ReplaceAll(l, "X", "1")),
		aoc.ParseBin(strings.ReplaceAll(l, "X", "0"))
}

func emulateV1Chip(lines []string) int64 {
	mem := map[int]int64{}
	var andMask, orMask int64
	for _, l := range lines {
		if strings.HasPrefix(l, "mask = ") {
			andMask, orMask = parseMask(l[7:])
		} else {
			kv := strings.Split(l, " = ")
			mem[aoc.ToInt(kv[0][4:len(kv[0])-1])] = aoc.ToInt64(kv[1])&andMask | orMask
		}
	}

	var sum int64 = 0
	for _, v := range mem {
		sum += v
	}
	return sum
}

// func emulateV2Chipe(lines []string) int64 {
// 	mem := map[int64]int64{}
// 	var mask string
// 	for _, l := range lines {
// 		if strings.HasPrefix(l, "mask = ") {
// 			mask = l[7:]
// 		} else {
// 			kv := strings.Split(l, " = ")
// 			k, v := aoc.ToInt64(kv[0][4:len(kv[0])-1]), aoc.ToInt64(kv[1])

// 		}
// 	}
// }
