package aoc2020

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

var emptyDeck = []int{}

func Day22(reader *bufio.Scanner) (string, string) {
	deck1 := readDeck(reader)
	deck2 := readDeck(reader)

	player1, player2 := playCombat(deck1, deck2)
	part1 := 0
	if len(player1) > 0 {
		part1 = score(player1)
	} else {
		part1 = score(player2)
	}

	player1, player2 = playRecursiveCombat(deck1, deck2)
	part2 := 0
	if len(player1) > 0 {
		part2 = score(player1)
	} else {
		part2 = score(player2)
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func playCombat(deck1, deck2 []int) ([]int, []int) {
	player1 := make([]int, len(deck1))
	copy(player1, deck1)
	player2 := make([]int, len(deck2))
	copy(player2, deck2)

	for len(player1) > 0 && len(player2) > 0 {
		c1, c2 := player1[0], player2[0]
		player1, player2 = player1[1:], player2[1:]
		if c1 > c2 {
			player1 = append(player1, c1, c2)
		} else {
			player2 = append(player2, c2, c1)
		}
	}

	return player1, player2
}

func playRecursiveCombat(deck1, deck2 []int) ([]int, []int) {
	seen := map[string]bool{}
	for len(deck1) > 0 && len(deck2) > 0 {
		check := strings.Join([]string{checksum(deck1), checksum(deck2)}, "|")
		if _, ok := seen[check]; ok {
			return deck1, emptyDeck
		}
		seen[check] = true

		card1, card2 := deck1[0], deck2[0]
		deck1, deck2 = deck1[1:], deck2[1:]
		if len(deck1) >= card1 && len(deck2) >= card2 {
			nd1 := make([]int, card1)
			copy(nd1, deck1)
			nd2 := make([]int, card2)
			copy(nd2, deck2)
			sub1, _ := playRecursiveCombat(nd1, nd2)
			if len(sub1) > 0 {
				deck1 = append(deck1, card1, card2)
			} else {
				deck2 = append(deck2, card2, card1)
			}
		} else {
			if card1 > card2 {
				deck1 = append(deck1, card1, card2)
			} else {
				deck2 = append(deck2, card2, card1)
			}
		}

	}

	return deck1, deck2
}

func score(deck []int) int {
	s := 0
	for i := 0; i < len(deck); i++ {
		s += deck[i] * (len(deck) - i)
	}

	return s
}

func checksum(deck []int) string {
	return aoc.IntArrayAsString(deck, ",")
}

func readDeck(reader *bufio.Scanner) []int {
	deck := []int{}
	if !strings.HasPrefix(aoc.Read(reader), "Player") {
		panic("No player found")
	}
	for reader.Scan() {
		l := reader.Text()
		if len(l) == 0 {
			break
		}

		deck = append(deck, aoc.ToInt(l))
	}

	return deck
}
