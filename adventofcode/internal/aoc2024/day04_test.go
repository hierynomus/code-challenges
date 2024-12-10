package aoc2024

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

const A2024D04Sample = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func TestDay04_Sample(t *testing.T) {
	d := day.TestDay(t, Day04)
	d.WithInput(A2024D04Sample, "18", "4")
}
