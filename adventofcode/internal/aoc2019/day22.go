package aoc2019

import (
	"bufio"
	"fmt"
	"math/big"
	"strings"
)

// Linear represents a linear function of the kind a*x+b
type Linear struct {
	A, B *big.Int
}

func (l *Linear) Calculate(x *big.Int, p *big.Int) *big.Int {
	result := new(big.Int)
	result.Mul(l.A, x).Add(result, l.B).Mod(result, p)

	return result
}

func (l *Linear) Compose(a *big.Int, b *big.Int, p *big.Int) *Linear {
	newA := new(big.Int)
	newB := new(big.Int)

	newA.Mul(l.A, a).Mod(newA, p)              // (A * a) % p
	newB.Mul(l.B, a).Add(newB, b).Mod(newB, p) // (B * a + b) % p

	return &Linear{A: newA, B: newB}
}

type Deck []int

func Day22(scanner *bufio.Scanner) (string, string) {
	moves := []string{}
	for scanner.Scan() {
		moves = append(moves, scanner.Text())
	}

	NegOne := big.NewInt(-1)
	One := big.NewInt(1)
	Zero := new(big.Int)

	f := &Linear{A: One, B: Zero}
	size := big.NewInt(10007)

	for _, m := range moves {
		switch {
		case m == "deal into new stack":
			f = f.Compose(NegOne, NegOne, size)
			// funcs = append(funcs, &Linear{A: NegOne, B: NegOne})
		case strings.HasPrefix(m, "deal with increment"):
			i, ok := new(big.Int).SetString(m[20:], 10)
			if !ok {
				panic(fmt.Errorf("boom"))
			}

			f = f.Compose(i, Zero, size)
		case strings.HasPrefix(m, "cut"):
			i, ok := new(big.Int).SetString(m[4:], 10)
			if !ok {
				panic(fmt.Errorf("boom"))
			}

			i.Neg(i)

			f = f.Compose(One, i, size)
		}
	}

	card := big.NewInt(2019)

	card = f.Calculate(card, size)

	// card2019 := -1

	// for i, c := range spaceDeck {
	// 	if c == 2019 {
	// 		card2019 = i
	// 		break
	// 	}
	// }

	return card.String(), ""
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
