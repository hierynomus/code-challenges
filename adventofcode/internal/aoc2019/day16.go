package aoc2019

import (
	"bufio"
	"fmt"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day16(scanner *bufio.Scanner) (string, string) {
	if !scanner.Scan() {
		panic(fmt.Errorf("could not read input"))
	}

	line := scanner.Text()
	input := []int{}

	for _, c := range line {
		input = append(input, int(c-'0'))
	}

	p1 := make([]int, len(input))
	copy(p1, input)

	for i := 0; i < 100; i++ {
		p1 = roundA(p1)
	}

	// l2 := strings.Repeat(line, 10000)
	// input2 := make([]int, len(l2))
	// for i, c := range l2 {
	// 	input[i] = int(c - '0')
	// }

	// offset = strconv.Itoa(l2[0:7])
	// for i := 0; i < 100; i++ {
	// 	input2 = roundB(input2, offset)
	// }
	return aoc.IntArrayAsString(p1[0:8], ""), ""
}

func roundA(arr []int) []int {
	out := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		out[i] = fft(i, arr)
	}

	return out
}

func fft(idx int, arr []int) int {
	sum := 0
	rpt := 1 + idx
	patIdx := 1

	for i := 0; i < len(arr); i++ {
		pat := (patIdx / rpt) % 4
		if pat == 1 { // 0, *1*, 0, -1
			sum += arr[i]
		} else if pat == 3 { // 0, 1, 0, *-1*
			sum -= arr[i]
		}
		patIdx++
	}

	return aoc.Abs(sum) % 10
}
