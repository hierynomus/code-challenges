package days

import (
	"testing"
)

func TestDay13_Real(t *testing.T) {
	d := TestDay(&Day13{}, t)
	d.WithFile("../input/day13.in", "361", "17590")
}
