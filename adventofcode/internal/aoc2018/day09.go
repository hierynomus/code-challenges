package aoc2018

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type Marble struct {
	Value int
	Left  *Marble
	Right *Marble
}

func NewMarble(v int) *Marble {
	m := &Marble{
		Value: v,
	}
	m.Left = m
	m.Right = m
	return m
}

func (m *Marble) Insert(o *Marble) {
	right := m.Right
	o.Left, o.Right = m, right
	right.Left = o
	m.Right = o
}

func (m *Marble) Remove() int {
	left, right := m.Left, m.Right
	left.Right = right
	right.Left = left
	return m.Value
}

func Day09(reader *bufio.Scanner) (string, string) {
	if !reader.Scan() {
		panic("no input")
	}

	l := strings.Split(reader.Text(), " ")
	players, lastMarble := aoc.ToInt(l[0]), aoc.ToInt(l[6])

	return strconv.Itoa(playMarbleGame(players, lastMarble)), strconv.Itoa(playMarbleGame(players, lastMarble*100))
}

func playMarbleGame(players, lastMarble int) int {
	currentMarble := NewMarble(0)
	scores := make([]int, players)
	player := 0

	for mv := 1; mv <= lastMarble; mv++ {
		if mv%23 == 0 {
			scores[player] += mv
			currentMarble = currentMarble.Left.Left.Left.Left.Left.Left
			score := currentMarble.Remove()
			scores[player] += score
		} else {
			currentMarble = currentMarble.Right.Right
			m := NewMarble(mv)
			currentMarble.Insert(m)
		}
		player = (player + 1) % players
	}

	return aoc.Max(scores)
}
