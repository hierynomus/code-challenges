package aoc2025

import (
	"bufio"
	"sort"
	"strconv"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type JunctionBox aoc.Point3D

type Circuit struct {
	boxes aoc.Set[JunctionBox]
}

func Day08(reader *bufio.Scanner) (string, string) {
	return ConnectJunctionBoxes(999)(reader)
}

func ConnectJunctionBoxes(nrConnections int) func(reader *bufio.Scanner) (string, string) {
	return func(reader *bufio.Scanner) (string, string) {
		var part1, part2 int
		lines := aoc.ReadStringArray(reader)
		circuits := map[JunctionBox]*Circuit{}
		for _, line := range lines {
			cs := strings.Split(line, ",")
			box := JunctionBox{
				X: aoc.ToInt64(cs[0]),
				Y: aoc.ToInt64(cs[1]),
				Z: aoc.ToInt64(cs[2]),
			}
			circuit := Circuit{
				boxes: aoc.NewSet([]JunctionBox{box}),
			}
			circuits[box] = &circuit
		}

		combos := aoc.Combinations(aoc.MapKeys(circuits))
		sort.Slice(combos, func(i, j int) bool {
			return aoc.Euclidian3D((aoc.Point3D)(combos[i][0]), (aoc.Point3D)(combos[i][1])) < aoc.Euclidian3D((aoc.Point3D)(combos[j][0]), (aoc.Point3D)(combos[j][1]))
		})

		connectedCircuits := aoc.Set[*Circuit]{}
		for _, c := range circuits {
			connectedCircuits.Add(c)
		}

		for c, combo := range combos {
			c1, c2 := circuits[combo[0]], circuits[combo[1]]
			if c1 == c2 {
				// fmt.Printf("Already connected: %v and %v (nr circuits: %d)\n", combo[0], combo[1], len(connectedCircuits))
				continue
			} else {
				// fmt.Printf("Connecting %v and %v (nr circuits: %d)\n", combo[0], combo[1], len(connectedCircuits))
				newCircuit := Circuit{boxes: aoc.Set[JunctionBox]{}}
				for b := range c1.boxes {
					newCircuit.boxes.Add(b)
				}
				for b := range c2.boxes {
					newCircuit.boxes.Add(b)
				}
				for b := range newCircuit.boxes {
					circuits[b] = &newCircuit
				}
				connectedCircuits.Delete(c1)
				connectedCircuits.Delete(c2)
				connectedCircuits.Add(&newCircuit)

				if len(connectedCircuits) == 1 { // Last circuit connected
					part2 = int(combo[0].X) * int(combo[1].X)
					break
				}
			}
			if c >= nrConnections-1 && part1 == 0 {
				cs := connectedCircuits.AsSlice()
				sort.Slice(cs, func(i int, j int) bool {
					return len(cs[i].boxes) > len(cs[j].boxes)
				})
				// fmt.Printf("Nr circuits: %d, Largest circuits sizes: %d, %d, %d\n", len(cs), len(cs[0].boxes), len(cs[1].boxes), len(cs[2].boxes))
				part1 = len(cs[0].boxes) * len(cs[1].boxes) * len(cs[2].boxes)
			}
		}

		return strconv.Itoa(part1), strconv.Itoa(part2)
	}
}
