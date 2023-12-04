package day

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/util"
	"github.com/stretchr/testify/assert"
)

type Tester struct {
	day Solver
	t   *testing.T
}

func RunDays(t *testing.T, days map[int]struct {
	S     Solver
	Part1 string
	Part2 string
}) {
	for dd, ss := range days {
		d := dd
		s := ss
		pkg := util.GetPackageName(s.S)
		t.Run(fmt.Sprintf("%s/Day%02d", pkg, d), func(t *testing.T) {
			day := TestDay(t, s.S)
			day.WithFile(fmt.Sprintf("../../input/%s/day%02d.in", pkg, d), s.Part1, s.Part2)
		})
	}
}

func TestDay(t *testing.T, s Solver) *Tester {
	return &Tester{
		day: s,
		t:   t,
	}
}

func (d *Tester) WithFile(file string, expected1 string, expected2 string) {
	fileHandle, err := os.Open(file)
	if err != nil {
		if os.IsNotExist(err) {
			d.t.Skipf("Skipping test, input file %s does not exist", file)
		} else {
			d.t.Fatalf("Error opening file %s: %v", file, err)
		}
		return
	}

	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)
	d.withScanner(fileScanner, expected1, expected2)
}

func (d *Tester) WithInput(input string, expected1 string, expected2 string) {
	r := strings.NewReader(input)
	scanner := bufio.NewScanner(r)
	d.withScanner(scanner, expected1, expected2)
}

func (d *Tester) withScanner(scanner *bufio.Scanner, expected1 string, expected2 string) {
	o1, o2 := d.day(scanner)
	assert.Equal(d.t, expected1, o1)
	assert.Equal(d.t, expected2, o2)
}
