package aoc2022

import (
	"bufio"
	"sort"
	"strconv"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type MonkeyAway struct {
	Operation   func(old int64) int64
	Items       []int64
	Divisor     int64
	TrueMonkey  int
	FalseMonkey int
}

func ParseMonkey(reader *bufio.Scanner) MonkeyAway {
	monkey := MonkeyAway{
		Items: []int64{},
	}
	itemLine := aoc.Read(reader)
	for _, item := range strings.Split(itemLine[18:], ", ") {
		i := aoc.ToInt64(item)
		monkey.Items = append(monkey.Items, i)
	}
	monkey.Operation = ParseOperation(strings.TrimSpace(aoc.Read(reader)))
	monkey.Divisor = ParseDivisor(strings.TrimSpace(aoc.Read(reader)))
	monkey.TrueMonkey = aoc.ToInt(strings.Split(strings.TrimSpace(aoc.Read(reader)), " ")[5])
	monkey.FalseMonkey = aoc.ToInt(strings.Split(strings.TrimSpace(aoc.Read(reader)), " ")[5])
	return monkey
}

func ParseOperation(line string) func(old int64) int64 {
	if !strings.HasPrefix(line, "Operation:") {
		panic("Not an operation")
	}

	s := strings.Split(line, " ")
	switch s[4] {
	case "*":
		if s[5] == "old" {
			return func(old int64) int64 {
				return old * old
			}
		} else {
			i := aoc.ToInt64(s[5])
			return func(old int64) int64 {
				return old * i
			}
		}
	case "+":
		if s[5] == "old" {
			return func(old int64) int64 {
				return old + old
			}
		} else {
			i := aoc.ToInt64(s[5])
			return func(old int64) int64 {
				return old + i
			}
		}
	default:
		panic("Unknown operation")
	}
}

func ParseDivisor(line string) int64 {
	if !strings.HasPrefix(line, "Test:") {
		panic("Not a test")
	}

	s := strings.Split(line, " ")
	x := aoc.ToInt64(s[3])

	return x
}

func PlayMakeAway(monkeys map[int]MonkeyAway, rounds int, worryReduction func(int64) int64) int {
	items := map[int][]int64{}
	for i := 0; i < len(monkeys); i++ {
		items[i] = []int64{}
		items[i] = append(items[i], monkeys[i].Items...)
	}

	actions := make([]int, len(monkeys))
	for i := 0; i < rounds; i++ {
		for m := 0; m < len(monkeys); m++ {
			monkey := monkeys[m]
			for _, item := range items[m] {
				item = monkey.Operation(item)
				item = worryReduction(item)
				if item%monkey.Divisor == 0 {
					items[monkey.TrueMonkey] = append(items[monkey.TrueMonkey], item)
				} else {
					items[monkey.FalseMonkey] = append(items[monkey.FalseMonkey], item)
				}
				actions[m]++
			}
			items[m] = []int64{}
		}
	}

	sort.Ints(actions)
	return actions[len(actions)-1] * actions[len(actions)-2]
}

func Day11(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	monkeys := map[int]MonkeyAway{}
	for reader.Scan() {
		l := reader.Text()
		if !strings.HasPrefix(l, "Monkey ") {
			continue
		}

		m := ParseMonkey(reader)
		monkeys[aoc.ToInt(l[7:8])] = m
	}

	part1 = PlayMakeAway(monkeys, 20, func(i int64) int64 {
		return i / 3
	})

	divs := []int64{}
	for _, m := range monkeys {
		divs = append(divs, m.Divisor)
	}
	lcm := aoc.LcmArray(divs)

	part2 = PlayMakeAway(monkeys, 10000, func(i int64) int64 {
		if i > lcm {
			return i % lcm
		} else {
			return i
		}
	})

	return strconv.Itoa(part1), strconv.Itoa(part2)
}
