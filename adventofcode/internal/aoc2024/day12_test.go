package aoc2024

import (
	"strconv"
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
	"github.com/stretchr/testify/assert"
)

const A2024D12Sample = `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`

func TestDay12_Sample(t *testing.T) {
	d := day.TestDay(t, Day12)
	d.WithInput(A2024D12Sample, "1930", "1206")
}

func TestDay12_Corners(t *testing.T) {
	sets := []struct {
		Region   aoc.PointSet
		Expected int
	}{
		{
			aoc.NewPointSet([]aoc.Point{{X: 0, Y: 0}}), 4,
		},
		{
			aoc.NewPointSet([]aoc.Point{{X: 0, Y: 0}, {X: 0, Y: 1}}), 4,
		},
		{
			aoc.NewPointSet([]aoc.Point{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}}), 6,
		},
	}

	for i, s := range sets {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			c := countCorners(s.Region)
			assert.Equal(t, s.Expected, c)
		})
	}
}
