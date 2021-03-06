package aoc2019

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/internal/aoc2019/intcode"
	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type TractorBeam struct {
	icm *intcode.IntCodeMachine
}

func (tb *TractorBeam) Pulled(x, y int) bool {
	go tb.icm.Run()

	tb.icm.Input.Write(x)
	tb.icm.Input.Write(y)

	out := tb.icm.Output.Read()
	<-tb.icm.ClosedCh
	tb.icm.Reset()

	return out > 0
}

func Day19(scanner *bufio.Scanner) (string, string) {
	if !scanner.Scan() {
		panic(fmt.Errorf("boom"))
	}

	program := aoc.AsIntArray(scanner.Text())
	icm := intcode.NewIntCodeMachine(program)
	tb := &TractorBeam{icm}

	nr := 0

	for y := 0; y < 50; y++ {
		for x := 0; x < 50; x++ {
			if tb.Pulled(x, y) {
				nr++
			}
		}
	}

	xCount := 1
	yCount := 1
	pos := aoc.Point{X: 800, Y: 800} // Works from any set of coords on left side of TB

	for !(xCount >= 100 && yCount >= 100) {
		if tb.Pulled(pos.X+xCount, pos.Y) {
			xCount++
		} else if xCount < 100 {
			xCount = 1
			if yCount > 1 {
				yCount--
			}
			pos = pos.AddXY(0, 1)
			for !tb.Pulled(pos.X, pos.Y) {
				pos = pos.AddXY(1, 0)
			}
		}

		if tb.Pulled(pos.X, pos.Y+yCount) {
			yCount++
		} else {
			yCount = 1
			if xCount > 1 {
				xCount--
			}
			pos = pos.AddXY(1, 0)
		}
	}

	return strconv.Itoa(nr), strconv.Itoa(pos.X*10000 + pos.Y)
}
