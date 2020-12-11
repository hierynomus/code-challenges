package aoc

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
		for _, c := range []rune(l) {
			if c == r {
				count += 1
			}
		}
	}

	return count
}
