package aoc

func BinSearchInt(start, end int, f func(int) bool) int {
	for start < end {
		mid := (start + end) / 2
		if f(mid) {
			end = mid
		} else {
			start = mid + 1
		}
	}

	return start
}
