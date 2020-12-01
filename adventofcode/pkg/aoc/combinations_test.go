package aoc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntCombinationsN_2(t *testing.T) {
	in := []int{1, 2, 3, 4}
	exp := [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}
	collector := [][]int{}
	for c := range IntCombinationsN(in, 2) {
		collector = append(collector, c)
	}

	assert.Equal(t, exp, collector)
}

func TestIntCombinationsN_3(t *testing.T) {
	in := []int{1, 2, 3, 4}
	exp := [][]int{{1, 2, 3}, {1, 2, 4}, {1, 3, 4}, {2, 3, 4}}
	collector := [][]int{}
	for c := range IntCombinationsN(in, 3) {
		collector = append(collector, c)
	}

	assert.Equal(t, exp, collector)
}
