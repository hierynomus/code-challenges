package aoc2022

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
	"github.com/stretchr/testify/assert"
)

const A2022D25Sample = `1=-0-2
12111
2=0=
21
2=01
111
20012
112
1=-1=
1-12
12
1=
122
`

func TestToSnafu(t *testing.T) {
	assert.Equal(t, "1", ToSnafu(1))
	assert.Equal(t, "2", ToSnafu(2))
	assert.Equal(t, "1=", ToSnafu(3))
	assert.Equal(t, "1-", ToSnafu(4))
	assert.Equal(t, "10", ToSnafu(5))
	assert.Equal(t, "11", ToSnafu(6))
	assert.Equal(t, "12", ToSnafu(7))
	assert.Equal(t, "2=", ToSnafu(8))
	assert.Equal(t, "2-", ToSnafu(9))
	assert.Equal(t, "20", ToSnafu(10))
	assert.Equal(t, "1=0", ToSnafu(15))
	assert.Equal(t, "1-0", ToSnafu(20))
	assert.Equal(t, "1=11-2", ToSnafu(2022))
	assert.Equal(t, "1-0---0", ToSnafu(12345))
	assert.Equal(t, "1121-1110-1=0", ToSnafu(314159265))
}

func TestParseSnafu(t *testing.T) {
	assert.Equal(t, 1, ParseSnafu("1"))
	assert.Equal(t, 2, ParseSnafu("2"))
	assert.Equal(t, 3, ParseSnafu("1="))
	assert.Equal(t, 4, ParseSnafu("1-"))
	assert.Equal(t, 5, ParseSnafu("10"))
	assert.Equal(t, 6, ParseSnafu("11"))
	assert.Equal(t, 7, ParseSnafu("12"))
	assert.Equal(t, 8, ParseSnafu("2="))
	assert.Equal(t, 9, ParseSnafu("2-"))
	assert.Equal(t, 10, ParseSnafu("20"))
	assert.Equal(t, 15, ParseSnafu("1=0"))
	assert.Equal(t, 20, ParseSnafu("1-0"))
	assert.Equal(t, 2022, ParseSnafu("1=11-2"))
	assert.Equal(t, 12345, ParseSnafu("1-0---0"))
	assert.Equal(t, 314159265, ParseSnafu("1121-1110-1=0"))
}

func TestDay25_Sample(t *testing.T) {
	d := day.TestDay(t, Day25)
	d.WithInput(A2022D25Sample, "2=-1=0", "0")
}
