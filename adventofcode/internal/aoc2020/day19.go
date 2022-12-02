package aoc2020

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

var ErrNoMatch = errors.New("no match")

type Rule interface {
	match(s string, r []int, ruleset map[int]Rule) bool
}

var rules map[int]Rule = map[int]Rule{}
var rules2 map[int]Rule = map[int]Rule{}

type OrRule struct {
	opts [][]int
	i    int
}

type RuneRule struct {
	i int
	r rune
}

func NewOr(i int, s string) *OrRule {
	lr := strings.Split(s, " | ")
	opts := [][]int{}
	for _, r := range lr {
		opts = append(opts, aoc.AsIntArrayS(r, " "))
	}
	return &OrRule{i: i, opts: opts}
}

func NewRune(i int, s string) *RuneRule {
	return &RuneRule{
		i: i,
		r: rune(s[1]),
	}
}

// func (rule *RuneRule) match(s string, r []int, ruleset map[int]Rule) (int, error) {
// 	// fmt.Printf("%s => %s %d", r, s, idx)
// 	if idx < len(s) && rune(s[idx]) == r.r {
// 		// fmt.Println(" MATCH")
// 		return idx + 1, nil
// 	}
// 	// fmt.Println(" NO MATCH")
// 	return 0, NoMatch
// }

func (rule *RuneRule) match(s string, r []int, ruleset map[int]Rule) bool {
	if len(s) == 0 {
		return false
	}

	if rune(s[0]) == rule.r {
		if len(r) > 0 {
			return ruleset[r[0]].match(s[1:], r[1:], ruleset)
		} else {
			return len(s) == 1 // Last character and no rules left
		}
	}

	return false
}

func (rule *RuneRule) String() string {
	return fmt.Sprintf("RuneRule[%d]: %c", rule.i, rule.r)
}

func (rule *OrRule) match(s string, r []int, ruleset map[int]Rule) bool {
	found := false
	for _, opt := range rule.opts {
		newR := []int{}
		newR = append(newR, opt...)
		if len(r) > 0 {
			newR = append(newR, r...)
		}
		found = found || ruleset[newR[0]].match(s, newR[1:], ruleset)
	}

	return found
}

// func (rule *OrRule) match(s string, r []int, ruleset map[int]Rule) (int, error) {
// 	// fmt.Printf("%s => %s %d\n", r, s, idx)
// 	for _, opt := range rule.opts {
// 		i := idx
// 		var err error
// 		for _, x := range opt {
// 			i, err = ruleset[x].match(s, i, ruleset)
// 			if err != nil {
// 				break // Try next option
// 			}
// 		}
// 		if err == nil {
// 			return i, nil
// 		}
// 	}

// 	return 0, NoMatch
// }

func (rule *OrRule) String() string {
	s := aoc.IntArrayAsString(rule.opts[0], " ")
	for i := 1; i < len(rule.opts); i++ {
		s = fmt.Sprintf("%s | %s", s, aoc.IntArrayAsString(rule.opts[i], " "))
	}

	return fmt.Sprintf("OrRule[%d]: %s", rule.i, s)
}

func Day19(reader *bufio.Scanner) (string, string) {
	for reader.Scan() {
		l := reader.Text()
		if len(l) == 0 {
			break
		}
		kv := strings.Split(l, ": ")
		nr := aoc.ToInt(kv[0])
		if strings.Contains(kv[1], "\"") {
			rules[nr] = NewRune(nr, kv[1])
			rules2[nr] = NewRune(nr, kv[1])
		} else {
			rules[nr] = NewOr(nr, kv[1])
			rules2[nr] = NewOr(nr, kv[1])
		}
	}

	rules2[8] = NewOr(8, "42 | 42 8")
	rules2[11] = NewOr(11, "42 31 | 42 11 31")

	part1 := 0
	part2 := 0
	for reader.Scan() {
		l := reader.Text()

		if rules[0].match(l, []int{}, rules) {
			part1++
		} else if rules2[0].match(l, []int{}, rules2) {
			part2++
		}
		// 	if j, err := rules[0].match(l, []int{0}, rules); err == nil {
		// 		if j == len(l) {
		// 			println(l)
		// 			part1++
		// 			continue
		// 		}
		// 	}
		// 	// println("Attempt 2")
		// 	if j, err := rules2[0].match(l, []int{0}, rules2); err == nil {
		// 		if j == len(l) {
		// 			println(l)
		// 			part2++
		// 		}
		// 	}
		// }
	}

	return strconv.Itoa(part1), strconv.Itoa(part1 + part2)
}

// def test(s,seq):
//     if s == '' or seq == []:
//         return s == '' and seq == [] # if both are empty, True. If only one, False.

//     r = rules[seq[0]]
//     if '"' in r:
//         if s[0] in r:
//             return test(s[1:], seq[1:]) # strip first character
//         else:
//             return False # wrong first character
//     else:
//         return any(test(s, t + seq[1:]) for t in r) # expand first term
