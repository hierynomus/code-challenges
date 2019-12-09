package days

import (
	"testing"
)

func TestDay09_Real(t *testing.T) {
	d := TestDay(&Day09{}, t)
	d.WithFile("../input/day09.in", "2427443564", "87221")
}
