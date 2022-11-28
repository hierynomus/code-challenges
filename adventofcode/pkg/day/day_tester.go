package day

import (
	"bufio"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Tester struct {
	day Solver
	t   *testing.T
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
