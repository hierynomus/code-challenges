package aoc2023

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type RockField []string

func (r RockField) String() string {
	return strings.Join(r, "\n")
}

func Day13(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	fields := []RockField{}
	curr := RockField{}
	for reader.Scan() {
		line := reader.Text()
		if line == "" {
			fields = append(fields, curr)
			curr = RockField{}
			continue
		}

		curr = append(curr, line)
	}
	fields = append(fields, curr)

	for _, field := range fields {
		// yMirror, xMirror := 0, 0
		for y := 0; y < len(field)-1; y++ {
			if checkMirror(field, y, y+1) {
				part1 += 100 * (y + 1)
			}
			if checkSmudgedMirror(field, y, y+1, 0) {
				part2 += 100 * (y + 1)
			}
		}

		t := RockField(aoc.TransposeString(field))
		for x := 0; x < len(t)-1; x++ {
			if checkMirror(t, x, x+1) {
				part1 += x + 1
			}
			if checkSmudgedMirror(t, x, x+1, 0) {
				part2 += x + 1
			}
		}
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func checkMirror(field RockField, y1, y2 int) bool {
	if y1 < 0 || y2 >= len(field) {
		return true
	}
	return field[y1] == field[y2] && checkMirror(field, y1-1, y2+1)
}

func checkSmudgedMirror(field RockField, y1, y2, smudges int) bool {
	if y1 < 0 || y2 >= len(field) {
		return smudges == 1
	}

	for x := 0; x < len(field[y1]); x++ {
		if field[y1][x] != field[y2][x] {
			smudges++
		}
	}

	return smudges <= 1 && checkSmudgedMirror(field, y1-1, y2+1, smudges)
}
