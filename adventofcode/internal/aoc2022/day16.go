package aoc2022

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type Memo struct {
	Memo map[string]map[int]Flow
}

type Valve struct {
	Tunnels  map[string]int
	Name     string
	FlowRate int
}

type Flow struct {
	ValvesOpen []string
	Rate       int
}

func (f Flow) Add(open string, rate int) Flow {
	return Flow{
		ValvesOpen: append(f.ValvesOpen, open),
		Rate:       f.Rate + rate,
	}
}

var EmptyFlow = Flow{ValvesOpen: []string{}, Rate: 0}

func ParseValve(line string) Valve {
	s := strings.Split(line, " ")
	name := s[1]
	flowRate := aoc.ToInt(strings.TrimRight(strings.Split(s[4], "=")[1], ";"))
	tunnels := map[string]int{}

	for i := 9; i < len(s); i++ {
		tunnels[strings.TrimRight(s[i], ",")] = 1
	}

	return Valve{
		Tunnels:  tunnels,
		Name:     name,
		FlowRate: flowRate,
	}
}

type CaveSystem map[string]Valve

func (c CaveSystem) FlowRate(openedValves map[string]int) int {
	total := 0
	for k, v := range openedValves {
		total += (c[k].FlowRate * (30 - v))
	}

	return total
}

func (c CaveSystem) Simplify() {
	for _, valve := range c {
		if valve.Name == "AA" {
			continue
		}
		if valve.FlowRate == 0 {
			delete(c, valve.Name) // take me out
			for source, dist := range valve.Tunnels {
				tunnelsOfSource := c[source].Tunnels
				delete(tunnelsOfSource, valve.Name)
				for dest, destDist := range valve.Tunnels {
					if dest != source {
						tunnelsOfSource[dest] = dist + destDist
					}
				}
			}
		}
	}
}

func copyMap[K, V comparable](m map[K]V) map[K]V {
	result := make(map[K]V)
	for k, v := range m {
		result[k] = v
	}
	return result
}

func (c CaveSystem) MaxFlowRate(currValve string, time int, valvesOpened []string, memo *Memo) Flow {
	memoed := memo.Memo[currValve]
	if f, ok := memoed[time]; ok {
		return f
	}

	if time > 29 {
		return EmptyFlow
	}

	maxFlow := EmptyFlow
	valve := c[currValve]
	if !aoc.StringArrayContains(valvesOpened, currValve) && valve.FlowRate > 0 {
		r := valve.FlowRate * (30 - (time + 1))
		for v, d := range valve.Tunnels {
			f := c.MaxFlowRate(v, time+d+1, append(valvesOpened, currValve), memo)
			f = f.Add(currValve, r)

			if f.Rate > maxFlow.Rate {
				maxFlow = f
			}
		}
	}

	for v, d := range valve.Tunnels {
		f := c.MaxFlowRate(v, time+d, valvesOpened, memo)
		if f.Rate > maxFlow.Rate {
			maxFlow = f
		}
	}

	memo.Memo[currValve][time] = maxFlow

	return maxFlow
}

// func FindMaxFlowRate(caveSystem CaveSystem, currValve string, valvesOpen []string, memo map[string]map[int]Flow, time int) Flow {
// 	if _, ok := memo[currValve]; !ok {
// 		memo[currValve] = map[int]Flow{}
// 	}

// 	v, ok := memo[currValve][time]
// 	if ok {
// 		return v
// 	}

// 	if time >= 30 {
// 		return
// 	}

// 	v := caveSystem[currValve]
// 	maxFR := 0
// 	var maxOpened map[string]int
// 	if _, ok := opened[currValve]; !ok {
// 		time++
// 		newOpened := copyMap(opened)
// 		newOpened[currValve] = time
// 		if time == 30 {
// 			return newOpened
// 		}

// 		for t, d := range v.Tunnels {
// 			o := FindMaxFlowRate(caveSystem, t, newOpened, time+d)
// 			fr := caveSystem.FlowRate(o)
// 			if fr > maxFR {
// 				maxFR = fr
// 				maxOpened = o
// 			}
// 		}
// 	}

// 	for t, d := range v.Tunnels {
// 		o := FindMaxFlowRate(caveSystem, t, opened, time+d)
// 		fr := caveSystem.FlowRate(o)
// 		if fr > maxFR {
// 			maxFR = fr
// 			maxOpened = o
// 		}
// 	}

// 	return maxOpened
// }

func Day16(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	caveSystem := CaveSystem{}

	for reader.Scan() {
		line := reader.Text()
		v := ParseValve(line)
		caveSystem[v.Name] = v
	}

	caveSystem.Simplify()

	memo := &Memo{
		Memo: map[string]map[int]Flow{},
	}

	for _, v := range caveSystem {
		memo.Memo[v.Name] = map[int]Flow{}
	}

	flow := caveSystem.MaxFlowRate("AA", 0, []string{}, memo)

	part1 = flow.Rate
	// part1 = caveSystem.FlowRate(opened)

	return strconv.Itoa(part1), strconv.Itoa(part2)
}
