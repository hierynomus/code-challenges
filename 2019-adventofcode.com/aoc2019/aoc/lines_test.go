package aoc

import "testing"

import "gotest.tools/v3/assert"

func TestShouldIntersect(t *testing.T) {
	l1 := Line{Point{0, 0}, Point{10, 0}}
	l2 := Line{Point{5, 5}, Point{5, -5}}

	intersect := l1.Intersection(l2)
	assert.Equal(t, intersect, Point{5, 0})
}

func TestIsInSegmentVertical(t *testing.T) {
	l1 := Line{Point{0, 0}, Point{0, 10}}
	assert.Check(t, l1.Contains(Point{0, 5}))
	assert.Check(t, !l1.Contains(Point{0, 11}))
	assert.Check(t, !l1.Contains(Point{0, -1}))
}

func TestIsInSegmentHorizontal(t *testing.T) {
	l1 := Line{Point{20, 0}, Point{10, 0}}
	assert.Check(t, l1.Contains(Point{15, 0}))
	assert.Check(t, !l1.Contains(Point{9, 0}))
	assert.Check(t, !l1.Contains(Point{21, 0}))
}
