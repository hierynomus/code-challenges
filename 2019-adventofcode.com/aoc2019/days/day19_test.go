package days

import (
	"testing"
)

func TestDay19_Real(t *testing.T) {
	d := TestDay(&Day19{}, t)
	d.WithFile("../input/day19.in", "234", "9290812")
}
