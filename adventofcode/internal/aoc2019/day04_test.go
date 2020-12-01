package aoc2019

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAsc(t *testing.T) {
	assert.True(t, isAsc(123))
	assert.True(t, !isAsc(121))
	assert.True(t, isAsc(12))
	assert.True(t, isAsc(111111))
	assert.True(t, isAsc(135679))
}

func TestHasPair(t *testing.T) {
	assert.True(t, !hasPair(123))
	assert.True(t, hasPair(112))
	assert.True(t, !hasPair(12))
	assert.True(t, hasPair(111111))
	assert.True(t, !hasPair(135679))
	assert.True(t, hasPair(1356679))
}
func TestHasExactPair(t *testing.T) {
	assert.True(t, !hasExactPair(123))
	assert.True(t, hasExactPair(112))
	assert.True(t, !hasExactPair(12))
	assert.True(t, !hasExactPair(111111))
	assert.True(t, hasExactPair(11122111))
	assert.True(t, !hasExactPair(135679))
	assert.True(t, hasExactPair(1356679))
}
