package aoc2020

import (
	"bufio"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type Cup struct {
	Nr   int
	Next *Cup
}

func Day23(reader *bufio.Scanner) (string, string) {
	l := aoc.Read(reader)
	m, current := parseCups(l)
	// _ = playCrabCupsArray(game, 100)
	playCrabCups(m, current, 100)

	part1 := []int{}
	c := m[1]
	for i := 0; i < 8; i++ {
		part1 = append(part1, c.Next.Nr)
		c = c.Next
	}

	m, current = parseCups(l)
	last := m[int(l[len(l)-1]-'0')]
	for i := 10; i <= 1000000; i++ {
		n := &Cup{Nr: i}
		m[i] = n
		last.Next = n
		last = n
	}

	last.Next = current

	playCrabCups(m, current, 10000000)
	part2 := m[1].Next.Nr * m[1].Next.Next.Nr
	return aoc.IntArrayAsString(part1, ""), strconv.Itoa(part2)
}

func parseCups(l string) (map[int]*Cup, *Cup) {
	var current, cup *Cup = nil, nil
	m := map[int]*Cup{}
	for _, x := range []rune(l) {
		c := &Cup{Nr: int(x - '0')}
		m[c.Nr] = c
		if current == nil {
			current = c
			cup = c
		} else if cup != nil {
			cup.Next = c
			cup = c
		}
	}

	cup.Next = current

	return m, current
}

func playCrabCups(lookup map[int]*Cup, cup *Cup, rounds int) {
	current := cup
	for i := 0; i < rounds; i++ {
		c1, c2, c3 := current.Next, current.Next.Next, current.Next.Next.Next
		current.Next = c3.Next
		dest := current.Nr - 1
		for c1.Nr == dest || c2.Nr == dest || c3.Nr == dest || dest <= 0 {
			dest--
			if dest <= 0 {
				dest = len(lookup)
			}
		}

		c3.Next = lookup[dest].Next
		lookup[dest].Next = c1

		current = current.Next
	}
}

func playCrabCupsArray(game []int, rounds int) []int {
	round := make([]int, len(game))
	copy(round, game)
	for i := 0; i < rounds; i++ {
		current := round[0]
		removed := round[1:4]
		find, foundIdx := current, 0
		for foundIdx == 0 {
			find--
			if find <= 0 {
				find = 9
			}
			for j := 4; j < len(game); j++ {
				if round[j] == find {
					foundIdx = j
					break
				}
			}
		}

		newRound := []int{}
		newRound = append(newRound, round[4:foundIdx+1]...)
		newRound = append(newRound, removed...)
		newRound = append(newRound, round[foundIdx+1:]...)
		newRound = append(newRound, round[0])
		round = newRound
	}

	return round
}
