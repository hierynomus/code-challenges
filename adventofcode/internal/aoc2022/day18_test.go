package aoc2022

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

const A2022D18Sample = `2,2,2
1,2,2
3,2,2
2,1,2
2,3,2
2,2,1
2,2,3
2,2,4
2,2,6
1,2,5
3,2,5
2,1,5
2,3,5
`

func TestDay18_Sample(t *testing.T) {
	d := day.TestDay(t, Day18)
	d.WithInput(A2022D18Sample, "64", "58")
}
