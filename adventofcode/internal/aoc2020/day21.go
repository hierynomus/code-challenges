package aoc2020

import (
	"bufio"
	"strconv"
	"strings"
)

func Day21(reader *bufio.Scanner) (string, string) {
	// a2i := map[string][]string{}
	i2a := map[string][]string{}
	// mapping := map[string]string{}

	for reader.Scan() {
		food := reader.Text()
		ia := strings.Split(food, " (contains ")
		ing := strings.Split(ia[0], " ")
		all := strings.Split(ia[1][0:len(ia[1])-1], " ")
		for _, i := range ing {
			if _, ok := i2a[i]; !ok {
				s := []string{}
				s = append(s, all...)
				i2a[i] = s
			}

		}
	}
	part1 := 0
	return strconv.Itoa(part1), ""
}
