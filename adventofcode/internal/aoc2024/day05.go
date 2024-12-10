package aoc2024

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type PrintOrder struct {
	Left, Right int
}

func (po PrintOrder) String() string {
	return fmt.Sprintf("%d|%d", po.Left, po.Right)
}

func (po PrintOrder) FindIndices(job []int) (int, int) {
	idxLeft, idxRight := -1, -1
	for i, j := range job {
		if j == po.Left {
			idxLeft = i
		}
		if j == po.Right {
			idxRight = i
		}
	}

	return idxLeft, idxRight
}

func (po PrintOrder) Satisfies(job []int) bool {
	idxLeft, idxRight := po.FindIndices(job)
	if idxLeft == -1 || idxRight == -1 {
		return true
	}

	return idxLeft < idxRight
}

func (po PrintOrder) Reorder(job []int) []int {
	idxLeft, idxRight := po.FindIndices(job)

	if idxLeft == -1 || idxRight == -1 || idxLeft < idxRight {
		return job
	}

	job[idxLeft], job[idxRight] = job[idxRight], job[idxLeft]
	return job
}

func Day05(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	orders := []PrintOrder{}
	for reader.Scan() {
		line := reader.Text()
		if line == "" {
			break
		}
		lr := strings.Split(line, "|")
		orders = append(orders, PrintOrder{Left: aoc.ToInt(lr[0]), Right: aoc.ToInt(lr[1])})
	}

	jobs := [][]int{}

	for reader.Scan() {
		line := reader.Text()
		if line == "" {
			break
		}

		jobs = append(jobs, aoc.ToIntArray(strings.Split(line, ",")))
	}

	invalidJobs := make([][]int, 0)
	for _, job := range jobs {
		valid := true
		for _, order := range orders {
			if !order.Satisfies(job) {
				invalidJobs = append(invalidJobs, job)
				valid = false
				break
			}
		}

		if valid {
			part1 += job[len(job)/2]
		}
	}

	for _, job := range invalidJobs {
		validJob := checkAndReorder(job, orders)
		part2 += validJob[len(validJob)/2]
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func checkAndReorder(job []int, orders []PrintOrder) []int {
	for _, order := range orders {
		if !order.Satisfies(job) {
			job = order.Reorder(job)
			return checkAndReorder(job, orders)
		}
	}

	return job
}
