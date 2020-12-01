package day

import "bufio"

func Day0X(scanner *bufio.Scanner) (string, string) {
	for scanner.Scan() {
		l := scanner.Text()
		println(l)
	}

	return "", ""
}
