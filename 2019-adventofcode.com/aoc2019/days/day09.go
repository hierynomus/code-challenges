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
		panic(fmt.Errorf("boom"))
	}

	program := aoc.AsIntArray(scanner.Text())
	icm := intcode.NewIntCodeMachine(program)

	go icm.Run()
	icm.Input.Write(1)
	part1 := icm.Output.Read()

	icm.Reset()

	go icm.Run()
	icm.Input.Write(2)
	part2 := icm.Output.Read()

	return strconv.Itoa(part1), strconv.Itoa(part2)
}
