package aoc2022

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type RockPaperScissors int

const (
	ROCK RockPaperScissors = iota
	PAPER
	SCISSORS
)

func Day02(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	moves := [][]rune{}
	lines := aoc.ReadStringArray(reader)
	for _, line := range lines {
		moves = append(moves, []rune{rune(line[0]), rune(line[2])})
	}

	moves1 := [][]RockPaperScissors{}
	for _, m := range moves {
		moves1 = append(moves1, []RockPaperScissors{ConvertMove(m[0]), ConvertMove(m[1])})
	}

	part1 = PlayRockPaperScissors(moves1)

	moves2 := [][]RockPaperScissors{}
	for _, m := range moves {
		p1 := ConvertMove(m[0])
		moves2 = append(moves2, []RockPaperScissors{p1, DecodeGuide(m[1], p1)})
	}

	part2 = PlayRockPaperScissors(moves2)

	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func ConvertMove(r rune) RockPaperScissors {
	switch r {
	case 'A', 'B', 'C':
		return RockPaperScissors(r - 'A')
	case 'X', 'Y', 'Z':
		return RockPaperScissors(r - 'X')
	default:
		panic("invalid move " + string(r))
	}
}

func DecodeGuide(outcome rune, p1 RockPaperScissors) RockPaperScissors {
	return (p1 - (1 - ConvertMove(outcome)) + 3) % 3
}

func PlayRockPaperScissors(moves [][]RockPaperScissors) int {
	score := 0
	for _, m := range moves {
		rps1, rps2 := m[0], m[1]
		score += 1 + int(rps2)
		switch {
		case rps1 == SCISSORS && rps2 == ROCK:
			score += 6
		case rps1 == ROCK && rps2 == PAPER:
			score += 6
		case rps1 == PAPER && rps2 == SCISSORS:
			score += 6
		case rps1 == rps2:
			score += 3
		default:
			continue
		}
	}

	return score
}
