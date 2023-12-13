package aoc

func Transpose(slice [][]int) [][]int {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]int, xl)

	for i := range result {
		result[i] = make([]int, yl)
	}

	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}

	return result
}

func TransposeString(slice []string) []string {
	xl := len(slice[0])
	yl := len(slice)
	result := make([]string, xl)

	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i] += string(slice[j][i])
		}
	}

	return result
}
