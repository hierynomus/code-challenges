package aoc2019

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/internal/aoc2019/intcode"
	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type AmplifierArray []*intcode.IntCodeMachine

func (a AmplifierArray) Run() {
	for _, amp := range a {
		go amp.Run()
	}
}

func (a AmplifierArray) Reset() {
	for _, amp := range a {
		<-amp.ClosedCh
		amp.Reset()
	}
}

func (a AmplifierArray) PassValues(in int) int {
	for i := 0; i < 5; i++ {
		a[i].Input.Write(in)
		in = a[i].Output.Read()
	}

	return in
}

func Day07(scanner *bufio.Scanner) (string, string) {
	amplifiers := AmplifierArray{}

	if scanner.Scan() {
		program := aoc.AsIntArray(scanner.Text())

		for i := 0; i < 5; i++ {
			amplifiers = append(amplifiers, intcode.NewIntCodeMachine(program))
		}

		permutations := aoc.Permutations([]int{0, 1, 2, 3, 4})

		maxOut := 0

		for _, perm := range permutations {
			for i, x := range perm {
				amplifiers[i].Input.Write(x)
			}

			amplifiers.Run()

			newMax := amplifiers.PassValues(0)

			if newMax > maxOut {
				maxOut = newMax
			}

			amplifiers.Reset()
		}

		permutations = aoc.Permutations([]int{5, 6, 7, 8, 9})

		maxOut2 := 0
		for _, perm := range permutations {
			for i, x := range perm {
				amplifiers[i].Input.Write(x)
			}

			amplifiers.Run()

			newMax := 0

			for !amplifiers[0].Closed {
				newMax = amplifiers.PassValues(newMax)
			}
			amplifiers.Reset()

			if newMax > maxOut2 {
				maxOut2 = newMax
			}
		}

		return strconv.Itoa(maxOut), strconv.Itoa(maxOut2)
	}

	return "", ""
}
