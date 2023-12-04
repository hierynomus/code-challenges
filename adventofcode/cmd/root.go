package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Config struct {
	InputDir string
	Time     bool
}

func RootCommand(config *Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "solve",
		Short: "Run the solver",
	}

	cmd.PersistentFlags().StringVarP(&config.InputDir, "input", "i", "", "The input directory")
	cmd.PersistentFlags().BoolVarP(&config.Time, "time", "t", false, "Time solutions")

	return cmd
}

func Execute(ctx context.Context) {
	config := &Config{}
	cmd := RootCommand(config)

	for year := range Years {
		cmd.AddCommand(AocCommand(year, config))
	}

	if err := cmd.ExecuteContext(ctx); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
