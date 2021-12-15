package aoc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoint3D(t *testing.T) {
	assert.Equal(t, Point3D{0, 1, 2}, Point3D{0, 1, 2})
}

func TestMap(t *testing.T) {
	m := map[Point]int{}
	m[Point{0, 0}] = 1
	assert.Equal(t, 1, m[NewPoint(0, 0)])
}
