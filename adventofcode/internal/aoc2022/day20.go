package aoc2022

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type GroveCoord struct {
	Next  *GroveCoord
	Prev  *GroveCoord
	Val   int
	Moved bool
}

func (c *GroveCoord) LinkNext(next *GroveCoord) {
	c.Next = next
	next.Prev = c
}

func (c *GroveCoord) Ith(i int) *GroveCoord {
	cc := c
	for i > 0 {
		cc = cc.Next
		i--
	}

	return cc
}

func Mix(inp []*GroveCoord) {
	for i := 0; i < len(inp); i++ {
		c := inp[i]
		p := c
		c.Prev.LinkNext(c.Next) // detach c
		if c.Val > 0 {
			for i := 0; i < c.Val%(len(inp)-1); i++ {
				p = p.Next
			}
		} else {
			for i := 0; i < (-c.Val % (len(inp) - 1)); i++ {
				p = p.Prev
			}
			p = p.Prev
		}
		c.LinkNext(p.Next) // c -> p.Next
		p.LinkNext(c)      // p -> c
	}
}

func Reset(inp []*GroveCoord) {
	for i := 0; i < len(inp)-1; i++ {
		inp[i].LinkNext(inp[i+1])
	}
	inp[len(inp)-1].LinkNext(inp[0])
}

func Day20(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	inp := []*GroveCoord{}
	inp = append(inp, &GroveCoord{Val: aoc.ToInt(aoc.Read(reader))})

	i := 0
	zero := 0
	for reader.Scan() {
		line := aoc.ToInt(reader.Text())
		c := &GroveCoord{Val: line}
		inp[i].LinkNext(c)
		i++
		inp = append(inp, c)
		if line == 0 {
			zero = i
		}
	}

	inp[i].LinkNext(inp[0])

	Mix(inp)
	x := inp[zero].Ith(1000)
	y := x.Ith(1000)
	z := y.Ith(1000)
	part1 = x.Val + y.Val + z.Val

	Reset(inp)
	for i := 0; i < len(inp); i++ {
		inp[i].Val *= 811589153
	}

	for i := 0; i < 10; i++ {
		Mix(inp)
	}

	x2 := inp[zero].Ith(1000)
	y2 := x2.Ith(1000)
	z2 := y2.Ith(1000)
	part2 = x2.Val + y2.Val + z2.Val

	return strconv.Itoa(part1), strconv.Itoa(part2)
}
