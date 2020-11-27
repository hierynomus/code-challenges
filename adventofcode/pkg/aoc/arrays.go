package aoc

import (
	"strconv"
	"strings"
)

func IntArrayAsString(arr []int, sep string) string {
	if len(arr) == 0 {
		return ""
	}

	s := strconv.Itoa(arr[0])

	for _, i := range arr[1:] {
		s += sep
		s += strconv.Itoa(i)
	}

	return s
}

func AsIntArray(line string) []int {
	arr := strings.Split(line, ",")
	iArr := make([]int, len(arr))

	for i, c := range arr {
		n, err := strconv.Atoi(c)
		if err != nil {
			panic(err)
		}

		iArr[i] = n
	}

	return iArr
}

func AsRuneArray(line string) ([]rune, error) {
	return []rune(line), nil
}

func AsStringArray(line string) ([]string, error) {
	return strings.Split(line, ","), nil
}

func StringArrayContains(haystack []string, needle string) bool {
	for _, s := range haystack {
		if s == needle {
			return true
		}
	}

	return false
}

func StringArrayIndex(haystack []string, needle string) int {
	for i, s := range haystack {
		if s == needle {
			return i
		}
	}

	return -1
}

func IntArrayIndex(haystack []int, needle int) int {
	for i, s := range haystack {
		if s == needle {
			return i
		}
	}

	return -1
}
