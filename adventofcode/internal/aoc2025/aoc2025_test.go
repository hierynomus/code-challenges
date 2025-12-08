package aoc2025

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

func TestDays(t *testing.T) {
	tests := map[int]struct {
		S     day.Solver
		Part1 string
		Part2 string
	}{
		1: {S: Day01, Part1: "1182", Part2: "6907"},
		2: {S: Day02, Part1: "9188031749", Part2: "11323661261"},
		3: {S: Day03, Part1: "17383", Part2: "172601598658203"},
		4: {S: Day04, Part1: "1419", Part2: "8739"},
		5: {S: Day05, Part1: "623", Part2: "353507173555373"},
		6: {S: Day06, Part1: "4364617236318", Part2: "9077004354241"},
		7: {S: Day07, Part1: "1678", Part2: "0"},
		8: {S: Day08, Part1: "244188", Part2: "8361881885"},
		// 9:  {S: Day09, Part1: "6320029754031", Part2: "6347435485773"},
		// 10: {S: Day10, Part1: "682", Part2: "1511"},
		// 11: {S: Day11, Part1: "191690", Part2: "228651922369703"},
		// 12: {S: Day12, Part1: "1477762", Part2: "923480"}, // TODO part 2
	}

	day.RunDays(t, tests)
}
