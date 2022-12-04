package challenge

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/nlowe/aoc2021/util"

	"github.com/spf13/viper"
)

type Input struct {
	scanner *bufio.Scanner

	lines chan string
}

func FromFile() *Input {
	path := viper.GetString("input")
	if path == "" {
		_, f, _, ok := runtime.Caller(1)
		if !ok {
			panic("failed to determine input path, provide it with -i instead")
		}

		path = filepath.Join(filepath.Dir(f), "input.txt")
	}

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	return newInputFromReader(f, f)
}

func FromLiteral(input string) *Input {
	return newInputFromReader(strings.NewReader(input), nil)
}

func newInputFromReader(r io.Reader, c io.Closer) *Input {
	result := &Input{
		scanner: bufio.NewScanner(r),
		lines:   make(chan string),
	}

	go func() {
		defer func() {
			if c != nil {
				_ = c.Close()
			}
		}()

		for result.scanner.Scan() {
			result.lines <- result.scanner.Text()
		}

		close(result.lines)
	}()

	return result
}

func (c *Input) Lines() <-chan string {
	return c.lines
}

func (c *Input) Ints() <-chan int {
	result := make(chan int)

	go func() {
		defer close(result)

		for line := range c.lines {
			result <- util.MustAtoI(line)
		}
	}()

	return result
}

func (c *Input) LineSlice() (result []string) {
	for line := range c.Lines() {
		result = append(result, line)
	}

	return
}

func (c *Input) IntSlice() (result []int) {
	for line := range c.Lines() {
		result = append(result, util.MustAtoI(line))
	}

	return
}
