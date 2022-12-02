package aoc2020

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day13(reader *bufio.Scanner) (string, string) {
	if !reader.Scan() {
		panic("Nothing to read")
	}
	t := aoc.ToInt(reader.Text())
	buses := map[int]int{}
	if !reader.Scan() {
		panic("missing data")
	}

	for i, b := range strings.Split(reader.Text(), ",") {
		if b != "x" {
			buses[aoc.ToInt(b)] = i
		}
	}

	var bus, min int
	for b := range buses {
		m := b - (t % b)
		if bus == 0 || m < min {
			bus, min = b, m
		}
	}

	t, period := 0, 1
	for b, i := range buses {
		for {
			if (t+i)%b == 0 {
				period *= b
				break
			}
			t += period
		}
	}

	part2 := t

	return strconv.Itoa(bus * min), strconv.Itoa(part2)
}
