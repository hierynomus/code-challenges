package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/internal/aoc2018"
	"github.com/hierynomus/code-challenges/adventofcode/internal/aoc2019"
	"github.com/hierynomus/code-challenges/adventofcode/internal/aoc2020"
	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
	"github.com/hierynomus/code-challenges/adventofcode/pkg/util"
	"github.com/spf13/cobra"
)

var Years = map[int]map[int]day.Solver{
	2018: {
		1: aoc2018.Day01,
		2: aoc2018.Day02,
		3: aoc2018.Day03,
		4: aoc2018.Day04,
		5: aoc2018.Day05,
	},
	2019: {
		1:  aoc2019.Day01,
		2:  aoc2019.Day02,
		3:  aoc2019.Day03,
		4:  aoc2019.Day04,
		5:  aoc2019.Day05,
		6:  aoc2019.Day06,
		7:  aoc2019.Day07,
		8:  aoc2019.Day08,
		9:  aoc2019.Day09,
		10: aoc2019.Day10,
		11: aoc2019.Day11,
		12: aoc2019.Day12,
		13: aoc2019.Day13,
		14: aoc2019.Day14,
		15: aoc2019.Day15,
		16: aoc2019.Day16,
		17: aoc2019.Day17,
		// 18: aoc2019.Day18,
		19: aoc2019.Day19,
		// 20: aoc2019.Day20,
		// 21: aoc2019.Day21,
		22: aoc2019.Day22,
		// 23: aoc2019.Day23,
		24: aoc2019.Day24,
		// 25: aoc2019.Day25,
	},
	2020: {
		1: aoc2020.Day01,
		// 2:  aoc2020.Day02,
		// 3:  aoc2020.Day03,
		// 4:  aoc2020.Day04,
		// 5:  aoc2020.Day05,
		// 6:  aoc2020.Day06,
		// 7:  aoc2020.Day07,
		// 8:  aoc2020.Day08,
		// 9:  aoc2020.Day09,
		// 10: aoc2020.Day10,
		// 11: aoc2020.Day11,
		// 12: aoc2020.Day12,
		// 13: aoc2020.Day13,
		// 14: aoc2020.Day14,
		// 15: aoc2020.Day15,
		// 16: aoc2020.Day16,
		// 17: aoc2020.Day17,
		// // 18: aoc2020.Day18,
		// 19: aoc2020.Day19,
		// // 20: aoc2020.Day20,
		// // 21: aoc2020.Day21,
		// 22: aoc2020.Day22,
		// // 23: aoc2020.Day23,
		// 24: aoc2020.Day24,
		// // 25: aoc2020.Day25,
	},
}

func AocCommand(year int, config *Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   fmt.Sprintf("aoc%d", year),
		Short: fmt.Sprintf("AoC Solutions %d", year),
		Run: func(cmf *cobra.Command, args []string) {
			if config.InputDir != "" {
				var keys []int
				for k := range Years[year] {
					keys = append(keys, k)
				}
				sort.Ints(keys)

				for _, d := range keys {
					f := fmt.Sprintf("%s/aoc%d/day%02d.in", config.InputDir, year, d)
					day.RunDayWithInput(d, Years[year][d], f)
				}
			} else {
				panic(fmt.Errorf("cannot run all days without input dir"))
			}
		},
	}

	for d, s := range Years[year] {
		cmd.AddCommand(DayCommand(year, d, s, config))
	}

	return cmd
}

func DayCommand(year int, d int, s day.Solver, config *Config) *cobra.Command {
	var f string
	cmd := &cobra.Command{
		Use:   fmt.Sprintf("day%02d", d),
		Short: fmt.Sprintf("Solve day %d.%02d", year, d),
		Run: func(cmd *cobra.Command, args []string) {
			switch {
			case strings.TrimSpace(f) != "":
				day.RunDayWithInput(d, s, f)
			case strings.TrimSpace(config.InputDir) != "":
				f := fmt.Sprintf("%s/aoc%d/day%02d.in", config.InputDir, year, d)
				day.RunDayWithInput(d, s, f)
			case util.StdInAvailable():
				day.RunDay(d, s)
			default:
				panic(fmt.Errorf("no input available"))
			}

		},
	}

	cmd.Flags().StringVarP(&f, "file", "f", "", "Input file")

	return cmd
}
