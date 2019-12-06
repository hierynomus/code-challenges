package days

import "testing"

import "gotest.tools/v3/assert"

func TestIsAsc(t *testing.T) {
	assert.Check(t, isAsc(123))
	assert.Check(t, !isAsc(121))
	assert.Check(t, isAsc(12))
	assert.Check(t, isAsc(111111))
	assert.Check(t, isAsc(135679))
}

func TestHasPair(t *testing.T) {
	assert.Check(t, !hasPair(123))
	assert.Check(t, hasPair(112))
	assert.Check(t, !hasPair(12))
	assert.Check(t, hasPair(111111))
	assert.Check(t, !hasPair(135679))
	assert.Check(t, hasPair(1356679))
}
func TestHasExactPair(t *testing.T) {
	assert.Check(t, !hasExactPair(123))
	assert.Check(t, hasExactPair(112))
	assert.Check(t, !hasExactPair(12))
	assert.Check(t, !hasExactPair(111111))
	assert.Check(t, hasExactPair(11122111))
	assert.Check(t, !hasExactPair(135679))
	assert.Check(t, hasExactPair(1356679))
}

func TestDay04(t *testing.T) {
	d := TestDay(&Day04{}, t)
	d.WithFile("../input/day04.in", "945", "617")
}
