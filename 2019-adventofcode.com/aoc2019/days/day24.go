package days

import (
	"bufio"
	"fmt"
	"math"

	"github.com/hierynomus/aoc2019/aoc"
)

type Day24 struct{}

func (d *Day24) Solve(scanner *bufio.Scanner) (string, string) {
	initialEris := [][]rune{}
	initialEris = append(initialEris, []rune("......."))

	for scanner.Scan() {
		initialEris = append(initialEris, []rune(fmt.Sprintf(".%s.", scanner.Text())))
	}

	initialEris = append(initialEris, []rune("......."))

	eris := make([][]rune, 7)
	copy(eris, initialEris)

	bioSeen := map[float64]struct{}{}

	var part1 float64

	for {
		bio := Biodiversity(eris)
		if _, ok := bioSeen[bio]; ok {
			part1 = bio
			break
		} else {
			bioSeen[bio] = struct{}{}
		}

		eris = SimulateEris(eris)
	}

	return fmt.Sprintf("%.f", part1), ""
}

func SimulateEris(eris [][]rune) [][]rune {
	newEris := [][]rune{}
	newEris = append(newEris, []rune("......."))

	for y := 1; y < 6; y++ {
		l := "."

		for x := 1; x < 6; x++ {
			bugCount := 0

			for _, n := range aoc.Neighbours(x, y) {
				if eris[n.Y][n.X] == '#' {
					bugCount++
				}
			}

			if eris[y][x] == '#' && bugCount != 1 {
				l += "."
			} else if eris[y][x] == '.' && (bugCount == 1 || bugCount == 2) {
				l += "#"
			} else {
				l += string(eris[y][x])
			}
		}

		l += "."
		newEris = append(newEris, []rune(l))
	}
	newEris = append(newEris, []rune("......."))

	return newEris
}

func Biodiversity(eris [][]rune) float64 {
	bio := float64(0)

	for y := 1; y < 6; y++ {
		for x := 1; x < 6; x++ {
			if eris[y][x] == '#' {
				pow := ((y-1)*5 + (x - 1))
				bio += math.Pow(2, float64(pow))
			}
		}
	}

	return bio
}
