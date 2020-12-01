package aoc

func HammingDistance(s1, s2 string) int {
	d := 0
	for i, c := range s1 {
		if c != rune(s2[i]) {
			d += 1
		}
	}
	return d
}
