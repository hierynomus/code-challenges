package aoc2022

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
	"golang.org/x/exp/maps"
)

type Production string
type Robot string

const (
	OreRobot      = Robot("Ore Robot")
	ClayRobot     = Robot("Clay Robot")
	ObsidianRobot = Robot("Obsidian Robot")
	GeodeRobot    = Robot("Geode Robot")
	Ore           = Production("Ore")
	Clay          = Production("Clay")
	Obsidian      = Production("Obsidian")
	Geode         = Production("Geode")
)

func (r Robot) Produce() Production {
	switch r {
	case OreRobot:
		return Ore
	case ClayRobot:
		return Clay
	case ObsidianRobot:
		return Obsidian
	case GeodeRobot:
		return Geode
	}

	panic("unknown robot")
}

type Blueprint []Rule

func NewBlueprint() Blueprint {
	return Blueprint{}
}

func (b Blueprint) Run(state State) State {
	for state.Time < 24 {
		state = b.RunOnce(state)
	}

	return state
}

func (b Blueprint) RunOnce(state State) State {
	newState := state.Copy()
	newState.Time++
	for _, rule := range b {
		if rule.CanApply(state.Stock) {
			for k, v := range rule.Needs {
				newState.Stock[k] -= v
			}
			newState.Robots[rule.Target]++
		}
	}

	for r, n := range state.Robots {
		if n > 0 {
			newState.Stock[r.Produce()] += n
		}
	}

	return newState
}

type Rule struct {
	Needs  map[Production]int
	Target Robot
}

func (r Rule) CanApply(stock map[Production]int) bool {
	for k, v := range r.Needs {
		if stock[k] < v {
			return false
		}
	}

	return true
}

type State struct {
	Robots map[Robot]int
	Stock  map[Production]int
	Time   int
}

func NewState() State {
	return State{
		Robots: map[Robot]int{
			OreRobot:      1,
			ClayRobot:     0,
			ObsidianRobot: 0,
			GeodeRobot:    0,
		},
		Stock: map[Production]int{
			Ore:      0,
			Clay:     0,
			Obsidian: 0,
			Geode:    0,
		},
		Time: 0,
	}
}

func (s State) Copy() State {
	return State{
		Robots: maps.Clone(s.Robots),
		Stock:  maps.Clone(s.Stock),
		Time:   s.Time,
	}
}

func Day19(reader *bufio.Scanner) (string, string) {
	var part1, part2 int
	blueprints := []Blueprint{}

	for reader.Scan() {
		bp := reader.Text()
		s := strings.Split(bp, " ")
		b := NewBlueprint()
		b = append(b,
			Rule{Needs: map[Production]int{Ore: aoc.ToInt(s[6])}, Target: OreRobot},
			Rule{Needs: map[Production]int{Clay: aoc.ToInt(s[12])}, Target: ClayRobot},
			Rule{Needs: map[Production]int{Ore: aoc.ToInt(s[18]), Clay: aoc.ToInt(s[21])}, Target: ObsidianRobot},
			Rule{Needs: map[Production]int{Ore: aoc.ToInt(s[27]), Obsidian: aoc.ToInt(s[30])}, Target: GeodeRobot},
		)
		blueprints = append(blueprints, b)
	}

	for i, blueprint := range blueprints {
		s := NewState()
		out := blueprint.Run(s)
		part1 += (i + 1) * out.Stock[Geode]
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}
