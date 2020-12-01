package aoc

func StringCombinations(s []string) [][]string {
	arr := [][]string{}

	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			arr = append(arr, []string{s[i], s[j]})
		}
	}

	return arr
}

func IntCombinations(s []int) [][]int {
	arr := [][]int{}

	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			arr = append(arr, []int{s[i], s[j]})
		}
	}

	return arr
}
