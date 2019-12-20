package days

import (
	"testing"
)

func TestDay15_Real(t *testing.T) {
	d := TestDay(&Day15{}, t)
	d.WithFile("../input/day15.in", "258", "372")
}
