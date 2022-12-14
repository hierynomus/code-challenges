package aoc2022

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

// const Tick struct{} = struct{}{}

// type HandheldDevice struct {
// 	Registers map[string]int
// 	Input chan Tick
// 	Output chan int
// 	Instructions []string
// 	InstructionPointer int
// 	CPUClock int
// }

// func NewHandheldDevice(instructions []string) *HandheldDevice {
// 	return &HandheldDevice{
// 		Registers: map[string]int{
// 			"X": 1,
// 		},
// 		InstructionPointer: 0,
// 		Instructions: instructions,
// 		CPUClock: 0,
// 	}
// }

// func (h *HandheldDevice) Run() {
// 	for {
// 		select {
// 		case <-h.Input:

// 	}
// 	for h.InstructionPointer < len(h.Instructions) {
// 		h.Tick()
// 	}
// }

type CRT struct {
	Display []string
}

func NewCRT() *CRT {
	return &CRT{
		Display: []string{},
	}
}

func (c *CRT) Tick(cycle, x int) {
	if cycle%40 == 0 {
		c.Display = append(c.Display, "")
	}

	if x-1 == cycle%40 || x+1 == cycle%40 || x == cycle%40 {
		c.Display[cycle/40] += "X"
	} else {
		c.Display[cycle/40] += "."
	}
}

func Day10(reader *bufio.Scanner) (string, string) {
	var part1 int

	output := []int{20, 60, 100, 140, 180, 220}
	instructions := aoc.ReadStringArray(reader)
	cycle := 0
	X := 1
	crt := NewCRT()
	for _, instruction := range instructions {
		i := strings.Split(instruction, " ")
		if i[0] == "noop" {
			crt.Tick(cycle, X)
			cycle++
			if aoc.IntArrayContains(output, cycle) {
				part1 += cycle * X
			}
		} else if i[0] == "addx" {
			crt.Tick(cycle, X)
			cycle++
			if aoc.IntArrayContains(output, cycle) {
				part1 += cycle * X
			}
			crt.Tick(cycle, X)
			cycle++
			if aoc.IntArrayContains(output, cycle) {
				part1 += cycle * X
			}
			X += aoc.ToInt(i[1])
		}
	}

	part2 := strings.Join(crt.Display, "\n")

	return strconv.Itoa(part1), part2
}
