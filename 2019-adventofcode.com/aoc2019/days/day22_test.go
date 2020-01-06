package days

import (
	"testing"
)

func TestDay22_Real(t *testing.T) {
	d := TestDay(&Day22{}, t)
	d.WithFile("../input/day22.in", "18370591", "")
}
