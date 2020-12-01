package aoc2018

import (
	"bufio"
	"sort"
	"strconv"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

func Day04(reader *bufio.Scanner) (string, string) {
	input := []string{}

	for reader.Scan() {
		input = append(input, reader.Text())
	}

	sort.Strings(input)

	var currentGuard, sleepyTime, wakeyTime int
	totalSleep := map[int]int{}
	sleepsAt := map[int]aoc.IntHistogram{}

	for _, s := range input {
		if strings.Contains(s, "Guard") {
			currentGuard = aoc.ToInt(strings.Split(s, " ")[3][1:])
		} else if strings.Contains(s, "asleep") {
			t := strings.Split(s, " ")[1]
			sleepyTime = aoc.ToInt(strings.Split(t[0:len(t)-1], ":")[1])
		} else if strings.Contains(s, "wakes up") {
			t := strings.Split(s, " ")[1]
			wakeyTime = aoc.ToInt(strings.Split(t[0:len(t)-1], ":")[1])
			totalSleep[currentGuard] += wakeyTime - sleepyTime

			if _, ok := sleepsAt[currentGuard]; !ok {
				sleepsAt[currentGuard] = aoc.IntHistogram{}
			}

			sleepsAt[currentGuard].Adds(aoc.Range(sleepyTime, wakeyTime))
		}
	}

	var sleepiest int
	max := 0
	for k, v := range totalSleep {
		if v > max {
			sleepiest, max = k, v
		}
	}

	min, _ := sleepsAt[sleepiest].Max()
	part1 := sleepiest * min

	var sameMinuteSleeper int
	minute, times := 0, 0
	for g, h := range sleepsAt {
		m, t := h.Max()
		if t > times {
			sameMinuteSleeper, minute, times = g, m, t
		}
	}

	part2 := sameMinuteSleeper * minute

	return strconv.Itoa(part1), strconv.Itoa(part2)
}
