package days

import (
	"testing"
)

func TestDay03(t *testing.T) {
	d := TestDay(&Day03{}, t)
	d.WithFile("../input/day03.in", "446", "9006")
}
