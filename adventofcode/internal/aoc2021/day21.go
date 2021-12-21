package aoc2021

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type DeterministicDie struct {
	counter int
	sides   int
	rolls   int
}

type DiracDicePlayer struct {
	position int
	score    int
}

func NewPlayer(position int) *DiracDicePlayer {
	return &DiracDicePlayer{position: position, score: 0}
}

func (p *DiracDicePlayer) Move(d *DeterministicDie) {
	newPos := p.position + d.Throw3()
	for newPos > 10 {
		newPos -= 10
	}
	p.position = newPos
	p.score += p.position
}

func NewDeterminisiticDie(sides int) *DeterministicDie {
	return &DeterministicDie{counter: 1, sides: sides}
}

func (d *DeterministicDie) Throw3() int {
	r := 0
	for i := 0; i < 3; i++ {
		r += d.counter
		d.counter = (d.counter + 1) % d.sides
	}

	d.rolls += 3

	return r
}

func Day21(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	var pos1, pos2 int
	_, err := fmt.Sscanf(aoc.Read(reader), "Player 1 starting position: %d", &pos1)
	if err != nil {
		panic(err)
	}

	_, err = fmt.Sscanf(aoc.Read(reader), "Player 2 starting position: %d", &pos2)
	if err != nil {
		panic(err)
	}

	player1, player2 := NewPlayer(pos1), NewPlayer(pos2)
	die := NewDeterminisiticDie(100)
	for {
		player1.Move(die)
		if player1.score >= 1000 {
			part1 = player2.score * die.rolls
			break
		} else {
			player2.Move(die)
			if player2.score >= 1000 {
				part1 = player1.score * die.rolls
				break
			}
		}
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}
