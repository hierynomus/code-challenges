package aoc2025

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

const A2025D01Sample = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

func TestDay01_Sample(t *testing.T) {
	d := day.TestDay(t, Day01)
	d.WithInput(A2025D01Sample, "3", "6")
}

func TestDay01_R1000(t *testing.T) {
	d := day.TestDay(t, Day01)
	d.WithInput(`R1000`, "0", "10")
}

func TestDay01_L1000(t *testing.T) {
	d := day.TestDay(t, Day01)
	d.WithInput(`L1000`, "0", "10")
}

func TestDay01_L1050(t *testing.T) {
	d := day.TestDay(t, Day01)
	d.WithInput(`L1050`, "1", "11")
}
