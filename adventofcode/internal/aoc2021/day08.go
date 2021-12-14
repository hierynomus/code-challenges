package aoc2021

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

//   0:      1:      2:      3:      4:
//  aaaa    ....    aaaa    aaaa    ....
// b    c  .    c  .    c  .    c  b    c
// b    c  .    c  .    c  .    c  b    c
//  ....    ....    dddd    dddd    dddd
// e    f  .    f  e    .  .    f  .    f
// e    f  .    f  e    .  .    f  .    f
//  gggg    ....    gggg    gggg    ....

//   5:      6:      7:      8:      9:
//  aaaa    aaaa    aaaa    aaaa    aaaa
// b    .  b    .  .    c  b    c  b    c
// b    .  b    .  .    c  b    c  b    c
//  dddd    dddd    ....    dddd    dddd
// .    f  e    f  .    f  e    f  .    f
// .    f  e    f  .    f  e    f  .    f
//  gggg    gggg    ....    gggg    gggg

// Occurrences
// a = 8 // done
// b = 6 // done
// c = 8 // done
// d = 7
// e = 4 // done
// f = 9 // done
// g = 7

type SevSeg map[rune]rune

var SevSegNumber = map[string]int{
	"cf":      1,
	"acdeg":   2,
	"acdfg":   3,
	"bcdf":    4,
	"abdfg":   5,
	"abdefg":  6,
	"acf":     7,
	"abcdefg": 8,
	"abcdfg":  9,
	"abcefg":  0,
}

func (ss SevSeg) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	for from, to := range ss {
		sb.WriteRune(from)
		sb.WriteString(" -> ")
		sb.WriteRune(to)
		sb.WriteString(", ")
	}
	sb.WriteString("]\n")
	return sb.String()
}

func (ss SevSeg) Map(from, to rune) {
	ss[from] = to
}

func (ss SevSeg) RenderNumber(s string) string {
	col := func(r1, r2 rune) string {
		out := ""
		for i := 0; i < 2; i++ {
			if strings.ContainsRune(s, ss[r1]) {
				out += "X  "
			} else {
				out += ".  "
			}
			if strings.ContainsRune(s, ss[r2]) {
				out += "  X\n"
			} else {
				out += "  .\n"
			}
		}

		return out
	}

	line := func(r rune) string {
		if strings.ContainsRune(s, ss[r]) {
			return " XXXX \n"
		} else {
			return " .... \n"
		}
	}

	out := ""
	out += line('a')
	out += col('b', 'c')
	out += line('d')
	out += col('e', 'f')
	out += line('g')

	return out
}

func (ss SevSeg) MapOccurrenceCount(m map[int]rune, inp ...string) {
	hist := aoc.MakeRuneHistogram([]rune(strings.Join(inp, "")))
	for k, v := range hist {
		if from, ok := m[int(v)]; ok {
			ss.Map(from, k)
		}
	}
}

func (ss SevSeg) StripKnowns(inp string) string {
	for _, v := range ss {
		inp = strings.ReplaceAll(inp, string(v), "")
	}
	return inp
}

func (ss SevSeg) AsNumber(s string) int {
	translated := ""
	for from, to := range ss {
		if strings.ContainsRune(s, to) {
			translated += string(from)
		}
	}

	return SevSegNumber[aoc.SortStringByCharacter(translated)]
}

func Day08(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	for reader.Scan() {
		line := strings.Split(reader.Text(), " | ")
		inp, out := line[0], line[1]
		ss := Demangle(inp)

		for i, v := range strings.Split(out, " ") {
			l := len(v)
			if l == 2 || l == 3 || l == 4 || l == 7 {
				part1++
			}

			nr := ss.AsNumber(v)
			for j := 3 - i; j > 0; j-- {
				nr *= 10
			}
			part2 += nr
		}

	}
	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func Demangle(inp string) SevSeg {
	var one, four, seven, eight string
	for _, v := range strings.Split(inp, " ") {
		switch len(v) {
		case 2:
			one = v
		case 3:
			seven = v
		case 4:
			four = v
		case 7:
			eight = v
		}
	}

	inp2 := strings.ReplaceAll(inp, " ", "")

	sevSeg := make(SevSeg)
	// Basic knowledge about occurrences of segments in the input
	sevSeg.MapOccurrenceCount(map[int]rune{4: 'e', 6: 'b', 9: 'f'}, inp2)
	// Between 1 and 7 only segment 'a' occurs once
	sevSeg.MapOccurrenceCount(map[int]rune{1: 'a'}, one, seven)
	// Because we know 'a', 'c' is now known
	sevSeg.MapOccurrenceCount(map[int]rune{8: 'c'}, sevSeg.StripKnowns(inp2))
	// Only 'd' and 'g' need to be mapped, using four and eight with knowledge of all knowns
	sevSeg.MapOccurrenceCount(map[int]rune{2: 'd', 1: 'g'}, sevSeg.StripKnowns(four), sevSeg.StripKnowns(eight))

	return sevSeg
}
