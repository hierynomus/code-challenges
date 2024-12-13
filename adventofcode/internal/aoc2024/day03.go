package aoc2024

import (
	"bufio"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day03(reader *bufio.Scanner) (string, string) {
	var part1, part2 int64

	instructions := aoc.ReadStringArray(reader)
	enabled := true
	for _, instruction := range instructions {
		for i := 0; i < len(instruction)-4; i++ {
			if instruction[i:i+4] == "mul(" {
				for j := i + 4; j < len(instruction); j++ {
					if instruction[j] == ')' {
						s := strings.Split(instruction[i+4:j], ",")
						if len(s) != 2 {
							break
						}
						n1, err := aoc.ToIntE(s[0])
						if err != nil {
							break
						}
						n2, err := aoc.ToIntE(s[1])
						if err != nil {
							break
						}
						part1 += int64(n1 * n2)
						if enabled {
							part2 += int64(n1 * n2)
						}
						break
					}
				}
			} else if instruction[i:i+4] == "do()" {
				enabled = true
			} else if i < len(instruction)-7 && instruction[i:i+7] == "don't()" {
				enabled = false
			}
		}
	}

	return aoc.Int64ToString(part1), aoc.Int64ToString(part2)
}
