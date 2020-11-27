package aoc2019

import (
	"bufio"
	"strings"
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
	"github.com/stretchr/testify/assert"
)

func TestDay10_Real(t *testing.T) {
	d := day.TestDay(t, Day10)
	d.WithFile("../../input/aoc2019/day10.in", "263", "1110")
}

func TestDay10_1(t *testing.T) {
	input := `#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##`
	d := AsteroidBelt{}
	d.BuildAsteroidBelt(bufio.NewScanner(strings.NewReader(input)))
	p := d.FindMostLinesOfSight()
	assert.Equal(t, len(d[p]), 210)
	assert.Equal(t, p, &aoc.Point{X: 11, Y: 13})
}
func TestDay10_2(t *testing.T) {
	input := `.#..#
.....
#####
....#
...##`
	d := AsteroidBelt{}
	d.BuildAsteroidBelt(bufio.NewScanner(strings.NewReader(input)))
	p := d.FindMostLinesOfSight()
	assert.Equal(t, len(d[p]), 8)
	assert.Equal(t, p, &aoc.Point{X: 3, Y: 4})
}
func TestDay10_3(t *testing.T) {
	input := `......#.#.
#..#.#....
..#######.
.#.#.###..
.#..#.....
..#....#.#
#..#....#.
.##.#..###
##...#..#.
.#....####`
	d := AsteroidBelt{}
	d.BuildAsteroidBelt(bufio.NewScanner(strings.NewReader(input)))
	p := d.FindMostLinesOfSight()
	assert.Equal(t, len(d[p]), 33)
	assert.Equal(t, p, &aoc.Point{X: 5, Y: 8})
}
