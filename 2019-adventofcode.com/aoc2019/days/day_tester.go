package days

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

type DayTester struct {
	solver Solver
	t      *testing.T
}

func TestDay(s Solver, t *testing.T) *DayTester {
	return &DayTester{
		solver: s,
		t:      t,
	}
}

func (d *DayTester) WithFile(file string, expected1 string, expected2 string) {
	fileHandle, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)
	d.withScanner(fileScanner, expected1, expected2)
}

func (d *DayTester) WithInput(input string, expected1 string, expected2 string) {
	r := strings.NewReader(input)
	scanner := bufio.NewScanner(r)
	d.withScanner(scanner, expected1, expected2)
}

func (d *DayTester) withScanner(scanner *bufio.Scanner, expected1 string, expected2 string) {
	o1, o2 := d.solver.Solve(scanner)
	assert.Equal(d.t, o1, expected1)
	assert.Equal(d.t, o2, expected2)
}
