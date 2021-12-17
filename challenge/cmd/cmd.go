// Code generated by 'go run ./gen'; DO NOT EDIT

package cmd

import (
	"fmt"
	"time"

	"github.com/pkg/profile"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/nlowe/aoc2021/challenge/day1"
	"github.com/nlowe/aoc2021/challenge/day10"
	"github.com/nlowe/aoc2021/challenge/day11"
	"github.com/nlowe/aoc2021/challenge/day12"
	"github.com/nlowe/aoc2021/challenge/day13"
	"github.com/nlowe/aoc2021/challenge/day14"
	"github.com/nlowe/aoc2021/challenge/day15"
	"github.com/nlowe/aoc2021/challenge/day16"
	"github.com/nlowe/aoc2021/challenge/day17"
	"github.com/nlowe/aoc2021/challenge/day2"
	"github.com/nlowe/aoc2021/challenge/day3"
	"github.com/nlowe/aoc2021/challenge/day4"
	"github.com/nlowe/aoc2021/challenge/day5"
	"github.com/nlowe/aoc2021/challenge/day6"
	"github.com/nlowe/aoc2021/challenge/day7"
	"github.com/nlowe/aoc2021/challenge/day8"
	"github.com/nlowe/aoc2021/challenge/day9"
	"github.com/nlowe/aoc2021/challenge/example"
)

func addDays(root *cobra.Command) {
	example.AddCommandsTo(root)
	day1.AddCommandsTo(root)
	day2.AddCommandsTo(root)
	day3.AddCommandsTo(root)
	day4.AddCommandsTo(root)
	day5.AddCommandsTo(root)
	day6.AddCommandsTo(root)
	day7.AddCommandsTo(root)
	day8.AddCommandsTo(root)
	day9.AddCommandsTo(root)
	day10.AddCommandsTo(root)
	day11.AddCommandsTo(root)
	day12.AddCommandsTo(root)
	day13.AddCommandsTo(root)
	day14.AddCommandsTo(root)
	day15.AddCommandsTo(root)
	day16.AddCommandsTo(root)
	day17.AddCommandsTo(root)
}

type prof interface {
	Stop()
}

func NewRootCommand() *cobra.Command {
	var (
		start    time.Time
		profiler prof
	)

	result := &cobra.Command{
		Use:     "aoc2021",
		Short:   "Advent of Code 2021 Solutions",
		Long:    "Golang implementations for the 2021 Advent of Code problems",
		Example: "go run main.go 1 a -i ./challenge/day1/input.txt",
		Args:    cobra.ExactArgs(1),
		PersistentPreRun: func(_ *cobra.Command, _ []string) {
			if viper.GetBool("profile") {
				profiler = profile.Start()
			}

			start = time.Now()
		},
		PersistentPostRun: func(_ *cobra.Command, _ []string) {
			if profiler != nil {
				profiler.Stop()
			}

			fmt.Println("Took", time.Since(start))
		},
	}

	addDays(result)

	flags := result.PersistentFlags()

	flags.StringP("input", "i", "", "Input File to read. If not specified, assumes ./challenge/dayN/input.txt for the currently running challenge")
	flags.Bool("profile", false, "Profile implementation performance")

	_ = viper.BindPFlags(flags)

	return result
}
