package aoc2023

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

const A2023D15Sample = `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7
`

func TestDay15_Sample(t *testing.T) {
	d := day.TestDay(t, Day15)
	d.WithInput(A2023D15Sample, "1320", "145")
}
