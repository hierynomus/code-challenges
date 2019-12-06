package days

import "bufio"

type Day07 struct{}

func (d *Day07) Solve(scanner bufio.Scanner) (string, string) {
	for scanner.Scan() {
		l := scanner.Text()
		println(l)
	}

	return "", ""
}
