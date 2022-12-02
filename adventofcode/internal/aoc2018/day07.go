package aoc2018

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type Clock struct {
	Time int
}

func NewClock() *Clock {
	return &Clock{Time: 0}
}

func (c *Clock) Tick() {
	c.Time += 1
}

type WorkerElf struct {
	Clock    *Clock
	WorkItem rune
	ReadyAt  int
}

func NewElf(c *Clock) *WorkerElf {
	return &WorkerElf{Clock: c, WorkItem: rune(0)}
}

func (e *WorkerElf) WorkOn(item rune) {
	e.WorkItem = item
	e.ReadyAt = e.Clock.Time + int(item) - 4
}

func (e *WorkerElf) Ready() bool {
	return e.WorkItem != rune(0) && e.ReadyAt <= e.Clock.Time
}

func (e *WorkerElf) Idle() bool {
	return e.WorkItem == rune(0)
}

func (e *WorkerElf) Clear() {
	e.WorkItem = rune(0)
	e.ReadyAt = 0
}

func Day07(reader *bufio.Scanner) (string, string) {
	depsForward := map[rune][]rune{}
	depsReverse := map[rune][]rune{}
	for reader.Scan() {
		l := strings.Split(reader.Text(), " ")

		r1, r2 := rune(l[1][0]), rune(l[7][0])
		if _, ok := depsForward[r1]; !ok {
			depsForward[r1] = []rune{}
		}
		depsForward[r1] = append(depsForward[r1], r2)

		if _, ok := depsReverse[r2]; !ok {
			depsReverse[r2] = []rune{}
		}
		depsReverse[r2] = append(depsReverse[r2], r1)
	}

	part1 := d7p1(depsForward, depsReverse)
	_, part2 := d7p2(depsForward, depsReverse, 6)
	return part1, strconv.Itoa(part2)
}

func d7p1(forward, reverse map[rune][]rune) string {
	forwardDeps := map[rune][]rune{}
	for k, v := range forward {
		forwardDeps[k] = v
	}

	nextNodes := aoc.RuneSet{}
	for k := range forwardDeps {
		nextNodes.Add(k)
	}

	for k := range reverse {
		nextNodes.Delete(k)
	}

	result := []rune{}
	seen := aoc.RuneSet{}
	for len(nextNodes) > 0 {
		nextStart := nextNodes.Min()
		result = append(result, nextStart)
		seen.Add(nextStart)
		nextNodes.Delete(nextStart)
		if _, ok := forwardDeps[nextStart]; ok {
			next := forwardDeps[nextStart]
			for _, n := range next {
				if !seen.Contains(n) {
					needed := reverse[n]
					all := true
					for _, r := range needed {
						all = all && seen.Contains(r)
					}

					if all {
						nextNodes.Add(n)
					}
				}
			}
			delete(forwardDeps, nextStart)
		}
	}

	return string(result)
}

//nolint:funlen
func d7p2(forward, reverse map[rune][]rune, nrWorkers int) (string, int) {
	forwardDeps := map[rune][]rune{}
	for k, v := range forward {
		forwardDeps[k] = v
	}

	nextNodes := aoc.RuneSet{}
	for k := range forwardDeps {
		nextNodes.Add(k)
	}

	for k := range reverse {
		nextNodes.Delete(k)
	}

	result := []rune{}
	seen := aoc.RuneSet{}
	clock := NewClock()
	workers := []*WorkerElf{}
	for i := 0; i < nrWorkers; i++ {
		workers = append(workers, NewElf(clock))
	}

	for _, w := range workers {
		if len(nextNodes) > 0 {
			next := nextNodes.Min()
			w.WorkOn(next)
			nextNodes.Delete(next)
		}
	}

	for !allIdle(workers) {
		clock.Tick()
		for _, w := range workers {
			if w.Ready() {
				ready := w.WorkItem
				w.Clear()
				result = append(result, ready)
				seen.Add(ready)

				next := forwardDeps[ready]
				for _, n := range next {
					if !seen.Contains(n) {
						needed := reverse[n]
						all := true
						for _, r := range needed {
							all = all && seen.Contains(r)
						}

						if all {
							nextNodes.Add(n)
						}
					}
				}
				delete(forwardDeps, ready)
				for _, w := range workers {
					if w.Idle() && len(nextNodes) > 0 {
						i := nextNodes.Min()
						nextNodes.Delete(i)
						w.WorkOn(i)
					}
				}
			}
		}
	}

	return string(result), clock.Time
}

func allIdle(workers []*WorkerElf) bool {
	allIdle := true
	for _, w := range workers {
		allIdle = allIdle && w.Idle()
	}

	return allIdle
}
