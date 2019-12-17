package days

import (
	"testing"
)

func TestDay17_Real(t *testing.T) {
	d := TestDay(&Day17{}, t)
	d.WithFile("../input/day17.in", "3888", "927809")
}
