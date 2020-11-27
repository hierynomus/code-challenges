package aoc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLcm(t *testing.T) {
	assert.Equal(t, Lcm(6, 10), int64(30))
	assert.Equal(t, Lcm(10, 6), int64(30))
}
