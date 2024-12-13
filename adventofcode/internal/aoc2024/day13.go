package aoc2024

import (
	"bufio"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type ClawMachine struct {
	Ax, Ay, Bx, By int64
	Px, Py         int64
}

func (m ClawMachine) Solve() (int64, int64) {
	// (A*Ax + B*Bx) = Px
	// (A*Ya + B*Yb) = Py
	// B = (Px - A*Ax) / Bx
	// B = (Py - A*Ya) / Yb
	// (Px - A*Ax) / Bx = (Py - A*Ya) / Yb
	// (Px - A*Ax) * Yb = (Py - A*Ya) * Bx
	// Px*Yb - A*Ax*Yb = Py*Bx - A*Ya*Bx
	// A*Ya*Bx - A*Ax*Yb = Py*Bx - Px*Yb
	// A(Ya*Bx - Ax*Yb) = Py*Bx - Px*Yb
	// A = (Py*Bx - Px*Yb) / (Ya*Bx - xa*Yb)

	// Check for remainder
	if (m.Py*m.Bx-m.Px*m.By)%(m.Ay*m.Bx-m.Ax*m.By) != 0 {
		return -1, -1
	}

	a := (m.Py*m.Bx - m.Px*m.By) / (m.Ay*m.Bx - m.Ax*m.By)

	if (m.Px-a*m.Ax)%m.Bx != 0 {
		return -1, -1
	}

	b := (m.Px - a*m.Ax) / m.Bx

	return a, b
}

func Day13(reader *bufio.Scanner) (string, string) {
	var part1, part2 int64

	machines := []ClawMachine{}
	curr := ClawMachine{}
	for reader.Scan() {
		line := reader.Text()
		if line == "" {
			machines = append(machines, curr)
			curr = ClawMachine{}
			continue
		}

		splt := strings.Split(line, " ")

		if splt[1] == "A:" || splt[1] == "B:" {
			x := aoc.ToInt(splt[2][2 : len(splt[2])-1])
			y := aoc.ToInt(splt[3][2:])
			if splt[1] == "A:" {
				curr.Ax, curr.Ay = int64(x), int64(y)
			} else {
				curr.Bx, curr.By = int64(x), int64(y)
			}
		} else {
			x := aoc.ToInt(splt[1][2 : len(splt[1])-1])
			y := aoc.ToInt(splt[2][2:])
			curr.Px, curr.Py = int64(x), int64(y)
		}
	}
	machines = append(machines, curr)

	for _, m := range machines {
		a, b := m.Solve()

		if a > 0 && b > 0 {
			part1 += a*3 + b
		}

		m.Px, m.Py = m.Px+10000000000000, m.Py+10000000000000
		a, b = m.Solve()
		if a > 0 && b > 0 {
			part2 += a*3 + b
		}
	}

	return aoc.Int64ToString(part1), aoc.Int64ToString(part2)
}
