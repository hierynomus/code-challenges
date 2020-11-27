package day

import (
	"bufio"
	"fmt"
	"os"
)

type Solver func(r *bufio.Scanner) (string, string)

func RunDayWithInput(day int, s Solver, f string) {
	fileHandle, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)
	RunDayWithScanner(day, s, fileScanner)
}

func RunDay(day int, s Solver) {
	RunDayWithScanner(day, s, bufio.NewScanner(os.Stdin))
}

func RunDayWithScanner(day int, s Solver, scanner *bufio.Scanner) {
	p1, p2 := s(scanner)
	f := bufio.NewWriter(os.Stdout)

	fmt.Fprintf(f, "Day %02d.1: %s\n", day, p1)
	fmt.Fprintf(f, "Day %02d.2: %s\n", day, p2)
	f.Flush()
}
