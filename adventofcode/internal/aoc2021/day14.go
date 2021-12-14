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

	rules := map[string]string{}
	aoc.Skip(reader)

	for reader.Scan() {
		line := reader.Text()
		fromto := strings.Split(line, " -> ")
		rules[fromto[0]] = fromto[1]
	}

	part1 = D14P1(inp, rules)

	pairs := aoc.MakeStringHistogram([]string{})
	for i := 0; i < len(inp)-1; i++ {
		pairs.Add(inp[i : i+2])
	}

	pairRules := map[string][]string{}
	for from, to := range rules {
		pairRules[from] = []string{fmt.Sprintf("%c%s", from[0], to), fmt.Sprintf("%s%c", to, from[1])}
	}

	for i := 0; i < 40; i++ {
		pairs = PolymerTemplatePairs(pairs, pairRules)
	}

	hist := aoc.MakeRuneHistogram([]rune{})
	for k, v := range pairs {
		hist[rune(k[0])] += int64(v)
	}
	hist[rune(inp[len(inp)-1])]++
	_, max := hist.Max()
	_, min := hist.Min()
	part2 = max - min

	return aoc.Int64ToString(part1), aoc.Int64ToString(part2)
}

func D14P1(inp string, rules map[string]string) int64 {
	s := inp
	for i := 0; i < 10; i++ {
		s = PolymerTemplateSimple(s, rules)
	}

	hist := aoc.MakeRuneHistogram([]rune(s))
	_, max := hist.Max()
	_, min := hist.Min()
	return max - min
}

func PolymerTemplateSimple(inp string, rules map[string]string) string {
	s := strings.Builder{}
	for i := 0; i < len(inp)-1; i++ {
		if rules[inp[i:i+2]] != "" {
			s.WriteByte(inp[i])
			s.WriteString(rules[inp[i:i+2]])
		} else {
			s.WriteByte(inp[i])
		}
	}
	s.WriteByte(inp[len(inp)-1])

	return s.String()
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
