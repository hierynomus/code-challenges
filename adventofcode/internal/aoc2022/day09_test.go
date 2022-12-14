package aoc2022

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

const A2022D09Sample = `R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2
`

func TestDay09_Sample(t *testing.T) {
	d := day.TestDay(t, Day09)
	d.WithInput(A2022D09Sample, "13", "1")
}

func TestDay09_LargeSample(t *testing.T) {
	sample := `R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20
`
	d := day.TestDay(t, Day09)
	d.WithInput(sample, "88", "36")
}
