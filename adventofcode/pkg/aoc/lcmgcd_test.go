package aoc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGcd(t *testing.T) {
	assert.Equal(t, Gcd(10, 5), int64(5))
	assert.Equal(t, Gcd(17, 5), int64(1))
	assert.Equal(t, Gcd(120, 32), int64(8))
}

func TestLcm(t *testing.T) {
	assert.Equal(t, Lcm(6, 10), int64(30))
	assert.Equal(t, Lcm(10, 6), int64(30))
}
