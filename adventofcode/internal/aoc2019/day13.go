package aoc2019

import (
	"bufio"
	"fmt"
	"strconv"
	"sync"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
	"github.com/hierynomus/code-challenges/adventofcode/pkg/intcode"
)

type Arkanoid struct {
	Machine          *intcode.IntCodeMachine
	Ball             aoc.Point
	Blocks           map[aoc.Point]struct{}
	Paddle           aoc.Point
	Score            int
	BallUpdate       chan aoc.Point
	PaddleUpdate     chan aoc.Point
	wg               *sync.WaitGroup
	JoystickPosition int
}

func (a *Arkanoid) ReadAndUpdateState() {
	x, y, k := 0, 0, 0

	for {
		x, y, k = a.read()
		if x == -1 && y == 0 {
			a.Score = k
			if len(a.Blocks) == 0 {
				close(a.BallUpdate)
				break
			}
		}

		a.Update(aoc.Point{X: x, Y: y}, k)

		if k == Ball {
			a.BallUpdate <- a.Ball
		}
	}
	a.wg.Done()
}

func (a *Arkanoid) read() (x, y, k int) {
	return a.Machine.Output.Read(), a.Machine.Output.Read(), a.Machine.Output.Read()
}

func (a *Arkanoid) Update(p aoc.Point, k int) {
	switch k {
	case Ball:
		a.Ball = p
	case Paddle:
		a.Paddle = p
	case Block:
		a.Blocks[p] = struct{}{}
	case Empty:
		delete(a.Blocks, p)
	}
}

func (a *Arkanoid) Init() {
	x, y, k := 0, 0, 0

	for {
		x, y, k = a.read()
		if x == -1 && y == 0 {
			break
		}

		a.Update(aoc.Point{X: x, Y: y}, k)
	}

	a.Score = k
}

func (a *Arkanoid) PlayMove() {
	joystick := 0
	if a.Paddle.X > a.Ball.X {
		joystick = -1
	} else if a.Paddle.X < a.Ball.X {
		joystick = 1
	}

	a.JoystickPosition = joystick
	a.Machine.Input.Write(joystick)
}

const (
	Empty  int = 0
	Wall   int = 1
	Block  int = 2
	Paddle int = 3
	Ball   int = 4
)

func Day13(scanner *bufio.Scanner) (string, string) {
	if !scanner.Scan() {
		panic(fmt.Errorf("boom"))
	}

	program := aoc.AsIntArray(scanner.Text())
	icm := intcode.NewIntCodeMachine(program)

	go icm.Run()

	nrBlocks := 0

	for !icm.Closed {
		icm.Output.Read()
		icm.Output.Read()
		k := icm.Output.Read()

		if k == Block {
			nrBlocks++
		}
	}

	program[0] = 2
	arcade := intcode.NewIntCodeMachine(program)

	go arcade.Run()

	game := &Arkanoid{
		Machine:      arcade,
		Score:        0,
		Blocks:       map[aoc.Point]struct{}{},
		PaddleUpdate: make(chan aoc.Point),
		BallUpdate:   make(chan aoc.Point),
		wg:           &sync.WaitGroup{},
	}
	game.Init()
	game.wg.Add(1)

	go game.ReadAndUpdateState()

	game.PlayMove()

	for range game.BallUpdate {
		game.PlayMove()
	}

	return strconv.Itoa(nrBlocks), strconv.Itoa(game.Score)
}
