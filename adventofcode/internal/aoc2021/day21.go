package aoc2021

import (
	"bufio"
	"fmt"

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
	var part1, part2 int64

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
			part1 = int64(player2.score * die.rolls)
			break
		} else {
			player2.Move(die)
			if player2.score >= 1000 {
				part1 = int64(player1.score * die.rolls)
				break
			}
		}
	}

	wins1, wins2 := PlayDiracDice(pos1, pos2)

	if wins1 > wins2 {
		part2 = wins1
	} else {
		part2 = wins2
	}
	return aoc.Int64ToString(part1), aoc.Int64ToString(part2)
}

type DiracDiceGame struct {
	pos1, pos2     int
	score1, score2 int
}

var ThreeRollDirac = map[int]int{
	3: 1,
	4: 3,
	5: 6,
	6: 7,
	7: 6,
	8: 3,
	9: 1,
}

func PlayDiracDice(pos1, pos2 int) (int64, int64) {
	games := map[DiracDiceGame]int{}
	games[DiracDiceGame{pos1, pos2, 0, 0}] = 1
	won1, won2 := int64(0), int64(0)
	for len(games) > 0 {
		for k, v := range games {
			if v == 0 {
				delete(games, k)
				continue
			}

			games[k] = 0

			for roll, nr := range ThreeRollDirac {
				pos := k.pos1 + roll
				for pos > 10 {
					pos -= 10
				}
				score1 := k.score1 + pos
				if score1 >= 21 {
					won1 += int64(nr * v)
					continue
				}

				for roll2, nr2 := range ThreeRollDirac {
					pos2 := k.pos2 + roll2
					for pos2 > 10 {
						pos2 -= 10
					}
					score2 := k.score2 + pos2
					if score2 >= 21 {
						won2 += int64(nr2 * v)
						continue
					}
					g := DiracDiceGame{pos, pos2, score1, score2}
					games[g] += nr * nr2 * v
				}
			}
		}
	}

	return won1, won2

}
