package aoc

import "strconv"

func Abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}

func Abs64(i int64) int64 {
	if i < 0 {
		return -i
	}

	return i
}

func ToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}
