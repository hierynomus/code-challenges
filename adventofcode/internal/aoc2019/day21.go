package aoc2019

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
	"github.com/hierynomus/code-challenges/adventofcode/pkg/intcode"
)

func Day21(scanner *bufio.Scanner) (string, string) {
	if !scanner.Scan() {
		panic(fmt.Errorf("boom"))
	}

	program := aoc.AsIntArray(scanner.Text())
	icm := intcode.NewIntCodeMachine(program)

	go icm.Run()

	readPrompt(icm)

	go func() {
		instructions := "NOT C T\nOR D J\nAND T J\nNOT A T\nOR T J\nWALK\n"
		for _, c := range instructions {
			icm.Input.Write(int(c))
		}
	}()

	Damage := readOut(icm)
	<-icm.ClosedCh

	icm.Reset()

	go icm.Run()

	readPrompt(icm)

	go func() {
		// NOT B T\nNOT C J\nOR T J\n -> If B or C has gap, T is true
		// AND D J\nNOT E T\nAND T J\n -> If both D and E have floor, don't jump yet
		// NOT A T\nOR T J -> Ensure we jump if we're directly in front of gap
		instructions := "NOT B T\nNOT C J\nOR T J\nAND D J\nOR H T\nAND T J\nNOT A T\nOR T J\nRUN\n"
		// instructions := "OR D T\nOR H J\nAND T J\nRUN"
		for _, c := range instructions {
			icm.Input.Write(int(c))
		}
	}()

	Damage2 := readOut(icm)

	return strconv.Itoa(Damage), strconv.Itoa(Damage2)
}

func readPrompt(icm *intcode.IntCodeMachine) {
	for {
		select {
		case <-icm.ClosedCh:
			return
		default:
			i := icm.Output.Read()

			switch {
			case i < 255 && i != 10:
				// fmt.Printf("%c", rune(i))
			default:
				return
			}
		}
	}
}

func readOut(icm *intcode.IntCodeMachine) int {
	for {
		select {
		case <-icm.ClosedCh:
			return -1
		default:
			i := icm.Output.Read()

			switch {
			case i < 255:
				// fmt.Printf("%c", rune(i))
			default:
				return i
			}
		}
	}
}
