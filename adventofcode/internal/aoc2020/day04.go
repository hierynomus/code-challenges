package aoc2020

import (
	"bufio"
	"strconv"
	"strings"
	"unicode"
)

type Passport map[string]string

func Day04(reader *bufio.Scanner) (string, string) {
	current := Passport{}
	part1 := 0
	part2 := 0
	for reader.Scan() {
		l := reader.Text()
		if len(strings.TrimSpace(l)) == 0 {
			if isValidPassport(current) {
				if isValidData(current) {
					part2 += 1
				}
				part1 += 1
			}
			current = Passport{}
		} else {
			for _, kvp := range strings.Split(l, " ") {
				kv := strings.Split(kvp, ":")
				current[kv[0]] = kv[1]
			}
		}
	}

	if isValidPassport(current) {
		if isValidData(current) {
			part2 += 1
		}
		part1 += 1
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func isValidPassport(p Passport) bool {
	_, cidPresent := p["cid"]
	return len(p) == 8 || (len(p) == 7 && !cidPresent)
}

//nolint:gomnd
func isValidData(p Passport) bool {
	byr := isIntBetween(p, "byr", 1920, 2002)
	iyr := isIntBetween(p, "iyr", 2010, 2020)
	eyr := isIntBetween(p, "eyr", 2020, 2030)
	hgt := hgtValid(p["hgt"])
	ecl := eclValid(p["ecl"])
	hcl := hclValid(p["hcl"])
	pid := pidValid(p["pid"])

	return byr && iyr && eyr && ecl && hcl && pid && hgt
}

func hclValid(c string) bool {
	if len(c) != 7 || c[0] != '#' {
		return false
	}

	for i := 1; i < 7; i++ {
		r := rune(c[i])
		if unicode.IsDigit(r) || 'a' == r || 'b' == r || 'c' == r || 'd' == r || 'e' == r || 'f' == r {
			continue
		} else {
			return false
		}
	}
	return true
}

func eclValid(c string) bool {
	valid := []string{"amb", "blu", "gry", "brn", "grn", "hzl", "oth"}
	for _, v := range valid {
		if c == v {
			return true
		}
	}
	return false
}
func isIntBetween(p Passport, key string, min int, max int) bool {
	v, err := strconv.Atoi(p[key])
	if err != nil {
		return false
	}

	return min <= v && v <= max
}

func pidValid(pid string) bool {
	if len(pid) != 9 {
		return false
	}

	_, err := strconv.Atoi(pid)
	return err == nil
}

func hgtValid(h string) bool {
	i, err := strconv.Atoi(h[0 : len(h)-2])
	if err != nil {
		return false
	}
	if strings.HasSuffix(h, "in") {
		return i >= 59 && i <= 76
	} else if strings.HasSuffix(h, "cm") {
		return i >= 150 && i <= 193
	}
	return false
}
