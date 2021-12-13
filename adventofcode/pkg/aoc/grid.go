package aoc

import "strings"

func RenderRuneGrid(grid [][]rune) string {
	s := ""
	for _, l := range grid {
		s += string(l) + "\n"
	}

	return s
}

func CountRuneGridOccurrences(grid [][]rune, r rune) int {
	count := 0
	for _, l := range grid {
		for _, c := range l {
			if c == r {
				count += 1
			}
		}
	}

	return count
}

func RenderStringGrid(grid [][]string) string {
	s := ""
	for _, l := range grid {
		s += strings.Join(l, " ") + "\n"
	}

	return s
}
func RenderIntGridS(grid [][]int, sep string) string {
	s := ""
	for _, l := range grid {
		s += IntArrayAsString(l, sep) + "\n"
	}

	return s
}
