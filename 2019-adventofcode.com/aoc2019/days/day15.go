package days

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/hierynomus/aoc2019/aoc"
	"github.com/hierynomus/aoc2019/intcode"
)

var emptyStruct struct{} = struct{}{} //nolint:gochecknoglobals

type Location struct {
	Pos   aoc.Point
	From  *Location
	Exits map[int]*Location
	Walls map[int]struct{}
}

func (l *Location) AddWall(d int) {
	l.Walls[d] = emptyStruct
}

func (l *Location) AddExit(d int, e *Location) {
	l.Exits[d] = e
	e.From = l

	switch d {
	case 1:
		e.Exits[2] = l
	case 2:
		e.Exits[1] = l
	case 3:
		e.Exits[4] = l
	case 4:
		e.Exits[3] = l
	}
}

func (l *Location) UnexploredDirections() []int {
	directions := []int{}

	for i := 1; i < 5; i++ {
		if _, ok := l.Exits[i]; ok {
			continue
		}

		if _, ok := l.Walls[i]; ok {
			continue
		}

		directions = append(directions, i)
	}

	return directions
}

func (l *Location) Backtrack() int {
	for d, e := range l.Exits {
		if e == l.From {
			return d
		}
	}

	panic(fmt.Errorf("could not backtrack to previous point"))
}

type Day15 struct{}

func (d *Day15) Solve(scanner *bufio.Scanner) (string, string) {
	maze := map[aoc.Point]*Location{}
	moves := map[int]aoc.Point{
		1: aoc.Point{X: 0, Y: -1}, //nolint:gofmt
		2: aoc.Point{X: 0, Y: 1},
		3: aoc.Point{X: 1, Y: 0},
		4: aoc.Point{X: -1, Y: 0},
	}

	if !scanner.Scan() {
		panic(fmt.Errorf("boom"))
	}

	program := aoc.AsIntArray(scanner.Text())
	icm := intcode.NewIntCodeMachine(program)

	rootPos := NewLocation(maze, aoc.Point{X: 0, Y: 0})
	rootPos.From = rootPos

	go icm.Run()

	curPos := rootPos

loop:
	for {
		select {
		case <-icm.ClosedCh:
			break loop
		default:
			unexplored := curPos.UnexploredDirections()
			if len(unexplored) == 0 { // Fully explored location, backtrack a step
				d := curPos.Backtrack()
				curPos = curPos.From

				icm.IO.Input <- d
				<-icm.IO.Output // Always a room...
			} else {
				d := unexplored[0]
				newPos := curPos.Pos.Add(moves[d])

				if l, ok := maze[newPos]; ok {
					curPos.Exits[d] = l
					continue
				}

				icm.IO.Input <- d
				o := <-icm.IO.Output

				switch o {
				case 0: // Wall
					curPos.Walls[d] = emptyStruct
				case 1: // Room
					nl := NewLocation(maze, newPos)
					curPos.AddExit(d, nl)
					curPos = nl
				case 2:
					nl := NewLocation(maze, newPos)
					curPos.AddExit(d, nl)
					curPos = nl
					break loop
				}
			}
		}
	}

	steps := 0

	for curPos != rootPos {
		steps++

		curPos = curPos.From
	}

	return strconv.Itoa(steps), ""
}

func NewLocation(maze map[aoc.Point]*Location, p aoc.Point) *Location {
	l := &Location{
		Pos:   p,
		From:  nil,
		Exits: map[int]*Location{},
		Walls: map[int]struct{}{},
	}
	maze[p] = l

	return l
}
