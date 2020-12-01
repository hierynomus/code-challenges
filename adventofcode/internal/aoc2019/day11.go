package aoc2019

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/internal/aoc2019/intcode"
	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type Robot struct {
	l aoc.Point
	d int
}

func (r *Robot) turn(turn int) {
	if turn == 0 {
		turn = -1
	}

	r.d = (4 + r.d + turn) % 4
}

func (r *Robot) move() {
	directions := map[int]aoc.Point{
		0: {X: 0, Y: -1},
		1: {X: 1, Y: 0},
		2: {X: 0, Y: 1},
		3: {X: -1, Y: 0},
	}
	dP := directions[r.d]
	r.l = aoc.Point{X: r.l.X + dP.X, Y: r.l.Y + dP.Y}
}

func Day11(scanner *bufio.Scanner) (string, string) {
	if !scanner.Scan() {
		panic(fmt.Errorf("boom"))
	}

	painted := map[aoc.Point]int{}
	program := aoc.AsIntArray(scanner.Text())
	icm := intcode.NewIntCodeMachine(program)

	// Part 1
	go icm.Run()

	robot := &Robot{l: aoc.Point{X: 0, Y: 0}, d: 0}

	for !icm.Closed {
		color, present := painted[robot.l]
		if !present {
			color = 0
		}
		icm.Input.Write(color)
		painted[robot.l] = icm.Output.Read()
		robot.turn(icm.Output.Read())
		robot.move()
	}
	<-icm.ClosedCh
	// Part 2
	icm.Reset()

	grid := make([][]int, 6)

	for y := 0; y < 6; y++ {
		grid[y] = make([]int, 43)
	}

	go icm.Run()

	robot = &Robot{l: aoc.Point{X: 0, Y: 0}, d: 0}
	grid[robot.l.Y][robot.l.X] = 1

	for !icm.Closed {
		color := grid[robot.l.Y][robot.l.X]
		icm.Input.Write(color)
		grid[robot.l.Y][robot.l.X] = icm.Output.Read()
		robot.turn(icm.Output.Read())
		robot.move()
	}

	return strconv.Itoa(len(painted)), PrintGrid(grid)
}

func PrintGrid(grid [][]int) string {
	out := "\n"

	for y := 0; y < len(grid); y++ {
		l := grid[y]
		for x := 0; x < len(l); x++ {
			if l[x] == 1 {
				out += "#"
			} else {
				out += "."
			}
		}

		out += "\n"
	}

	return out
}
