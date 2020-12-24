package aoc2020

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

func TestDay18_sample(t *testing.T) {
	inp := `1 + (2 * 3) + (4 * (5 + 6))
`

	d := day.TestDay(t, Day18)
	d.WithInput(inp, "51", "51")
}

func TestDay18_sample2(t *testing.T) {
	inp := `1 + 2 * 3 + 4 * 5 + 6
`

	d := day.TestDay(t, Day18)
	d.WithInput(inp, "71", "231")
}
func TestDay18_sample3(t *testing.T) {
	inp := `5 + (8 * 3 + 9 + 3 * 4 * 3)
`

	d := day.TestDay(t, Day18)
	d.WithInput(inp, "437", "1445")
}

func TestDay18_sample4(t *testing.T) {
	inp := `((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2
`

	d := day.TestDay(t, Day18)
	d.WithInput(inp, "13632", "23340")
}
