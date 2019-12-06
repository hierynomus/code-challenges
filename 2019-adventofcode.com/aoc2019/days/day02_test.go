package days

import (
	"testing"
)

func TestDay02(t *testing.T) {
	d := TestDay(&Day02{}, t)
	d.WithFile("../input/day02.in", "4138687", "6635")
}
