package aoc

import (
	"strconv"
	"strings"
)

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
