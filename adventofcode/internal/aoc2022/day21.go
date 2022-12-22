package aoc2022

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type Monkey struct {
	Name        string
	Left, Right *Monkey
	Operator    string
	Val         int
}

func (m *Monkey) Shout() int {
	if m.Operator == "" {
		return m.Val
	}

	switch m.Operator {
	case "+":
		return m.Left.Shout() + m.Right.Shout()
	case "*":
		return m.Left.Shout() * m.Right.Shout()
	case "-":
		return m.Left.Shout() - m.Right.Shout()
	case "/":
		return m.Left.Shout() / m.Right.Shout()
	}

	panic("unknown operator")
}

func GetOrCreateMonkey(monkeys map[string]*Monkey, name string) *Monkey {
	if _, ok := monkeys[name]; !ok {
		monkeys[name] = &Monkey{Name: name}
	}

	return monkeys[name]
}

func Day21(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	monkeys := map[string]*Monkey{}

	for reader.Scan() {
		line := reader.Text()
		s := strings.Split(line, ": ")
		rs := strings.Split(s[1], " ")
		m := GetOrCreateMonkey(monkeys, s[0])
		if len(rs) == 1 {
			m.Val = aoc.ToInt(rs[0])
		} else {
			m.Operator = rs[1]
			m.Left = GetOrCreateMonkey(monkeys, rs[0])
			m.Right = GetOrCreateMonkey(monkeys, rs[2])
		}
	}

	root := monkeys["root"]

	part1 = root.Shout()

	return strconv.Itoa(part1), strconv.Itoa(part2)
}
