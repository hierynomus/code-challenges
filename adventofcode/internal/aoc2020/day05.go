package aoc2020

import (
	"bufio"
	"sort"
	"strconv"
)

type Seat struct {
	Row, Column int
}

func (s *Seat) ID() int {
	return s.Row*8 + s.Column
}

func Day05(reader *bufio.Scanner) (string, string) {
	seats := []*Seat{}

	for reader.Scan() {
		l := reader.Text()
		row, err := strconv.ParseInt(toBin(l[0:7], 'B', 'F'), 2, 32)
		if err != nil {
			panic(err)
		}
		col, err := strconv.ParseInt(toBin(l[7:], 'R', 'L'), 2, 32)

		seats = append(seats, &Seat{Row: int(row), Column: int(col)})
	}

	part1 := seats[0].ID()
	ids := make([]int, len(seats))
	for i, s := range seats {
		id := s.ID()
		ids[i] = id
		if id > part1 {
			part1 = id
		}
	}

	part2 := 0

	sort.Ints(ids)
	for i, id := range ids {
		if i == 0 {
			continue
		}

		if ids[i-1] != id-1 && ids[i+1] == id+1 {
			part2 = id - 1
			break
		}

		if ids[i-1] == id-1 && ids[i+1] != id+1 {
			part2 = id + 1
			break
		}
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func toBin(s string, one, zero rune) string {
	o := []rune{}
	for _, c := range s {
		if c == one {
			o = append(o, '1')
		} else if c == zero {
			o = append(o, '0')
		} else {
			panic("unrecognized char")
		}
	}
	return string(o)
}
