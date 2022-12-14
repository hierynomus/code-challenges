package aoc2022

import (
	"bufio"
	"sort"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type Signal interface {
	Compare(other Signal) int
}

type SignalList []Signal

func (l SignalList) Compare(other Signal) int {
	ol, ok := other.(SignalList)
	if !ok {
		ol = SignalList{other}
	}

	for i := 0; i < len(l); i++ {
		if i >= len(ol) {
			return 1
		} else {
			cmp := l[i].Compare(ol[i])
			if cmp != 0 {
				return cmp
			}
		}
	}

	if len(l) < len(ol) {
		return -1
	}
	return 0
}

func (l SignalList) Parse(s string, idx int) (SignalList, int) {
	for idx < len(s) {
		switch {
		case s[idx] == '[':
			ll, i := SignalList{}.Parse(s, idx+1)
			idx = i
			l = append(l, ll)
		case s[idx] == ']':
			return l, idx + 1
		case s[idx] == ',':
			idx++
		default:
			end := idx
			for x := idx; x < len(s); x++ {
				if s[x] == ',' || s[x] == ']' {
					end = x
					break
				}
			}

			s := SignalInt(aoc.ToInt(s[idx:end]))
			l = append(l, s)
			idx = end
		}
	}

	return l, idx
}

type SignalInt int

func (i SignalInt) Compare(other Signal) int {
	oi, ok := other.(SignalInt)
	if ok {
		if i < oi {
			return -1
		} else if i > oi {
			return 1
		}
		return 0
	}

	return SignalList{i}.Compare(other)
}

var _ Signal = SignalInt(0)
var _ Signal = SignalList{}

func Day13(reader *bufio.Scanner) (string, string) {
	var part1, part2 int
	inp := aoc.ReadStringArray(reader)
	allSignals := []Signal{}
	p := 1

	for i := 0; i < len(inp); i += 3 {
		left, _ := SignalList{}.Parse(inp[i], 1)
		right, _ := SignalList{}.Parse(inp[i+1], 1)
		if left.Compare(right) <= 0 {
			part1 += p
		}
		p++
		allSignals = append(allSignals, left, right)
	}

	divider1, _ := SignalList{}.Parse("[[2]]", 1)
	divider2, _ := SignalList{}.Parse("[[6]]", 1)
	allSignals = append(allSignals, divider1, divider2)

	sort.SliceStable(allSignals, func(i, j int) bool {
		return allSignals[i].Compare(allSignals[j]) < 0
	})

	d1, d2 := 0, 0
	for i := 0; i < len(allSignals); i++ {
		if allSignals[i].Compare(divider1) == 0 {
			d1 = i + 1
		} else if allSignals[i].Compare(divider2) == 0 {
			d2 = i + 1
		}
	}

	part2 = d1 * d2

	return strconv.Itoa(part1), strconv.Itoa(part2)
}
