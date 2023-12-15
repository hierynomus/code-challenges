package aoc2023

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type Lens struct {
	Label string
	Focal int
}

type Box []*Lens

func Day15(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	lines := aoc.ReadStringArray(reader)

	boxes := map[int]Box{}

	for _, line := range lines {
		instructions := strings.Split(line, ",")
		for _, in := range instructions {
			part1 += Hash(in)

			if strings.Contains(in, "=") {
				prt := strings.Split(in, "=")
				lh := Hash(prt[0])
				if _, ok := boxes[lh]; !ok {
					boxes[lh] = Box{}
				}

				box := boxes[lh]

				lens := &Lens{Label: prt[0], Focal: aoc.ToInt(prt[1])}
				found := false
				for _, l := range box {
					if l.Label == lens.Label {
						l.Focal = lens.Focal
						found = true
						break
					}
				}

				if !found {
					boxes[lh] = append(box, lens)
				}
			} else if strings.Contains(in, "-") {
				label := in[:len(in)-1]
				lh := Hash(label)
				if box, ok := boxes[lh]; ok {
					for i, l := range box {
						if l.Label == label {
							box = append(box[:i], box[i+1:]...)
							break
						}
					}

					boxes[lh] = box
				}
			}
		}
	}

	for i, box := range boxes {
		for j, l := range box {
			m := i + 1
			m *= j + 1
			m *= l.Focal
			part2 += m
		}
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func Hash(instruction string) int {
	h := 0
	for _, c := range instruction {
		h += int(c)
		h *= 17
		h %= 256
	}

	return h
}
