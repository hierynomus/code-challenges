package aoc2019

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

func TestDay11_Real(t *testing.T) {
	out2 := `
...##.####.#....####.####..##..#..#.###....
....#.#....#....#....#....#..#.#..#.#..#...
....#.###..#....###..###..#....####.#..#...
....#.#....#....#....#....#.##.#..#.###....
.#..#.#....#....#....#....#..#.#..#.#......
..##..####.####.####.#.....###.#..#.#......
`
	d := day.TestDay(t, Day11)
	d.WithFile("../../input/aoc2019/day11.in", "2172", out2)
}
