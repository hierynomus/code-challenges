package aoc

import "fmt"

type IntHistogram map[int]int
type RuneHistogram map[rune]int64
type StringHistogram map[string]int

func MakeIntHistogram(list []int) IntHistogram {
	h := IntHistogram{}

	h.Adds(list)

	return h
}

func (h IntHistogram) Add(i int) {
	h[i]++
}

func (h IntHistogram) Adds(list []int) {
	for _, i := range list {
		h[i]++
	}
}

func (h IntHistogram) Max() (int, int) {
	key, max := 0, 0
	for k, v := range h {
		if v > max {
			key, max = k, v
		}
	}

	return key, max
}

func MakeRuneHistogram(list []rune) RuneHistogram {
	h := RuneHistogram{}

	for _, i := range list {
		h[i]++
	}

	return h
}

func (h RuneHistogram) Max() rune {
	k, _ := h.MaxC()
	return k
}

func (h RuneHistogram) Min() rune {
	k, _ := h.MinC()
	return k
}

func (h RuneHistogram) MaxC() (rune, int64) {
	var key rune
	var max int64
	for k, v := range h {
		if v > max {
			key, max = k, v
		}
	}

	return key, max
}

func (h RuneHistogram) MinC() (rune, int64) {
	var key rune
	var min int64 = 1 << 62
	for k, v := range h {
		if v < min {
			key, min = k, v
		}
	}

	return key, min
}

func MakeStringHistogram(list []string) StringHistogram {
	h := StringHistogram{}

	h.Adds(list)

	return h
}

func (h StringHistogram) Add(i string) {
	h[i]++
}

func (h StringHistogram) Adds(list []string) {
	for _, i := range list {
		h[i]++
	}
}

func (h StringHistogram) String() string {
	return fmt.Sprintf("%v", (map[string]int)(h))
}
