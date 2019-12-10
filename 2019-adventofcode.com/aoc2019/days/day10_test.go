package days

import (
	"bufio"
	"strings"
	"testing"

	"github.com/hierynomus/aoc2019/aoc"
	"gotest.tools/v3/assert"
)

func TestDay10_Real(t *testing.T) {
	d := TestDay(&Day10{}, t)
	d.WithFile("../input/day10.in", "263", "1110")
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
	d := &Day10{asteroids: AsteroidBelt{}}
	d.BuildAsteroidBelt(bufio.NewScanner(strings.NewReader(input)))
	p := d.FindMostLinesOfSight()
	assert.Equal(t, len(d.asteroids[p]), 210)
	assert.DeepEqual(t, p, &aoc.Point{X: 11, Y: 13})
}
func TestDay10_2(t *testing.T) {
	input := `.#..#
.....
#####
....#
...##`
	d := &Day10{asteroids: AsteroidBelt{}}
	d.BuildAsteroidBelt(bufio.NewScanner(strings.NewReader(input)))
	p := d.FindMostLinesOfSight()
	assert.Equal(t, len(d.asteroids[p]), 8)
	assert.DeepEqual(t, p, &aoc.Point{X: 3, Y: 4})
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
	d := &Day10{asteroids: AsteroidBelt{}}
	d.BuildAsteroidBelt(bufio.NewScanner(strings.NewReader(input)))
	p := d.FindMostLinesOfSight()
	assert.Equal(t, len(d.asteroids[p]), 33)
	assert.DeepEqual(t, p, &aoc.Point{X: 5, Y: 8})
}
