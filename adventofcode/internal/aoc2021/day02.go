package aoc2021

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day02(reader *bufio.Scanner) (string, string) {
	part1, part2 := 0, 0

	pos, depth := 0, 0
	aim, ps, dpth := 0, 0, 0
	commands := aoc.ReadStringArray(reader)
	for _, command := range commands {
		c := strings.Split(command, " ")
		if c[0] == "forward" {
			pos += aoc.ToInt(c[1])
			ps, dpth = ps+aoc.ToInt(c[1]), dpth+(aim*aoc.ToInt(c[1]))
		} else if c[0] == "down" {
			depth += aoc.ToInt(c[1])
			aim += aoc.ToInt(c[1])
		} else if c[0] == "up" {
			depth -= aoc.ToInt(c[1])
			aim -= aoc.ToInt(c[1])
		}
	}
	part1 = pos * depth
	part2 = ps * dpth
	return strconv.Itoa(part1), strconv.Itoa(part2)
}
