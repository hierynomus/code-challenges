package aoc2020

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBag(t *testing.T) {
	l := "vibrant lavender bags contain 4 bright chartreuse bags, 3 dark teal bags, 4 muted aqua bags."

	bags := map[string]*Bag{}
	b := addBag(l, bags)

	assert.Equal(t, "vibrant lavender", b.Type)
	assert.Equal(t, 3, len(b.Contents))
	assert.Equal(t, 4, len(bags))
	assert.Equal(t, 1, len(bags["muted aqua"].ContainedIn))
	fmt.Printf("%v", b)
}
