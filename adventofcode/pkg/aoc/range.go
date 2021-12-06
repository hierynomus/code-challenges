package aoc

func Range(start, end int) []int {
	res := []int{}
	if start > end {
		for x := start; x > end; x-- {
			res = append(res, x)
		}
	} else {
		for x := start; x < end; x++ {
			res = append(res, x)
		}
	}
	return res
}

func RangeIncl(start, end int) []int {
	res := []int{}
	if start > end {
		for x := start; x >= end; x-- {
			res = append(res, x)
		}
	} else {
		for x := start; x <= end; x++ {
			res = append(res, x)
		}
	}
	return res
}
