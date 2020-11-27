package aoc2019

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/intcode"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day02(reader *bufio.Scanner) (string, string) {
	for reader.Scan() {
		program := aoc.AsIntArray(reader.Text())

		program[1] = 12
		program[2] = 2

		icm := intcode.NewIntCodeMachine(program)
		retCode := make(chan int)

		go func() {
			err := icm.Run()
			retCode <- err
		}()

		<-icm.ClosedCh

		code := <-retCode
		part1 := strconv.Itoa(code)

		for x := 0; x <= 99; x++ {
			for y := 0; y <= 99; y++ {
				icm.Reset()
				icm.Mem[1] = x
				icm.Mem[2] = y
				out := icm.Run()
				<-icm.ClosedCh

				if out == 19690720 {
					return part1, strconv.Itoa(100*x + y)
				}
			}
		}
	}
	panic(-1)
}
