package days

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/hierynomus/aoc2019/aoc"
)

type Day04 struct{}

func (d *Day04) Solve(scanner *bufio.Scanner) (string, string) {
	if !scanner.Scan() {
		panic(fmt.Errorf("could not read"))
	}

	r := strings.Split(scanner.Text(), "-")
	lo, hi := aoc.ToInt(r[0]), aoc.ToInt(r[1])

	nrPasswords := 0
	nrPasswords2 := 0

	for i := lo; i <= hi; i++ {
		if isAsc(i) && hasPair(i) {
			nrPasswords++

			if hasExactPair(i) {
				nrPasswords2++
			}
		}
	}

	return strconv.Itoa(nrPasswords), strconv.Itoa(nrPasswords2)
}

func isAsc(i int) bool {
	x := i
	c := x % 10

	for x > 0 {
		x /= 10
		nc := x % 10

		if c < nc {
			return false
		}

		c = nc
	}

	return true
}

func hasPair(i int) bool {
	x := i
	c := x % 10

	for x > 0 {
		x /= 10
		nc := x % 10

		if c == nc {
			return true
		}

		c = nc
	}

	return false
}

func hasExactPair(i int) bool {
	x := i
	c := x % 10
	count := 1

	for x > 0 {
		x /= 10
		nc := x % 10

		if c == nc {
			count++
		}

		if c != nc && count == 2 {
			return true
		} else if c != nc {
			count = 1
		}

		c = nc
	}

	return false
}
