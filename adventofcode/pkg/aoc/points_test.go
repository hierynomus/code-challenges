package aoc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoint3D(t *testing.T) {
	assert.Equal(t, Point3D{0, 1, 2}, Point3D{0, 1, 2})
}
