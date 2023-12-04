package aoc2022

import (
	"bufio"
	"math"
	"strconv"
)

func ParseSnafu(nr string) int {
	out := 0
	for i := len(nr) - 1; i >= 0; i-- {
		power := int(math.Pow(5, float64(i)))
		switch nr[len(nr)-i-1] {
		case '1':
			out += power
		case '2':
			out += 2 * power
		case '0':
			continue
		case '-':
			out -= power
		case '=':
			out -= 2 * power
		}
	}

	return out
}

func ToSnafu(nr int) string {
	out := ""
	for nr > 0 {
		nn := nr % 5
		nr /= 5
		switch nn {
		case 1:
			out = "1" + out
		case 2:
			out = "2" + out
		case 0:
			out = "0" + out
		case 4:
			out = "-" + out
			nr += 1
		case 3:
			out = "=" + out
			nr += 1
		}
	}

	return out
}

func Day25(reader *bufio.Scanner) (string, string) {
	var part2 int

	sum := 0
	for reader.Scan() {
		sum += ParseSnafu(reader.Text())
	}

	part1 := ToSnafu(sum)

	return part1, strconv.Itoa(part2)
}
