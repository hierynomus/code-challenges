package aoc2022

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func StartOfPacket(arr []rune) int {
	for i := 0; i < len(arr)-3; i++ {
		a := arr[i : i+4]
		s := aoc.NewRuneSet(a)
		if len(s) == 4 {
			return i + 4
		}
	}

	return -1
}

func StartOfMessage(arr []rune) int {
	for i := 0; i < len(arr)-14; i++ {
		a := arr[i : i+14]
		s := aoc.NewRuneSet(a)
		if len(s) == 14 {
			return i + 14
		}
	}

	return -1
}

func Day06(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	data := []rune(aoc.Read(reader))
	part1 = StartOfPacket(data)
	part2 = StartOfMessage(data)
	return strconv.Itoa(part1), strconv.Itoa(part2)
}
