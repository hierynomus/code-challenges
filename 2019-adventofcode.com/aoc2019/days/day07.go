package days

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/aoc2019/aoc"
	"github.com/hierynomus/aoc2019/intcode"
)

type AmplifierArray []*intcode.IntCodeMachine

func (a AmplifierArray) Run() {
	for _, amp := range a {
		go amp.Run()
	}
}

func (a AmplifierArray) Reset() {
	for _, amp := range a {
		amp.Reset()
	}
}

type Day07 struct{}

func (d *Day07) Solve(scanner *bufio.Scanner) (string, string) {
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
				amplifiers[i].IO.Input <- x
			}
			amplifiers.Run()
			amplifiers[0].IO.Input <- 0
			amplifiers[1].IO.Input <- <-amplifiers[0].IO.Output
			amplifiers[2].IO.Input <- <-amplifiers[1].IO.Output
			amplifiers[3].IO.Input <- <-amplifiers[2].IO.Output
			amplifiers[4].IO.Input <- <-amplifiers[3].IO.Output
			newMax := <-amplifiers[4].IO.Output
			if newMax > maxOut {
				maxOut = newMax
			}
			amplifiers.Reset()
		}

		permutations = aoc.Permutations([]int{5, 6, 7, 8, 9})

		maxOut2 := 0
		for _, perm := range permutations {
			amplifiers.Reset()

			for i, x := range perm {
				amplifiers[i].IO.Input <- x
			}
			amplifiers.Run()
			newMax := 0
			for !amplifiers[0].Closed {
				amplifiers[0].IO.Input <- newMax
				amplifiers[1].IO.Input <- <-amplifiers[0].IO.Output
				amplifiers[2].IO.Input <- <-amplifiers[1].IO.Output
				amplifiers[3].IO.Input <- <-amplifiers[2].IO.Output
				amplifiers[4].IO.Input <- <-amplifiers[3].IO.Output
				newMax = <-amplifiers[4].IO.Output
			}
			if newMax > maxOut2 {
				maxOut2 = newMax
			}
		}

		return strconv.Itoa(maxOut), strconv.Itoa(maxOut2)
	}
	return "", ""
}
