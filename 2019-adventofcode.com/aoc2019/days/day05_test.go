package days

import (
	"testing"
)

func TestDay05(t *testing.T) {
	d := TestDay(&Day05{}, t)
	d.WithFile("../input/day05.in", "15426686", "11430197")
}
