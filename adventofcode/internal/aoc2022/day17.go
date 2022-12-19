package aoc2022

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type Chamber struct {
	Wind    string
	WindIdx int
	Taken   map[aoc.Point]bool
	Height  int
}

type Block struct {
	Points []aoc.Point
}

var blocks = []Block{
	{
		Points: []aoc.Point{
			{X: 0, Y: 0},
			{X: 1, Y: 0},
			{X: 2, Y: 0},
			{X: 3, Y: 0},
		}, // -
	}, {
		Points: []aoc.Point{
			{X: 1, Y: 0},
			{X: 0, Y: 1},
			{X: 1, Y: 1},
			{X: 2, Y: 1},
			{X: 1, Y: 2},
		}, // +
	}, {
		Points: []aoc.Point{
			{X: 0, Y: 0},
			{X: 1, Y: 0},
			{X: 2, Y: 0},
			{X: 2, Y: 1},
			{X: 2, Y: 2},
		}, // L
	}, {
		Points: []aoc.Point{
			{X: 0, Y: 0},
			{X: 0, Y: 1},
			{X: 0, Y: 2},
			{X: 0, Y: 3},
		}, // |
	}, {
		Points: []aoc.Point{
			{X: 0, Y: 0},
			{X: 1, Y: 0},
			{X: 0, Y: 1},
			{X: 1, Y: 1},
		}, // square
	},
}

func (c *Chamber) DropBlock(block Block) {
	pos := aoc.Point{X: 2, Y: c.Height + 3}
	for {
		w := c.Wind[c.WindIdx]
		c.WindIdx = (c.WindIdx + 1) % len(c.Wind)

		switch {
		case w == '<' && c.CanMoveInto(block, pos.West()):
			pos = pos.West()
		case w == '>' && c.CanMoveInto(block, pos.East()):
			pos = pos.East()
		}

		if c.CanMoveInto(block, pos.South()) {
			pos = pos.South()
		} else {
			break
		}
	}

	for _, p := range block.Points {
		np := pos.Add(p)
		if np.Y+1 > c.Height {
			c.Height = np.Y + 1
		}
		c.Taken[np] = true
	}
}

func (c *Chamber) CanMoveInto(block Block, pos aoc.Point) bool {
	for _, p := range block.Points {
		m := pos.Add(p)
		if m.X < 0 || m.X > 6 || m.Y < 0 {
			return false
		}

		if c.Taken[m] {
			return false
		}
	}

	return true
}

func Day17(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	wind := aoc.Read(reader)
	c := &Chamber{
		Wind:    wind,
		WindIdx: 0,
		Height:  0,
		Taken:   make(map[aoc.Point]bool),
	}

	for r := 0; r < 2022; r++ {
		c.DropBlock(blocks[r%len(blocks)])
	}

	part1 = c.Height

	return strconv.Itoa(part1), strconv.Itoa(part2)
}
