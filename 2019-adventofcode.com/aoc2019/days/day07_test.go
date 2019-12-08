package days

import (
	"testing"
)

func TestDay07(t *testing.T) {
	d := TestDay(&Day07{}, t)
	d.WithInput("3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0", "43210", "")
}
func TestDay07_2(t *testing.T) {
	d := TestDay(&Day07{}, t)
	d.WithInput("3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5", "0", "61696857")
}
func TestDay07_Real(t *testing.T) {
	d := TestDay(&Day07{}, t)
	d.WithFile("../input/day07.in", "929800", "15432220")
}
