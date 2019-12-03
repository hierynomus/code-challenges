package command

import (
	"bufio"
	"fmt"
	"os"

	"github.com/hierynomus/aoc2019/days"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(dayCommand(2, &days.Day02{}))
	rootCmd.AddCommand(dayCommand(3, &days.Day03{}))
}

func dayCommand(day int, s days.Solver) *cobra.Command {
	return &cobra.Command{
		Use:   fmt.Sprintf("day%02d", day),
		Short: fmt.Sprintf("Solve Day %02d", day),
		Run: func(cmf *cobra.Command, args []string) {
			reader := bufio.NewScanner(os.Stdin)
			p1, p2 := s.Solve(reader)
			fmt.Printf("Day %02d.1: %s\n", day, p1)
			fmt.Printf("Day %02d.2: %s\n", day, p2)
		},
	}
}

var rootCmd = &cobra.Command{
	Use:   "aoc2019",
	Short: "AoC 2019 Solutions",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
