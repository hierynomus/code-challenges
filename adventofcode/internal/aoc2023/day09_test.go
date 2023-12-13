package aoc2023

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

const A2023D09Sample = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45
`

func TestDay09_Sample(t *testing.T) {
	d := day.TestDay(t, Day09)
	d.WithInput(A2023D09Sample, "114", "2")
}
