package aoc2021

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day14(reader *bufio.Scanner) (string, string) {
	var part1, part2 int64

	inp := aoc.Read(reader)
	pairs := aoc.MakeStringHistogram([]string{})
	for i := 0; i < len(inp)-1; i++ {
		pairs.Add(inp[i : i+2])
	}

	pairRules := map[string][]string{}
	aoc.Skip(reader)

	for reader.Scan() {
		line := reader.Text()
		fromto := strings.Split(line, " -> ")
		pairRules[fromto[0]] = []string{fmt.Sprintf("%c%s", fromto[0][0], fromto[1]), fmt.Sprintf("%s%c", fromto[1], fromto[0][1])}
	}

	for i := 0; i < 10; i++ {
		pairs = PolymerTemplatePairs(pairs, pairRules)
	}

	hist := CountOccurrences(pairs)
	hist[rune(inp[len(inp)-1])]++
	part1 = hist[hist.Max()] - hist[hist.Min()]

	for i := 10; i < 40; i++ {
		pairs = PolymerTemplatePairs(pairs, pairRules)
	}

	hist = CountOccurrences(pairs)
	hist[rune(inp[len(inp)-1])]++
	part2 = hist[hist.Max()] - hist[hist.Min()]

	return aoc.Int64ToString(part1), aoc.Int64ToString(part2)
}

func CountOccurrences(pairs aoc.StringHistogram) aoc.RuneHistogram {
	hist := aoc.MakeRuneHistogram([]rune{})
	for k, v := range pairs {
		hist[rune(k[0])] += int64(v)
	}

	return hist
}

func PolymerTemplatePairs(inp aoc.StringHistogram, rules map[string][]string) aoc.StringHistogram {
	pairs := aoc.MakeStringHistogram([]string{})
	for k, v := range inp {
		if out, ok := rules[k]; ok {
			for _, o := range out {
				pairs[o] += v
			}
		} else {
			pairs[k] = v
		}
	}

	return pairs
}
