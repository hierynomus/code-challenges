package aoc2020

import (
	"bufio"
	"strings"
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
	"github.com/stretchr/testify/assert"
)

func TestDay11_sample(t *testing.T) {
	inp := `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL
`

	d := day.TestDay(t, Day11)
	d.WithInput(inp, "37", "26")
}

func TestDay11_countOccupied2(t *testing.T) {
	inp := aoc.ReadRuneGrid(bufio.NewScanner(strings.NewReader(`#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##
`)))
	assert.Equal(t, 5, countOccupied2(inp, 0, 2))
}
func TestDay11_countOccupied2_2(t *testing.T) {
	inp := aoc.ReadRuneGrid(bufio.NewScanner(strings.NewReader(`#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##
`)))
	assert.Equal(t, 4, countOccupied2(inp, 0, 1))
}
