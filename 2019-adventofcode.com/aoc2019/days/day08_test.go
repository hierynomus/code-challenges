package days

import (
	"testing"

	"github.com/hierynomus/aoc2019/aoc"
	"gotest.tools/v3/assert"
)

func TestHistogram(t *testing.T) {
	i := []int{1, 1, 2, 3, 4, 6, 8, 6, 3}
	h := aoc.MakeIntHistogram(i)
	assert.Equal(t, h[1], 2)
	assert.Equal(t, h[6], 2)
}

func TestDay08_Real(t *testing.T) {
	out := `
XXX  X     XX  X  X XXXX
X  X X    X  X X X  X
X  X X    X  X XX   XXX
XXX  X    XXXX X X  X
X X  X    X  X X X  X
X  X XXXX X  X X  X X
`
	d := TestDay(&Day08{}, t)
	d.WithFile("../input/day08.in", "1485", out)
}
