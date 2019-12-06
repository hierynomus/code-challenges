package days

import (
	"testing"
)

func TestDay06(t *testing.T) {
	d := TestDay(&Day06{}, t)
	d.WithFile("../input/day06.in", "117672", "277")
}
