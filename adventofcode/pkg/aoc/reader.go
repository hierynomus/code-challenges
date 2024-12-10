package aoc

import (
	"bufio"
	"strings"
)

func ReadStringArray(reader *bufio.Scanner) []string {
	lines := []string{}

	for reader.Scan() {
		lines = append(lines, reader.Text())
	}

	return lines
}

func ReadIntArray(reader *bufio.Scanner) []int {
	lines := []int{}
	for reader.Scan() {
		lines = append(lines, ToInt(reader.Text()))
	}

	return lines
}

func ReadIntArrayLine(reader *bufio.Scanner, sep string) []int {
	if !reader.Scan() {
		panic("No input")
	}

	l := strings.Split(reader.Text(), sep)
	line := []int{}
	for _, x := range l {
		line = append(line, ToInt(x))
	}

	return line
}

func ReadIntGrid(reader *bufio.Scanner, sep string) [][]int {
	grid := [][]int{}
	for reader.Scan() {
		l := strings.Split(reader.Text(), sep)
		line := []int{}
		for _, x := range l {
			line = append(line, ToInt(x))
		}
		grid = append(grid, line)
	}

	return grid
}

func ReadRuneGrid(reader *bufio.Scanner) [][]rune {
	grid := [][]rune{}
	for reader.Scan() {
		l := []rune(reader.Text())
		grid = append(grid, l)
	}

	return grid
}

func Read(reader *bufio.Scanner) string {
	if !reader.Scan() {
		panic("No input")
	}

	return reader.Text()
}

func Skip(reader *bufio.Scanner) {
	if !reader.Scan() {
		panic("no input")
	}

	reader.Text()
}
