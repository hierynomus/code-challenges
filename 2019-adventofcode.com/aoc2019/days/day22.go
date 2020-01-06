package days

import (
	"bufio"
	"strconv"
	"strings"
)

type Day22 struct{}

type Deck []int

func (d *Day22) Solve(scanner *bufio.Scanner) (string, string) {
	moves := []string{}
	for scanner.Scan() {
		moves = append(moves, scanner.Text())
	}

	spaceDeck := make(Deck, 10007)
	for i := 0; i < 10007; i++ {
		spaceDeck[i] = i
	}

	for _, m := range moves {
		switch {
		case m == "deal into new stack":
			spaceDeck.reverse()
		case strings.HasPrefix(m, "deal with increment"):
			i, err := strconv.Atoi(m[20:])
			if err != nil {
				panic(err)
			}

			spaceDeck = spaceDeck.deal(i)
		case strings.HasPrefix(m, "cut"):
			i, err := strconv.Atoi(m[4:])
			if err != nil {
				panic(err)
			}

			spaceDeck = spaceDeck.cut(i)
		}
	}

	card2019 := -1

	for i, c := range spaceDeck {
		if c == 2019 {
			card2019 = i
			break
		}
	}

	return strconv.Itoa(card2019), ""
}

func (d Deck) deal(i int) Deck {
	newDeck := make(Deck, 10007)
	pos := 0

	for _, c := range d {
		newDeck[pos] = c
		pos += i
		pos %= len(d)
	}

	return newDeck
}

func (d Deck) reverse() {
	for i := len(d)/2 - 1; i >= 0; i-- {
		opp := len(d) - 1 - i
		d[i], d[opp] = d[opp], d[i]
	}
}

func (d Deck) cut(n int) Deck {
	if n > 0 {
		return append(d[n:], d[0:n]...)
	}

	c := len(d) + n

	return append(d[c:], d[0:c]...)
}
