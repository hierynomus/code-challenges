package aoc2019

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/internal/aoc2019/intcode"
	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day05(scanner *bufio.Scanner) (string, string) {
	if scanner.Scan() {
		program := aoc.AsIntArray(scanner.Text())
		icm := intcode.NewIntCodeMachine(program)

		var part1, part2 string

		go icm.Run()

		icm.Input.Write(1)
	loop:
		for {
			select {
			case <-icm.ClosedCh:
				break loop
			default:
				i := icm.Output.Read()
				if i != 0 {
					part1 = strconv.Itoa(i)
					break
				}
			}
		}

		icm = intcode.NewIntCodeMachine(program)
		go icm.Run()
		icm.Input.Write(5)
		part2 = strconv.Itoa(icm.Output.Read())

		return part1, part2
	}

	return "", ""
}
