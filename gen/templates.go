package main

const (
	rootTemplate = `// Code generated by 'go run ./gen'; DO NOT EDIT
package cmd

import (
	"fmt"
	"time"

	"github.com/pkg/profile"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
{{- if gt .N 0 }}
{{ range $day := seq 1 .N }}
	"github.com/nlowe/aoc2021/challenge/day{{ $day }}"
{{- end }}
{{- end}}
	"github.com/nlowe/aoc2021/challenge/example"
)

func addDays(root *cobra.Command) {
	example.AddCommandsTo(root)
    {{- range $day := seq 1 .N }}
	day{{ $day }}.AddCommandsTo(root)
    {{- end }}
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
`

	glueTemplate = `package day{{ .N }}

import "github.com/spf13/cobra"

func AddCommandsTo(root *cobra.Command) {
	day := &cobra.Command{
		Use:   "{{ .N }}",
		Short: "Problems for Day {{ .N }}",
	}

	day.AddCommand(aCommand())
    day.AddCommand(bCommand())

	root.AddCommand(day)
}
`

	problemTemplate = `package day{{ .N }}

import (
    "fmt"

    "github.com/nlowe/aoc2021/challenge"
	"github.com/nlowe/aoc2021/util"
	"github.com/spf13/cobra"
)

func {{ .AB | toLower }}Command() *cobra.Command {
    return &cobra.Command{
        Use:   "{{ .AB | toLower }}",
        Short: "Day {{ .N }}, Problem {{ .AB }}",
        Run: func(_ *cobra.Command, _ []string) {
            fmt.Printf("Answer: %d\n", part{{ .AB | toUpper }}(challenge.FromFile()))
        },
    }
}

func part{{ .AB | toUpper }}(challenge *challenge.Input) int {
    panic("Not implemented!")
}
`

	testTemplate = `package day{{ .N }}

import (
	"testing"

	"github.com/nlowe/aoc2021/challenge"
	"github.com/stretchr/testify/require"
)

func Test{{ .AB }}(t *testing.T) {
	input := challenge.FromLiteral("foobar")

	result := part{{ .AB | toUpper }}(input)

	require.Equal(t, 42, result)
}
`

	benchmarkTemplate = `package day{{ .N }}

import (
	"testing"

	"github.com/nlowe/aoc2021/challenge"
)

func BenchmarkA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = partA(challenge.FromFile())
	}
}

func BenchmarkB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = partB(challenge.FromFile())
	}
}`
)
