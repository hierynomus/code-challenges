package days

import "bufio"

type Solver interface {
	Solve(r *bufio.Scanner) (string, string)
}
