package days

import (
	"testing"

	"github.com/hierynomus/aoc2019/aoc"
	"github.com/magiconair/properties/assert"
)

func TestHistogram(t *testing.T) {
	i := []int{1, 1, 2, 3, 4, 6, 8, 6, 3}
	h := aoc.MakeIntHistogram(i)
	assert.Equal(t, h[1], 2)
	assert.Equal(t, h[6], 2)
}

func TestDay08_Real(t *testing.T) {
	d := TestDay(&Day08{}, t)
	d.WithFile("../input/day08.in", "1485", "")
}
