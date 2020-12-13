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
	buses := []int{}
	if !reader.Scan() {
		panic("missing data")
	}

	for _, b := range strings.Split(reader.Text(), ",") {
		if "x" != b {
			buses = append(buses, aoc.ToInt(b))
		}
	}

	bus := buses[0]
	min := bus - (t % bus)
	for _, b := range buses {
		m := b - (t % b)
		if m < min {
			bus, min = b, m
		}
	}

	return strconv.Itoa(bus * min), ""
}
