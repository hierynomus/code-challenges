package days

import "bufio"

type Day interface {
	Solve(r *bufio.Scanner) (string, string)
}
