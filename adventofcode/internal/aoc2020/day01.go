package aoc2020

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day01(reader *bufio.Scanner) (string, string) {
	expense := []int{}
	for reader.Scan() {
		expense = append(expense, aoc.ToInt(reader.Text()))
	}

	var part1 int
	for _, c := range aoc.IntCombinations(expense) {
		if c[0]+c[1] == 2020 {
			part1 = c[0] * c[1]
			break
		}
	}

	part2 := 0
	for x := 0; part2 == 0 && x < len(expense)-2; x++ {
		for y := x + 1; part2 == 0 && y < len(expense)-1; y++ {
			for z := y + 1; part2 == 0 && z < len(expense); z++ {
				if expense[x]+expense[y]+expense[z] == 2020 {
					part2 = expense[x] * expense[y] * expense[z]
				}
			}
		}
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}
