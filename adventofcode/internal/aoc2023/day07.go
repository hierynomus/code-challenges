package aoc2023

import (
	"bufio"
	"sort"
	"strconv"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

var CamelCards = aoc.AsIndexMapping([]rune{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'})
var CamelCardsJoker = aoc.AsIndexMapping([]rune{'J', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'Q', 'K', 'A'})

type CamelCardHandType int

const (
	HighCard CamelCardHandType = iota
	OnePair
	TwoPairs
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type CamelCardHand struct {
	Hand  string
	Ranks map[rune]int
	Type  CamelCardHandType
	Bid   int
}

func (c CamelCardHand) Rank(cardNr int) int {
	return c.Ranks[rune(c.Hand[cardNr])]
}

type CamelCardSlice []CamelCardHand

func (c CamelCardSlice) Len() int {
	return len(c)
}

func (c CamelCardSlice) Less(i, j int) bool {
	t, o := c[i], c[j]
	if t.Type != o.Type {
		return t.Type < o.Type
	}

	for k := 0; k < 5; k++ {
		if t.Rank(k) != o.Rank(k) {
			return t.Rank(k) < o.Rank(k)
		}
	}

	return false
}

func (c CamelCardSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c CamelCardSlice) Score() int {
	score := 0
	for i := 0; i < len(c); i++ {
		score += c[i].Bid * (i + 1)
	}

	return score
}

// func parseHandType(hand string) CamelCardHandType {
// 	cardHisto := aoc.MakeRuneHistogram([]rune(hand))

// 	_, maxCount := cardHisto.MaxC()

// 	switch len(cardHisto) {
// 	case 1:
// 		return FiveOfAKind
// 	case 2:
// 		if maxCount == 4 {
// 			return FourOfAKind
// 		} else {
// 			return FullHouse
// 		}
// 	case 3:
// 		if maxCount == 3 {
// 			return ThreeOfAKind
// 		} else {
// 			return TwoPairs
// 		}
// 	case 4:
// 		return OnePair
// 	default:
// 		return HighCard
// 	}
// }

func parseHandType(hand string, joker bool) CamelCardHandType {
	cardHisto := aoc.MakeRuneHistogram([]rune(hand))

	jokers := int64(0)
	if joker {
		jokers = cardHisto['J']
		delete(cardHisto, 'J')
	}

	_, maxCount := cardHisto.MaxC()

	var t CamelCardHandType
	switch len(cardHisto) {
	case 0:
		// Corner case, 5 jokers
		t = FiveOfAKind
	case 1: // Could be 5 cards, or X cards + jokers
		t = FiveOfAKind
	case 2:
		if maxCount+jokers == 4 {
			t = FourOfAKind
		} else {
			t = FullHouse
		}
	case 3:
		if maxCount+jokers == 3 {
			t = ThreeOfAKind
		} else {
			t = TwoPairs
		}
	case 4:
		if maxCount+jokers == 2 {
			t = OnePair
		} else {
			t = HighCard
		}
	default:
		t = HighCard
	}

	return t
}

func parseCamelCardHand(line string, joker bool) CamelCardHand {
	f := strings.Fields(line)
	ranks := CamelCards
	if joker {
		ranks = CamelCardsJoker
	}

	hand := CamelCardHand{
		Hand:  f[0],
		Ranks: ranks,
		Bid:   aoc.ToInt(f[1]),
		Type:  parseHandType(f[0], joker),
	}

	return hand
}

func Day07(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	lines := aoc.ReadStringArray(reader)
	hands := make([]CamelCardHand, len(lines))
	jokerHands := make([]CamelCardHand, len(lines))
	for i, line := range lines {
		hands[i] = parseCamelCardHand(line, false)
		jokerHands[i] = parseCamelCardHand(line, true)
	}

	s := CamelCardSlice(hands)
	sort.Sort(s)
	part1 = s.Score()

	s = CamelCardSlice(jokerHands)
	sort.Sort(s)
	part2 = s.Score()

	return strconv.Itoa(part1), strconv.Itoa(part2)
}
