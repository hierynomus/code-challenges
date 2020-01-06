package days

import (
	"testing"
)

func TestDay24_Real(t *testing.T) {
	d := TestDay(&Day24{}, t)
	d.WithFile("../input/day24.in", "18370591", "")
}
