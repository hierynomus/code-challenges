package aoc2021

import (
	"bufio"
	"sort"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day10(reader *bufio.Scanner) (string, string) {
	var SyntaxScores = map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	var CounterPart = map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
		'>': '<',
	}
	var AutoCompleteScores = map[rune]int{
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}

	var part1, part2 int

	scores := make([]int, 0)
	for reader.Scan() {
		line := reader.Text()
		stack := make(aoc.RuneStack, 0)
		syntaxError := false
		for _, r := range line {
			switch r {
			case '(', '[', '{', '<':
				stack.Push(r)
			case ')', ']', '}', '>':
				if x, ok := stack.Pop(); !ok || x != CounterPart[r] {
					part1 += SyntaxScores[r]
					syntaxError = true
					break
				}
			}
		}

		if syntaxError {
			continue
		}

		score := 0
		for len(stack) > 0 {
			x, _ := stack.Pop()
			score = (score * 5) + AutoCompleteScores[x]
		}
		scores = append(scores, score)
	}

	sort.Sort(sort.IntSlice(scores))

	part2 = scores[len(scores)/2]

	return strconv.Itoa(part1), strconv.Itoa(part2)
}
