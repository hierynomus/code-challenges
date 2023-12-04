package cmd

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/hierynomus/code-challenges/adventofcode/internal/aoc2018"
	"github.com/hierynomus/code-challenges/adventofcode/internal/aoc2019"
	"github.com/hierynomus/code-challenges/adventofcode/internal/aoc2020"
	"github.com/hierynomus/code-challenges/adventofcode/internal/aoc2021"
	"github.com/hierynomus/code-challenges/adventofcode/internal/aoc2022"
	"github.com/hierynomus/code-challenges/adventofcode/internal/aoc2023"
	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
	"github.com/hierynomus/code-challenges/adventofcode/pkg/util"
	"github.com/spf13/cobra"
)

var Years = map[int]map[int]day.Solver{
	2018: aoc2018.AllDays,
	2019: aoc2019.AllDays,
	2020: aoc2020.AllDays,
	2021: aoc2021.AllDays,
	2022: aoc2022.AllDays,
	2023: aoc2023.AllDays,
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

				if config.Time {
					defer util.Timing(time.Now(), fmt.Sprintf("aoc%d", year))
				}

				for _, d := range keys {
					f := fmt.Sprintf("%s/aoc%d/day%02d.in", config.InputDir, year, d)
					day.RunDayWithInput(d, Years[year][d], f, config.Time)
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
				day.RunDayWithInput(d, s, f, config.Time)
			case strings.TrimSpace(config.InputDir) != "":
				f = fmt.Sprintf("%s/aoc%d/day%02d.in", config.InputDir, year, d)
				day.RunDayWithInput(d, s, f, config.Time)
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
