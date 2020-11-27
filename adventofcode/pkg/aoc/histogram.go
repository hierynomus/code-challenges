package aoc

type IntHistogram map[int]int
type RuneHistogram map[rune]int

func MakeIntHistogram(list []int) IntHistogram {
	h := IntHistogram{}

	for _, i := range list {
		h[i]++
	}

	return h
}

func MakeRuneHistogram(list []rune) RuneHistogram {
	h := RuneHistogram{}

	for _, i := range list {
		h[i]++
	}

	return h
}
