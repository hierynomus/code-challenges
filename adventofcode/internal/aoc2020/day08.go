package aoc2020

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type Instruction struct {
	Name   string
	Offset int
}

func Day08(reader *bufio.Scanner) (string, string) {
	lines := aoc.ReadStringArray(reader)
	program := []*Instruction{}
	for _, l := range lines {
		s := strings.Split(l, " ")
		program = append(program, &Instruction{Name: s[0], Offset: aoc.ToInt(s[1])})
	}

	part1, _ := runBootloader(program, 0, 0, aoc.IntSet{}, true)
	part2, success := runBootloader(program, 0, 0, aoc.IntSet{}, false)
	if !success {
		panic("Did not succeed")
	}
	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func runBootloader(program []*Instruction, ip int, acc int, seen aoc.IntSet, modified bool) (int, bool) {
	for ip < len(program) {
		if seen.Contains(ip) {
			return acc, false
		}
		seen.Add(ip)
		i := program[ip]
		switch i.Name {
		case "acc":
			acc += i.Offset
		case "nop":
			if !modified {
				macc, success := runBootloader(program, ip+i.Offset, acc, seen.Copy(), true)
				if success {
					return macc, true
				}
			}
		case "jmp":
			if !modified {
				macc, success := runBootloader(program, ip+1, acc, seen.Copy(), true)
				if success {
					return macc, true
				}
			}
			ip += i.Offset
			continue
		}
		ip += 1
	}

	return acc, true
}
