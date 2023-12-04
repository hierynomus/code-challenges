package aoc2021

import (
	"bufio"
	"strings"
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
	"github.com/stretchr/testify/assert"
)

func TestDay11_sample(t *testing.T) {
	inp := `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`

	d := day.TestDay(t, Day11)
	d.WithInput(inp, "1656", "195")
}

func TestDay11_1step(t *testing.T) {
	inp := `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`

	reader := bufio.NewScanner(strings.NewReader(inp))
	grid := make([][]int, 0)
	for reader.Scan() {
		line := aoc.AsIntArrayS(reader.Text(), "")
		grid = append(grid, line)
	}

	g, _ := OctopusStep(grid)
	assert.Equal(t, `6594254334
3856965822
6375667284
7252447257
7468496589
5278635756
3287952832
7993992245
5957959665
6394862637
`, aoc.RenderIntGridS(g, ""))
}

func TestDay11_2step(t *testing.T) {
	inp := `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`

	reader := bufio.NewScanner(strings.NewReader(inp))
	grid := make([][]int, 0)
	for reader.Scan() {
		line := aoc.AsIntArrayS(reader.Text(), "")
		grid = append(grid, line)
	}

	g, _ := OctopusStep(grid)
	g, f := OctopusStep(g)
	assert.Equal(t, 35, f)
	assert.Equal(t, `8807476555
5089087054
8597889608
8485769600
8700908800
6600088989
6800005943
0000007456
9000000876
8700006848
`, aoc.RenderIntGridS(g, ""))
}
