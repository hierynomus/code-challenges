package aoc2023

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func BoatDistance(raceTime, pressTime int) int {
	return pressTime * (raceTime - pressTime)
}

func NrValidPressTimes(raceTime int, distanceToBeat int) int {
	t := aoc.BinSearchInt(0, raceTime/2, func(t int) bool {
		return BoatDistance(raceTime, t) > distanceToBeat
	})

	valid := 0
	if raceTime%2 == 0 {
		valid = ((raceTime/2)-t)*2 + 1
	} else {
		valid = (((raceTime + 1) / 2) - t) * 2
	}

	return valid
}

func Day06(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	lines := aoc.ReadStringArray(reader)
	times := aoc.ToIntArray(strings.Fields(lines[0])[1:])
	distances := aoc.ToIntArray(strings.Fields(lines[1])[1:])

	part1 = 1
	for i := 0; i < len(times); i++ {
		part1 *= NrValidPressTimes(times[i], distances[i])
	}

	p2Time := aoc.ToInt(strings.Split(strings.ReplaceAll(lines[0], " ", ""), ":")[1])
	p2Dist := aoc.ToInt(strings.Split(strings.ReplaceAll(lines[1], " ", ""), ":")[1])
	part2 = NrValidPressTimes(p2Time, p2Dist)

	return strconv.Itoa(part1), strconv.Itoa(part2)
}
