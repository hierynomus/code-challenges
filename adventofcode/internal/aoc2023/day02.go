package aoc2023

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type CubeGame struct {
	Nr  int
	Max map[string]int
}

func Day02(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	fit := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	lines := aoc.ReadStringArray(reader)
	games := []CubeGame{}
	for _, line := range lines {
		p := strings.Split(line, ": ")
		gameIdx := aoc.ToInt(p[0][5:])
		game := CubeGame{Nr: gameIdx, Max: map[string]int{}}
		grabs := strings.Split(p[1], "; ")
		for _, turn := range grabs {
			colors := strings.Split(turn, ", ")
			for _, color := range colors {
				c := strings.Split(color, " ")
				cur := aoc.ToInt(c[0])
				if v, ok := game.Max[c[1]]; !ok {
					game.Max[c[1]] = cur
				} else if cur > v {
					game.Max[c[1]] = cur
				}
			}
		}
		games = append(games, game)
	}

	for _, game := range games {
		fits := true
		for color, count := range game.Max {
			if count > fit[color] {
				fits = false
				break
			}
		}

		if fits {
			part1 += game.Nr
		}

		power := 1
		for _, count := range game.Max {
			power *= count
		}

		part2 += power
	}
	return strconv.Itoa(part1), strconv.Itoa(part2)
}
