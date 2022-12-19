package aoc2022

import (
	"bufio"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type Scan int

const (
	Empty Scan = iota
	Droplet
	Water
)

func Day18(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	blocks := map[aoc.Point3D]Scan{}

	for reader.Scan() {
		line := reader.Text()
		s := strings.Split(line, ",")
		p := aoc.Point3D{X: aoc.ToInt64(s[0]), Y: aoc.ToInt64(s[1]), Z: aoc.ToInt64(s[2])}
		blocks[p] = Droplet
	}

	for b := range blocks {
		for _, n := range b.Neighbours6() {
			if blocks[n] != Droplet {
				part1++
			}
		}
	}

	min, max := aoc.FindBoundingBox3D(maps.Keys(blocks))
	min, max = min.AddXYZ(-1, -1, -1), max.AddXYZ(1, 1, 1)

	// Flood fill the "world" with water, so that it surrounds the entire droplet (but does not go inside!)
	blocks[min] = Water
	changed := true
	for changed {
		changed = false
		for b, s := range blocks {
			if s == Water {
				for _, n := range b.Neighbours6() {
					if n.X < min.X || n.Y < min.Y || n.Z < min.Z || n.X > max.X || n.Y > max.Y || n.Z > max.Z {
						continue
					}

					if blocks[n] == Empty {
						blocks[n] = Water
						changed = true
					}
				}
			}
		}
	}

	// Check each block of droplet, if it has a neighbouring water block, it has a side exposed.
	for z := min.Z; z <= max.Z; z++ {
		for y := min.Y; y <= max.Y; y++ {
			for x := min.X; x <= max.X; x++ {
				p := aoc.Point3D{X: x, Y: y, Z: z}
				if blocks[p] != Droplet {
					continue
				}

				for _, n := range p.Neighbours6() {
					if blocks[n] == Water {
						// blocks[n] = Steam
						part2++
					}
				}
			}
		}
	}
	return strconv.Itoa(part1), strconv.Itoa(part2)
}
