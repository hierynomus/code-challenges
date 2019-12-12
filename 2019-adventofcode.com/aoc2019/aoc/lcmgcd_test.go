package aoc

import "testing"

import "gotest.tools/v3/assert"

func TestLcm(t *testing.T) {
	assert.Equal(t, Lcm(6, 10), 30)
	assert.Equal(t, Lcm(10, 6), 30)
}
