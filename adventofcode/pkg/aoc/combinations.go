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

func IntCombinationsN(list []int, n int) <-chan []int {
	c := make(chan []int, 0)

	go func() {
		res := make([]int, n)
		for i, n := range list {
			res[0] = n
			generateIntCombinationN(c, res, list, i, 1)
		}
		close(c)
	}()

	return c
}

func generateIntCombinationN(c chan []int, out []int, in []int, idx int, sz int) {
	if sz == len(out) {
		res := make([]int, sz)
		copy(res, out)
		c <- res
		return
	}

	for i := idx + 1; i < len(in); i++ {
		out[sz] = in[i]
		generateIntCombinationN(c, out, in, i, sz+1)
	}
}
