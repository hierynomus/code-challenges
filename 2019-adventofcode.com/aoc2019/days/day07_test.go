package days

import (
	"testing"
)

func TestDay07_Real(t *testing.T) {
	d := TestDay(&Day07{}, t)
	d.WithFile("../input/day07.in", "929800", "15432220")
}
