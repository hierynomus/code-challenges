package aoc2019

import (
	"bufio"
	"fmt"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
	"github.com/hierynomus/code-challenges/adventofcode/pkg/intcode"
)

func Day09(scanner *bufio.Scanner) (string, string) {
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
