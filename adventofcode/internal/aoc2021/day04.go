package aoc2021

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type BingoCard [][]string

func (card BingoCard) score() int {
	var score int

	for _, row := range card {
		for _, nr := range row {
			if nr == "X" {
				continue
			}

			score += aoc.ToInt(nr)
		}
	}

	return score
}

func (card BingoCard) strike(inp string) bool {
	for y, row := range card {
		for x, nr := range row {
			if nr == inp {
				card[y][x] = "X"
				return true
			}
		}
	}

	return false
}

func (card BingoCard) isBingo() bool {
	return card.isBingoV() || card.isBingoH()
}

func (card BingoCard) isBingoV() bool {
	var bingo bool
	for x := 0; x < 5; x++ {
		bingo = true
		for _, row := range card {
			if row[x] != "X" {
				bingo = false
				break
			}
		}
		if bingo {
			return true
		}
	}

	return false
}

func (card BingoCard) isBingoH() bool {
	var bingo bool
	for _, row := range card {
		bingo = true
		for _, nr := range row {
			if nr != "X" {
				bingo = false
				break
			}
		}
		if bingo {
			return true
		}
	}

	return false
}

func Day04(reader *bufio.Scanner) (string, string) {
	part1, part2 := 0, 0

	nrs := strings.Split(aoc.Read(reader), ",")

	type Won struct {
		Bingo BingoCard
		Score int
		Turns int
	}

	var wins []Won

	for reader.Scan() {
		reader.Text() // Skip empty
		card := readBingo(reader)

		for i, nr := range nrs {
			if card.strike(nr) {
				if card.isBingo() {
					wins = append(wins, Won{
						Bingo: card,
						Score: card.score() * aoc.ToInt(nr),
						Turns: i,
					})
					break
				}
			}
		}
	}

	minTurns := len(nrs)
	for _, w := range wins {
		if w.Turns < minTurns {
			minTurns = w.Turns
			part1 = w.Score
		}
	}

	maxTurns := 0
	for _, w := range wins {
		if w.Turns > maxTurns {
			maxTurns = w.Turns
			part2 = w.Score
		}
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func readBingo(reader *bufio.Scanner) BingoCard {
	var card BingoCard
	space := regexp.MustCompile(`\s+`)
	for i := 0; i < 5; i++ {
		l := strings.TrimSpace(aoc.Read(reader))
		l = space.ReplaceAllString(l, " ")
		card = append(card, strings.Split(l, " "))
	}

	return card
}
