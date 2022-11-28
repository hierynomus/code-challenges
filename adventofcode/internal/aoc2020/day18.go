package aoc2020

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day18(reader *bufio.Scanner) (string, string) {
	lines := aoc.ReadStringArray(reader)

	part1 := 0
	part2 := 0
	for _, sum := range lines {
		s := strings.Split(strings.ReplaceAll(strings.ReplaceAll(sum, ")", " )"), "(", "( "), " ")
		sol, _ := solve(s, calculate1)
		part1 += sol
		sol2, _ := solve(s, calculate2)
		part2 += sol2
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func solve(sum []string, solver func(sum []string) int) (int, int) {
	startsAt := 0
	i := 0
	for i < len(sum) {
		if sum[i] == "(" {
			startsAt = i
			i++
		} else if sum[i] == ")" {
			sol := solver(sum[startsAt+1 : i])
			newSum := []string{}
			newSum = append(newSum, sum[0:startsAt]...)
			newSum = append(newSum, strconv.Itoa(sol))
			newSum = append(newSum, sum[i+1:]...)
			sum = newSum
			i = 0
		} else {
			i++
		}
	}
	return solver(sum), 0
}

func calculate1(sum []string) int {
	// fmt.Printf("%v", sum)
	solution := aoc.ToInt(sum[0])
	i := 1
	for i < len(sum) {
		switch sum[i] {
		case "+":
			solution += aoc.ToInt(sum[i+1])
			i += 2
		case "*":
			solution *= aoc.ToInt(sum[i+1])
			i += 2
		}
	}
	return solution
}

func calculate2(sum []string) int {
	// fmt.Printf("%v", sum)
	ints := []int{aoc.ToInt(sum[0])}
	i := 1
	for i < len(sum) {
		switch sum[i] {
		case "+":
			ints[len(ints)-1] += aoc.ToInt(sum[i+1])
			i += 2
		case "*":
			ints = append(ints, aoc.ToInt(sum[i+1]))
			i += 2
		}
	}

	solution := 1
	for _, i := range ints {
		solution *= i
	}
	return solution
}
