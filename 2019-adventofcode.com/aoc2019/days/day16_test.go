package days

import (
	"testing"
)

func TestDay16_Real(t *testing.T) {
	d := TestDay(&Day16{}, t)
	d.WithFile("../input/day16.in", "90744714", "")
}
