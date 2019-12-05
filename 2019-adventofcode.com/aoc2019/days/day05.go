package days

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/aoc2019/aoc"
	"github.com/hierynomus/aoc2019/intcode"
)

type Day05 struct{}

func (d *Day05) Solve(scanner *bufio.Scanner) (string, string) {
	if scanner.Scan() {
		program := aoc.AsIntArray(scanner.Text())
		icm := intcode.NewIntCodeMachine(program)
		var part1, part2 string
		go icm.Run()
		icm.IO.Input <- 1
		for i := range icm.IO.Output {
			if i != 0 {
				part1 = strconv.Itoa(i)
				break
			}
		}

		icm = intcode.NewIntCodeMachine(program)
		go icm.Run()
		icm.IO.Input <- 5
		part2 = strconv.Itoa(<-icm.IO.Output)
		return part1, part2
	}
	return "", ""
}
