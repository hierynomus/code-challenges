package aoc2018

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day03(reader *bufio.Scanner) (string, string) {
	claims := map[aoc.Point][]int{}
	allClaims := aoc.IntSet{}
	for reader.Scan() {
		c := reader.Text()
		split := strings.Split(c, " ")
		claim := aoc.ToInt(split[0][1:])
		allClaims.Add(claim)
		sxy := strings.Split(split[2][0:len(split[2])-1], ",")
		dxy := strings.Split(split[3], "x")

		sx, sy := aoc.ToInt(sxy[0]), aoc.ToInt(sxy[1])
		for dx := 0; dx < aoc.ToInt(dxy[0]); dx++ {
			for dy := 0; dy < aoc.ToInt(dxy[1]); dy++ {
				p := aoc.NewPoint(sx+dx, sy+dy)
				if _, ok := claims[p]; !ok {
					claims[p] = []int{claim}
				} else {
					claims[p] = append(claims[p], claim)
				}
			}
		}
	}

	part1 := 0
	for _, v := range claims {
		if len(v) > 1 {
			part1 += 1
		}
	}

	for _, v := range claims {
		if len(v) > 1 {
			for _, c := range v {
				allClaims.Delete(c)
			}
		}
	}

	part2 := 0
	for k := range allClaims {
		part2 = k
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}
