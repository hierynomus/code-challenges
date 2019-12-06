package days

import (
	"testing"
)

func TestDay01(t *testing.T) {
	d := TestDay(&Day01{}, t)
	d.WithFile("../input/day01.in", "3388015", "5079140")
}
