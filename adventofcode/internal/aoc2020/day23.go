package aoc2020

import (
	"bufio"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day23(reader *bufio.Scanner) (string, string) {
	l := aoc.Read(reader)
	game := make([]int, len(l))
	for i, x := range []rune(l) {
		game[i] = int(x - '0')
	}

	result1 := playCrabCups(game, 100)

	return aoc.IntArrayAsString(result1, ""), ""
}

func playCrabCups(game []int, rounds int) []int {
	round := make([]int, len(game))
	copy(round, game)
	for i := 0; i < rounds; i++ {
		current := round[0]
		removed := round[1:4]
		find, foundIdx := current, 0
		for foundIdx == 0 {
			find--
			if find <= 0 {
				find = 9
			}
			for j := 4; j < len(game); j++ {
				if round[j] == find {
					foundIdx = j
					break
				}
			}
		}

		newRound := []int{}
		newRound = append(newRound, round[4:foundIdx+1]...)
		newRound = append(newRound, removed...)
		newRound = append(newRound, round[foundIdx+1:]...)
		newRound = append(newRound, round[0])
		round = newRound
	}

	return round
}
