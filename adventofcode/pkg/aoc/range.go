package aoc

func Range(start, end int) []int {
	res := []int{}
	for x := start; x < end; x++ {
		res = append(res, x)
	}

	return res
}
