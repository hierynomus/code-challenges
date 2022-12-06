package aoc2022

import (
	"bufio"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type CargoMove struct {
	nr, from, to int
}

func ReadStacks(reader *bufio.Scanner) map[int]aoc.RuneStack {
	stacks := map[int]aoc.RuneStack{}
	for {
		l := aoc.Read(reader)
		if l == "" {
			return stacks
		}
		for i := 0; 1+(i*4) < len(l); i++ {
			idx := 1 + (i * 4)
			if l[idx] != ' ' {
				if _, ok := stacks[i+1]; !ok {
					stacks[i+1] = aoc.RuneStack{}
				}
				// Prepend to build up Stack correctly (we read top-to-bottom)
				stacks[i+1] = append([]rune{rune(l[idx])}, stacks[i+1]...)
			}
		}
	}
}

func ReadMoves(reader *bufio.Scanner) []*CargoMove {
	moves := []*CargoMove{}
	for _, m := range aoc.ReadStringArray(reader) {
		parts := strings.Split(m, " ")
		cm := &CargoMove{nr: aoc.ToInt(parts[1]), from: aoc.ToInt(parts[3]), to: aoc.ToInt(parts[5])}
		moves = append(moves, cm)
	}

	return moves
}

func MoveCargo9000(stacks []*aoc.RuneStack, moves []*CargoMove) {
	for _, move := range moves {
		for i := 0; i < move.nr; i++ {
			item, ok := stacks[move.from-1].Pop()
			if !ok {
				panic("Stack empty")
			}
			stacks[move.to-1].Push(item)
		}
	}
}

func MoveCargo9001(stacks []*aoc.RuneStack, moves []*CargoMove) {
	for _, move := range moves {
		inAir := &aoc.RuneStack{}
		for i := 0; i < move.nr; i++ {
			item, ok := stacks[move.from-1].Pop()
			if !ok {
				panic("Stack empty")
			}
			inAir.Push(item)
		}

		for !inAir.IsEmpty() {
			i, _ := inAir.Pop()
			stacks[move.to-1].Push(i)
		}
	}
}

func Day05(reader *bufio.Scanner) (string, string) {
	stacks := ReadStacks(reader)
	moves := ReadMoves(reader)

	tmpStacks := make([]*aoc.RuneStack, len(stacks))
	for i, s := range stacks {
		ss := make(aoc.RuneStack, len(s))
		copy(ss, s)
		tmpStacks[i-1] = &ss
	}

	MoveCargo9000(tmpStacks, moves)

	part1 := []rune{}
	for _, s := range tmpStacks {
		r, _ := s.Peek()
		part1 = append(part1, r)
	}

	tmpStacks = make([]*aoc.RuneStack, len(stacks))
	for i, s := range stacks {
		ss := make(aoc.RuneStack, len(s))
		copy(ss, s)
		tmpStacks[i-1] = &ss
	}

	MoveCargo9001(tmpStacks, moves)

	part2 := []rune{}
	for _, s := range tmpStacks {
		r, _ := s.Peek()
		part2 = append(part2, r)
	}

	return string(part1), string(part2)
}
