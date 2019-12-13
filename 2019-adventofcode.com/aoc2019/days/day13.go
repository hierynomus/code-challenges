package days

import (
	"bufio"
	"fmt"
	"strconv"
	"sync"

	"github.com/hierynomus/aoc2019/aoc"
	"github.com/hierynomus/aoc2019/intcode"
)

type Day13 struct{}

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
				break
			}
		}
		a.Update(aoc.Point{X: x, Y: y}, k)
		if k == Paddle {
			a.PaddleUpdate <- a.Paddle
		} else if k == Ball {
			a.BallUpdate <- a.Ball
		}
	}
	a.wg.Done()
}

func (a *Arkanoid) read() (x, y, k int) {
	return <-a.Machine.IO.Output, <-a.Machine.IO.Output, <-a.Machine.IO.Output
}

func (a *Arkanoid) Update(p aoc.Point, k int) {
	if k == Ball {
		a.Ball = p
	} else if k == Paddle {
		a.Paddle = p
	} else if k == Block {
		a.Blocks[p] = struct{}{}
	} else if k == Empty {
		delete(a.Blocks, p)
	}
}

func (a *Arkanoid) Init() {
	x, y, k := 0, 0, 0
	for {
		x, y, k = a.read()
		fmt.Printf("%d: (%d, %d)\n", k, x, y)
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
	a.Machine.IO.Input <- joystick
}

const (
	Empty  int = 0
	Wall       = 1
	Block      = 2
	Paddle     = 3
	Ball       = 4
)

func (d *Day13) Solve(scanner *bufio.Scanner) (string, string) {
	if !scanner.Scan() {
		panic(fmt.Errorf("Boom!"))
	}

	program := aoc.AsIntArray(scanner.Text())
	icm := intcode.NewIntCodeMachine(program)

	go icm.Run()

	nrBlocks := 0

	for !icm.Closed {
		<-icm.IO.Output
		<-icm.IO.Output
		k := <-icm.IO.Output
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
	for len(game.Blocks) > 0 {
		if game.JoystickPosition != 0 {
			<-game.PaddleUpdate
		}
		<-game.BallUpdate
		game.PlayMove()
	}

	return strconv.Itoa(nrBlocks), strconv.Itoa(game.Score)
}
