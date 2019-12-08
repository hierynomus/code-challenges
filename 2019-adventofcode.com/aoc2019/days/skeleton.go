package days

import "bufio"

type Day0X struct{}

func (d *Day0X) Solve(scanner *bufio.Scanner) (string, string) {
	for scanner.Scan() {
		l := scanner.Text()
		println(l)
	}

	return "", ""
}
