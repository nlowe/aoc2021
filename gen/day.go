package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/nlowe/aoc2021/util"
)

var (
	parts = [...]string{"A", "B"}
	funcs = template.FuncMap{
		"toLower": strings.ToLower,
		"toUpper": strings.ToUpper,
		"seq": func(start, end int) (result []int) {
			for i := start; i <= end; i++ {
				result = append(result, i)
			}
			return
		},
	}
)

type metadata struct {
	N  int
	AB string
}

func main() {
	n := dayLimit(time.Now())

	p, err := util.ChallengePath()
	if err != nil {
		abort(err)
	}

	fmt.Println("Regenerating ./challenge/cmd/cmd.go up to Day", n)

	_ = os.Remove(p)
	genFile(p, rootTemplate, funcs, metadata{N: n})
	goimports(p)

	for day := 1; day <= n; day++ {
		genDay(day)
	}
}

func goimports(file string) {
	goimports := exec.Command("goimports", "-local", "github.com/nlowe/aoc2021", "-w", file)
	if err := goimports.Run(); err != nil {
		abort(err)
	}
}

func dayLimit(now time.Time) int {
	now = now.UTC()

	if now.Year() > 2021 {
		return 25
	}

	if now.Month() != time.December {
		return 0
	}

	// Challenges unlock at 0500 UTC
	if now.Hour() < 5 {
		return util.IntClamp(0, now.Day()-1, 25)
	}

	return util.IntClamp(0, now.Day(), 25)
}

func genDay(n int) {
	p, err := util.PkgPath(n)
	if err != nil {
		abort(err)
	}

	if err := os.MkdirAll(p, 0744); err != nil {
		abort(err)
	}

	generated := false

	// Only try to generate code files if it looks like they're missing
	// Not all days last year were easily testable so we may not keep the _test.go
	// files around. This way we don't regenerate them if they get manually deleted.
	gluePath := filepath.Join(p, "import.go")
	if _, stat := os.Stat(gluePath); stat != nil && os.IsNotExist(stat) {
		generated = true
		genFile(gluePath, glueTemplate, funcs, metadata{N: n})

		for _, part := range parts {
			m := metadata{N: n, AB: part}
			genFile(filepath.Join(p, fmt.Sprintf("%s.go", strings.ToLower(part))), problemTemplate, funcs, m)
			genFile(filepath.Join(p, fmt.Sprintf("%s_test.go", strings.ToLower(part))), testTemplate, funcs, m)
		}

		genFile(filepath.Join(p, "benchmark_test.go"), benchmarkTemplate, funcs, metadata{N: n})
	}

	goimports(p)

	if _, stat := os.Stat(filepath.Join(p, "input.txt")); os.IsNotExist(stat) {
		generated = true
		fmt.Println("fetching input for day", n)
		problemInput, err := getInput(n)
		if err != nil {
			panic(err)
		}

		if err := os.WriteFile(filepath.Join(p, "input.txt"), problemInput, 0644); err != nil {
			panic(err)
		}
	}

	if generated {
		fmt.Printf("Generated problems for day %d\n", n)
	} else {
		fmt.Printf("Day %d up-to-date\n", n)
	}
}

func genFile(path, t string, funcs template.FuncMap, m metadata) {
	if _, stat := os.Stat(path); os.IsNotExist(stat) {
		fmt.Println("creating", path)
		t := template.Must(template.New(path).Funcs(funcs).Parse(t))
		cf, err := os.Create(path)
		if err != nil {
			abort(err)
		}

		defer mustClose(cf)
		if err := t.Execute(cf, m); err != nil {
			abort(err)
		}
	} else {
		fmt.Println(path, "already exists, skipping...")
	}
}

func getInput(day int) ([]byte, error) {
	// TODO: Revert back to kooky when https://github.com/zellyn/kooky/issues/49 is fixed
	token, set := os.LookupEnv("AOC_SESSION_TOKEN")
	if !set {
		return nil, fmt.Errorf("AOC_SESSION_TOKEN not set")
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://adventofcode.com/2021/day/%d/input", day), nil)
	if err != nil {
		return nil, err
	}

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: token,
	})

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer mustClose(resp.Body)

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code %s: %s", resp.Status, body)
	}

	body, err := io.ReadAll(resp.Body)
	return body, err
}

func mustClose(c io.Closer) {
	if c == nil {
		return
	}

	if err := c.Close(); err != nil {
		panic(fmt.Errorf("error closing io.Closer: %w", err))
	}
}

func abort(err error) {
	fmt.Printf("%s\n\nsyntax: go run gen/day.go <day> <a|b>\n", err.Error())
	os.Exit(1)
}
