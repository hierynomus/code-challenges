package days

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/aoc2019/aoc"
	"github.com/hierynomus/aoc2019/intcode"
)

type Day07 struct{}

func (d *Day07) Solve(scanner *bufio.Scanner) (string, string) {
	if scanner.Scan() {
		program := aoc.AsIntArray(scanner.Text())
		ampA := intcode.NewIntCodeMachine(program)
		ampB := intcode.NewIntCodeMachine(program)
		ampC := intcode.NewIntCodeMachine(program)
		ampD := intcode.NewIntCodeMachine(program)
		ampE := intcode.NewIntCodeMachine(program)

		permutations := Permutations([]int{0, 1, 2, 3, 4})

		maxOut := 0
		for _, perm := range permutations {
			ampA.IO.Input <- perm[0]
			ampB.IO.Input <- perm[1]
			ampC.IO.Input <- perm[2]
			ampD.IO.Input <- perm[3]
			ampE.IO.Input <- perm[4]
			go ampA.Run()
			go ampB.Run()
			go ampC.Run()
			go ampD.Run()
			go ampE.Run()
			ampA.IO.Input <- 0
			ampB.IO.Input <- <-ampA.IO.Output
			ampC.IO.Input <- <-ampB.IO.Output
			ampD.IO.Input <- <-ampC.IO.Output
			ampE.IO.Input <- <-ampD.IO.Output
			newMax := <-ampE.IO.Output
			if newMax > maxOut {
				maxOut = newMax
			}
			ampA.Reset()
			ampB.Reset()
			ampC.Reset()
			ampD.Reset()
			ampE.Reset()
		}

		permutations = Permutations([]int{5, 6, 7, 8, 9})

		maxOut2 := 0
		for _, perm := range permutations {
			ampA.Reset()
			ampB.Reset()
			ampC.Reset()
			ampD.Reset()
			ampE.Reset()

			ampA.IO.Input <- perm[0]
			ampB.IO.Input <- perm[1]
			ampC.IO.Input <- perm[2]
			ampD.IO.Input <- perm[3]
			ampE.IO.Input <- perm[4]
			go ampA.Run()
			go ampB.Run()
			go ampC.Run()
			go ampD.Run()
			go ampE.Run()
			newMax := 0
			for !ampA.Closed {
				ampA.IO.Input <- newMax
				ampB.IO.Input <- <-ampA.IO.Output
				ampC.IO.Input <- <-ampB.IO.Output
				ampD.IO.Input <- <-ampC.IO.Output
				ampE.IO.Input <- <-ampD.IO.Output
				newMax = <-ampE.IO.Output
			}
			if newMax > maxOut2 {
				maxOut2 = newMax
			}
		}

		return strconv.Itoa(maxOut), strconv.Itoa(maxOut2)
	}
	return "", ""
}

func Permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}
