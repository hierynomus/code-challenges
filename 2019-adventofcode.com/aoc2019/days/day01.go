package days

import (
	"bufio"

	"strconv"
)

type Day01 struct{}

func FuelFor(mass int) int {
	return mass/3 - 2
}

func (d *Day01) Solve(reader *bufio.Scanner) (string, string) {
	var input []int

	for reader.Scan() {
		i, err := strconv.Atoi(reader.Text())
		if err != nil {
			panic(err)
		}

		input = append(input, i)
	}

	part1 := 0
	for _, i := range input {
		part1 += FuelFor(i)
	}

	part2 := 0

	for _, i := range input {
		fuel := FuelFor(i)
		for fuel > 0 {
			part2 += fuel
			fuel = FuelFor(fuel)
		}
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}
