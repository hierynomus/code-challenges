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

func ToIntE(s string) (int, error) {
	i, err := strconv.Atoi(strings.TrimSpace(s))

	return i, err
}

func BinaryToInt8(s string) int {
	i, err := strconv.ParseInt(s, 2, 8)
	if err != nil {
		panic(err)
	}
	return int(i)
}

func BinaryToInt16(s string) int {
	i, err := strconv.ParseInt(s, 2, 16)
	if err != nil {
		panic(err)
	}
	return int(i)
}

func BinaryToInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		panic(err)
	}
	return i
}

func ToInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}

	return i
}

func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

func Pow(x, y int) int {
	if y == 0 {
		return 1
	}

	if y%2 == 0 {
		return Pow(x, y/2) * Pow(x, y/2)
	}

	return x * Pow(x, y/2) * Pow(x, y/2)
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

func MinOf(i1, i2 int) int {
	if i1 < i2 {
		return i1
	}

	return i2
}

func MaxOf(i1, i2 int) int {
	if i1 > i2 {
		return i1
	}

	return i2
}

func Sign(i int) int {
	if i < 0 {
		return -1
	}

	return 1
}

func ParseBin(s string) int64 {
	i, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		panic(err)
	}
	return i
}
