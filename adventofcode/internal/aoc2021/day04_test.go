package aoc2021

import (
	"bufio"
	"strings"
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
	"github.com/stretchr/testify/assert"
)

func TestDay04_sample(t *testing.T) {
	inp := `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7`

	d := day.TestDay(t, Day04)
	d.WithInput(inp, "4512", "1924")
}

func TestBingoH(t *testing.T) {
	inp := `X X X X X
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19
`
	c := readBingo(bufio.NewScanner(strings.NewReader(inp)))
	assert.True(t, c.isBingo())
}

func TestBingoV(t *testing.T) {
	inp := `X X 3 X X
 X  X 23  4 24
21  X 14 16  7
 6  X  3 18  5
 1  X 20 15 19
`
	c := readBingo(bufio.NewScanner(strings.NewReader(inp)))
	assert.True(t, c.isBingo())
}
