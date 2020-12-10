package aoc2020

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

func TestDay10_sample(t *testing.T) {
	inp := `16
10
15
5
1
11
7
19
6
12
4`

	d := day.TestDay(t, Day10)
	d.WithInput(inp, "35", "8")
}
func TestDay10_sample2(t *testing.T) {
	inp := `28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3
`

	d := day.TestDay(t, Day10)
	d.WithInput(inp, "220", "19208")
}
