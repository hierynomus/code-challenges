package aoc2020

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day10(reader *bufio.Scanner) (string, string) {
	adapters := []int{0}
	adapters = append(adapters, aoc.ReadIntArray(reader)...)
	sort.Ints(adapters)
	adapters = append(adapters, adapters[len(adapters)-1]+3)

	differences := map[int]int{}
	for i := 1; i < len(adapters); i++ {
		d := adapters[i] - adapters[i-1]
		differences[d] += 1
	}

	part1 := differences[1] * differences[3]

	from := 0
	var part2 int64 = 1
	for to := 1; to < len(adapters); to++ {
		if adapters[to]-adapters[to-1] == 3 {
			part2 *= countPermutations(adapters, from, to-1)
			from = to
		}
	}

	return strconv.Itoa(part1), fmt.Sprintf("%d", part2)
}

func countPermutations(adapters []int, from int, to int) int64 {
	switch to - from {
	case 0:
		return 1
	case 1:
		return 1
	case 2:
		return 2
	case 3:
		return 4
	case 4:
		return 7
	default:
		panic(fmt.Sprintf("Unknown permutations for %d", to-from))
	}
}
