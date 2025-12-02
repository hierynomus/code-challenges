package aoc2025

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

const A2025D02Sample = `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`

func TestDay02_Sample(t *testing.T) {
	d := day.TestDay(t, Day02)
	d.WithInput(A2025D02Sample, "1227775554", "4174379265")
}
