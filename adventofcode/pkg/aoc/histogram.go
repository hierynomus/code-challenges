package aoc

type IntHistogram map[int]int
type RuneHistogram map[rune]int

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
