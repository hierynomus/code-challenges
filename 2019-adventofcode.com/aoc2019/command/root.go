package command

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/hierynomus/aoc2019/days"
	"github.com/spf13/cobra"
)

var allDays map[int]days.Day = map[int]days.Day{ //nolint:gochecknoglobals
	1:  &days.Day01{},
	2:  &days.Day02{},
	3:  &days.Day03{},
	4:  &days.Day04{},
	5:  &days.Day05{},
	6:  &days.Day06{},
	7:  &days.Day07{},
	8:  &days.Day08{},
	9:  &days.Day09{},
	10: &days.Day10{},
	11: &days.Day11{},
	12: &days.Day12{},
	13: &days.Day13{},
	14: &days.Day14{},
	15: &days.Day15{},
	16: &days.Day16{},
	17: &days.Day17{},
	19: &days.Day19{},
}

var inputDir string //nolint:gochecknoglobals

func init() { //nolint:gochecknoinits
	for k, v := range allDays {
		rootCmd.AddCommand(dayCommand(k, v))
	}

	rootCmd.PersistentFlags().StringVarP(&inputDir, "input", "i", "", "Directory containing input files")
}

func dayCommand(day int, s days.Day) *cobra.Command {
	var f string

	cmd := &cobra.Command{
		Use:   fmt.Sprintf("day%02d", day),
		Short: fmt.Sprintf("Solve Day %02d", day),
		Run: func(cmf *cobra.Command, args []string) {
			switch {
			case strings.TrimSpace(f) != "":
				runDayWithInput(day, s, f)
			case strings.TrimSpace(inputDir) != "":
				f := fmt.Sprintf("%s/day%02d.in", inputDir, day)
				runDayWithInput(day, s, f)
			case stdInAvailable():
				runDay(day, s)
			default:
				panic(fmt.Errorf("no input available"))
			}
		},
	}

	cmd.Flags().StringVarP(&f, "file", "f", "", "Input file")

	return cmd
}

func stdInAvailable() bool {
	stat, _ := os.Stdin.Stat()
	return stat.Mode()&os.ModeCharDevice == 0
}

func runDayWithInput(day int, s days.Day, f string) {
	fileHandle, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)
	runDayWithScanner(day, s, fileScanner)
}

func runDay(day int, s days.Day) {
	runDayWithScanner(day, s, bufio.NewScanner(os.Stdin))
}

func runDayWithScanner(day int, s days.Day, scanner *bufio.Scanner) {
	p1, p2 := s.Solve(scanner)
	f := bufio.NewWriter(os.Stdout)

	fmt.Fprintf(f, "Day %02d.1: %s\n", day, p1)
	fmt.Fprintf(f, "Day %02d.2: %s\n", day, p2)
	f.Flush()
}

var rootCmd = &cobra.Command{ //nolint:gochecknoglobals
	Use:   "aoc2019",
	Short: "AoC 2019 Solutions",
	Run: func(cmd *cobra.Command, args []string) {
		if strings.TrimSpace(inputDir) != "" {
			var keys []int
			for k := range allDays {
				keys = append(keys, k)
			}
			sort.Ints(keys)

			for _, day := range keys {
				f := fmt.Sprintf("%s/day%02d.in", inputDir, day)
				runDayWithInput(day, allDays[day], f)
			}
		} else {
			panic(fmt.Errorf("cannot run all days without input dir"))
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
