package aoc2020

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day16(reader *bufio.Scanner) (string, string) {
	rules := map[string]aoc.IntSet{}
	allValidValues := aoc.IntSet{}
	validTickets := [][]int{}

	for reader.Scan() {
		l := reader.Text()
		if len(l) == 0 {
			// Done with rules
			break
		}
		fv := strings.Split(l, ": ")
		ranges := strings.Split(fv[1], " or ")

		rule := aoc.IntSet{}
		for _, r := range ranges {
			mm := strings.Split(r, "-")
			rr := aoc.Range(aoc.ToInt(mm[0]), aoc.ToInt(mm[1])+1)
			allValidValues.Adds(rr)
			rule.Adds(rr)
		}

		rules[fv[0]] = rule
	}

	if aoc.Read(reader) != "your ticket:" {
		panic("Expected ticket")
	}
	myTicket := aoc.AsIntArray(aoc.Read(reader))

	aoc.Read(reader) // Empty line
	if aoc.Read(reader) != "nearby tickets:" {
		panic("Expected other tickets")
	}

	part1 := 0
	for reader.Scan() {
		ticket := aoc.AsIntArray(reader.Text())
		valid := true
		for _, v := range ticket {
			if !allValidValues.Contains(v) {
				part1 += v
				valid = false
			}
		}
		if valid {
			validTickets = append(validTickets, ticket)
		}
	}

	fieldMapping := discoverFieldMapping(rules, validTickets)

	part2 := 1
	for f, i := range fieldMapping {
		if strings.HasPrefix(f, "departure") {
			part2 *= myTicket[i]
		}
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func discoverFieldMapping(rules map[string]aoc.IntSet, tickets [][]int) map[string]int {
	mapping := map[string]int{}
	reverseMapping := map[int]string{}

	tt := aoc.Transpose(tickets)
	for len(mapping) < len(rules) {
		for idx, t := range tt {
			if _, ok := reverseMapping[idx]; ok {
				continue // Already know idx
			}

			possibleRules := []string{}
			for field, rule := range rules {
				if _, ok := mapping[field]; ok {
					continue // Already know field
				}

				allValid := true
				for _, v := range t {
					allValid = allValid && rule.Contains(v)
				}

				if allValid {
					possibleRules = append(possibleRules, field)
				}
			}

			if len(possibleRules) == 1 {
				mapping[possibleRules[0]] = idx
				reverseMapping[idx] = possibleRules[0]
			}
		}
	}

	return mapping
}
