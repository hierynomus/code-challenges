package days

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/hierynomus/aoc2019/aoc"
	"github.com/hierynomus/aoc2019/intcode"
)

type Day09 struct{}

func (d *Day09) Solve(scanner *bufio.Scanner) (string, string) {
	if !scanner.Scan() {
		panic(fmt.Errorf("Boom!"))
	}
	program := aoc.AsIntArray(scanner.Text())
	icm := intcode.NewIntCodeMachine(program)

	go icm.Run()
	icm.IO.Input <- 1
	part1 := <-icm.IO.Output

	icm.Reset()
	go icm.Run()
	icm.IO.Input <- 2
	part2 := <-icm.IO.Output

	return strconv.Itoa(part1), strconv.Itoa(part2)
}
