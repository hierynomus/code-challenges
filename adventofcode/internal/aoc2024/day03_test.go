package aoc2024

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

const A2024D03Sample = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`

func TestDay03_Sample(t *testing.T) {
	d := day.TestDay(t, Day03)
	d.WithInput(A2024D03Sample, "161", "48")
}
