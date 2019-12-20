package days

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/hierynomus/aoc2019/aoc"
	"github.com/hierynomus/aoc2019/intcode"
)

var emptyStruct struct{} = struct{}{} //nolint:gochecknoglobals

type RepairDroid struct {
	brain    *intcode.IntCodeMachine
	Memory   map[aoc.Point]*Location
	position *Location
	Visited  chan *Location
	Done     chan struct{}
}

func NewRepairDroid(icm *intcode.IntCodeMachine) *RepairDroid {
	b := map[aoc.Point]*Location{}
	l := NewLocation(b, aoc.Point{X: 0, Y: 0})
	l.From = l // Start location

	return &RepairDroid{
		brain:    icm,
		Memory:   b,
		position: l,
		Visited:  make(chan *Location),
		Done:     make(chan struct{}, 1),
	}
}

func (d *RepairDroid) Explore() {
	go d.brain.Run()

	moves := map[int]aoc.Point{
		1: aoc.Point{X: 0, Y: -1}, //nolint:gofmt
		2: aoc.Point{X: 0, Y: 1},
		3: aoc.Point{X: 1, Y: 0},
		4: aoc.Point{X: -1, Y: 0},
	}

loop:
	for {
		select {
		case <-d.brain.ClosedCh:
			d.Done <- emptyStruct
			break loop
		default:
			unexplored := d.position.UnexploredDirections()
			if len(unexplored) == 0 && d.position.StartLocation() {
				// Fully explored current position and back at start!
				d.Done <- emptyStruct // Signal done!
				break loop
			}
			if len(unexplored) == 0 { // Fully explored location, backtrack a step
				dir := d.position.Backtrack()
				d.position = d.position.From

				d.brain.IO.Input <- dir
				<-d.brain.IO.Output // Always a room...
			} else {
				dir := unexplored[0]
				newPos := d.position.Pos.Add(moves[dir])

				if l, ok := d.Memory[newPos]; ok {
					d.position.Exits[dir] = l
					continue
				}

				d.brain.IO.Input <- dir
				o := <-d.brain.IO.Output

				if o == 0 { // Wall
					d.position.Walls[dir] = emptyStruct
				} else { // Room
					nl := NewLocation(d.Memory, newPos)
					nl.HasOxygen = o == 2
					d.position.AddExit(dir, nl)
					d.position = nl
					d.Visited <- nl
				}
			}
		}
	}
}

type Location struct {
	Pos       aoc.Point
	From      *Location
	Exits     map[int]*Location
	Walls     map[int]struct{}
	HasOxygen bool
}

func (l *Location) StartLocation() bool {
	return l.From == l
}

func (l *Location) Neighbours() []*Location {
	n := []*Location{}
	for _, v := range l.Exits {
		n = append(n, v)
	}

	return n
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
	if !scanner.Scan() {
		panic(fmt.Errorf("boom"))
	}

	program := aoc.AsIntArray(scanner.Text())
	icm := intcode.NewIntCodeMachine(program)

	droid := NewRepairDroid(icm)
	go droid.Explore()

	var oxygenTank *Location

visitLoop:
	for {
		select {
		case <-droid.Done:
			break visitLoop
		case r := <-droid.Visited:
			if r.HasOxygen {
				oxygenTank = r
			}
		}
	}

	stepsToStart := 0
	curPos := oxygenTank

	for !curPos.StartLocation() {
		stepsToStart++

		curPos = curPos.From
	}

	roomsToFill := oxygenTank.Neighbours()
	time := 0

	for len(roomsToFill) > 0 {
		time++

		nextTime := []*Location{}

		for _, r := range roomsToFill {
			if r.HasOxygen {
				continue
			}

			r.HasOxygen = true
			n := r.Neighbours()

			for _, x := range n {
				if x.HasOxygen {
					continue
				}

				nextTime = append(nextTime, x)
			}
		}

		roomsToFill = nextTime
	}

	return strconv.Itoa(stepsToStart), strconv.Itoa(time)
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
