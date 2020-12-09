package aoc

import (
	"strconv"
	"strings"
)

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
	i, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		panic(err)
	}

	return i
}

func Sum(l []int) int {
	s := 0
	for _, x := range l {
		s += x
	}

	return s
}

func Max(l []int) int {
	m := l[0]
	for _, x := range l {
		if x > m {
			m = x
		}
	}

	return m
}

func Min(l []int) int {
	m := l[0]
	for _, x := range l {
		if x < m {
			m = x
		}
	}

	return m
}
