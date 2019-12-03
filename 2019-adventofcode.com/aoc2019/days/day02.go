package days

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/aoc2019/aoc"
	"github.com/hierynomus/aoc2019/intcode"
)

type Day02 struct{}

func (d *Day02) Solve(reader *bufio.Scanner) (string, string) {
	for reader.Scan() {
		program, err := aoc.AsIntArray(reader.Text())
		if err != nil {
			panic(err)
		}

		program[1] = 12
		program[2] = 2

		icm := intcode.NewIntCodeMachine(program)
		part1 := strconv.Itoa(icm.Run())
		for x := 0; x <= 99; x++ {
			for y := 0; y <= 99; y++ {
				icm.SetNounVerb(x, y)
				out := icm.Run()
				if out == 19690720 {
					return part1, strconv.Itoa(100*x + y)
				}
			}
		}
	}
	panic(-1)
}