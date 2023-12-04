package aoc2022

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

var (
	North      = aoc.Point{X: 0, Y: -1}
	South      = aoc.Point{X: 0, Y: 1}
	East       = aoc.Point{X: 1, Y: 0}
	West       = aoc.Point{X: -1, Y: 0}
	Directions = []aoc.Point{North, South, West, East}
)

func ShouldMove(elfPositions aoc.PointSet, pos aoc.Point) bool {
	for _, p := range pos.Neighbours8() {
		if elfPositions.Contains(p) {
			return true
		}
	}

	return false
}

func CanMove(elfPositions aoc.PointSet, pos aoc.Point, dir aoc.Point) bool {
	switch dir {
	case North:
		return !(elfPositions.Contains(pos.AddXY(-1, -1)) || elfPositions.Contains(pos.AddXY(0, -1)) || elfPositions.Contains(pos.AddXY(1, -1)))
	case South:
		return !(elfPositions.Contains(pos.AddXY(-1, 1)) || elfPositions.Contains(pos.AddXY(0, 1)) || elfPositions.Contains(pos.AddXY(1, 1)))
	case East:
		return !(elfPositions.Contains(pos.AddXY(1, -1)) || elfPositions.Contains(pos.AddXY(1, 0)) || elfPositions.Contains(pos.AddXY(1, 1)))
	case West:
		return !(elfPositions.Contains(pos.AddXY(-1, -1)) || elfPositions.Contains(pos.AddXY(-1, 0)) || elfPositions.Contains(pos.AddXY(-1, 1)))
	default:
		panic("invalid direction")
	}
}

func ElfMoveRound(elfPositions aoc.PointSet, startWith int) (aoc.PointSet, bool) {
	newPositions := map[aoc.Point][]aoc.Point{}
	anyMoved := false

	for elf := range elfPositions {
		moved := false
		if ShouldMove(elfPositions, elf) {
			for i := startWith; i < startWith+4; i++ {
				dir := Directions[i%4]
				if CanMove(elfPositions, elf, dir) {
					newPos := elf.Add(dir)
					newPositions[newPos] = append(newPositions[newPos], elf)
					moved = true
					anyMoved = true
					break
				}
			}
		}
		if !moved {
			newPositions[elf] = append(newPositions[elf], elf)
		}
	}

	moved := aoc.PointSet{}
	for pos, elves := range newPositions {
		if len(elves) == 1 {
			moved.Add(pos)
		} else {
			for _, elf := range elves {
				moved.Add(elf)
			}
		}
	}

	return moved, anyMoved
}

func Day23(reader *bufio.Scanner) (string, string) {
	var part1, part2 int
	elfPositions := aoc.PointSet{}
	inp := aoc.ReadRuneGrid(reader)
	for y := 0; y < len(inp); y++ {
		for x := 0; x < len(inp[y]); x++ {
			if inp[y][x] == '#' {
				elfPositions.Add(aoc.Point{X: x, Y: y})
			}
		}
	}

	for i := 0; i < 10; i++ {
		elfPositions, _ = ElfMoveRound(elfPositions, i)
	}

	min, max := elfPositions.BoundingBox()
	part1 = (max.X-min.X+1)*(max.Y-min.Y+1) - len(elfPositions)

	i := 10
	moved := false
	for {
		elfPositions, moved = ElfMoveRound(elfPositions, i)
		if !moved {
			break
		}
		i++
	}

	part2 = i + 1
	return strconv.Itoa(part1), strconv.Itoa(part2)
}
