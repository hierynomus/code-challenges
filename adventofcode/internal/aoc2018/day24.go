package aoc2018

import (
	"bufio"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type Team string

const (
	None         Team = ""
	ImmuneSystem Team = "Immune System"
	Infection    Team = "Infection"
)

type Group struct {
	Team                          Team
	Name                          string
	AttackType                    string
	Weaknesses, Immunities        []string
	Number                        int
	Units, HP, Damage, Initiative int
}

func (g *Group) EffectivePower() int {
	return g.Units * g.Damage
}

func (g *Group) DamageTo(other *Group) int {
	if aoc.StringArrayContains(other.Immunities, g.AttackType) {
		return 0
	}

	damage := g.EffectivePower()
	if aoc.StringArrayContains(other.Weaknesses, g.AttackType) {
		damage *= 2
	}

	return damage
}

func ParseTeams(reader *bufio.Scanner) []Group {
	groups := []Group{}

	for reader.Scan() {
		teamS := reader.Text()
		team := Team(teamS[:len(teamS)-1])
		idx := 1
		for reader.Scan() {
			line := reader.Text()
			if line == "" {
				break
			}
			groups = append(groups, ParseGroup(line, team, idx))
			idx++
		}
	}

	return groups
}

func ParseGroup(line string, team Team, nr int) Group {
	g := Group{Name: fmt.Sprintf("%s group %d", team, nr), Number: nr, Team: team}
	r := regexp.MustCompile(`(\d+) units each with (\d+) hit points( \(.*\))? with an attack that does (\d+) (\w+) damage at initiative (\d+)`)
	m := r.FindStringSubmatch(line)
	g.Units, _ = strconv.Atoi(m[1])
	g.HP, _ = strconv.Atoi(m[2])
	g.Damage, _ = strconv.Atoi(m[4])
	g.AttackType = m[5]
	g.Initiative, _ = strconv.Atoi(m[6])

	if m[3] != "" {
		r = regexp.MustCompile(`(\w+) to ([\w, ]+)`)
		matches := r.FindAllStringSubmatch(m[3], -1)
		for _, m := range matches {
			switch m[1] {
			case "weak":
				g.Weaknesses = strings.Split(m[2], ", ")
			case "immune":
				g.Immunities = strings.Split(m[2], ", ")
			}
		}
	}

	return g
}

func Fight(orig []Group) (Team, int) {
	groups := copyGroups(orig)

	for {
		targets := SelectTargets(groups)
		groupsLeft, unitsDied := Attack(groups, targets)
		if !unitsDied {
			return None, 0
		}

		groups = RemoveDead(groupsLeft)

		immuneSystem, infection := CountGroups(groups)
		if immuneSystem == 0 || infection == 0 {
			// After RemoveDead there will only be one team left
			return groups[0].Team, immuneSystem + infection
		}
	}
}

func SelectTargets(groups []*Group) map[string]string {
	targets := map[string]string{}
	targeted := map[string]bool{}

	// Sort groups by effective power, then initiative
	sort.Slice(groups, func(i, j int) bool {
		if groups[i].EffectivePower() == groups[j].EffectivePower() {
			return groups[i].Initiative > groups[j].Initiative
		}
		return groups[i].EffectivePower() > groups[j].EffectivePower()
	})

	// Select targets
	for _, attacker := range groups {
		maxDamage := 0
		var currentTarget *Group
		for _, defender := range groups {
			if attacker.Name == defender.Name || targeted[defender.Name] || attacker.Team == defender.Team {
				continue
			}

			damage := attacker.DamageTo(defender)
			if damage > maxDamage {
				maxDamage = damage
				currentTarget = defender
			} else if damage == maxDamage && currentTarget != nil {
				if defender.EffectivePower() > currentTarget.EffectivePower() {
					currentTarget = defender
				} else if defender.EffectivePower() == currentTarget.EffectivePower() {
					if defender.Initiative > currentTarget.Initiative {
						currentTarget = defender
					}
				}
			}
		}

		if maxDamage > 0 {
			targets[attacker.Name] = currentTarget.Name
			targeted[currentTarget.Name] = true
		}
	}

	return targets
}

func Attack(groups []*Group, targets map[string]string) ([]*Group, bool) {
	// Sort groups by initiative
	sort.Slice(groups, func(i, j int) bool {
		return groups[i].Initiative > groups[j].Initiative
	})

	lookup := map[string]*Group{}
	for _, g := range groups {
		lookup[g.Name] = g
	}

	unitsDied := false

	// Attack
	for _, attacker := range groups {
		target, ok := targets[attacker.Name]
		if !ok {
			continue
		}

		if attacker.Units <= 0 {
			continue
		}

		defender := lookup[target]

		damage := attacker.DamageTo(defender)
		unitsKilled := aoc.Min([]int{damage / defender.HP, defender.Units})
		if unitsKilled > 0 {
			unitsDied = true
		}

		defender.Units -= unitsKilled
	}

	return groups, unitsDied
}

func RemoveDead(groups []*Group) []*Group {
	alive := []*Group{}
	for _, g := range groups {
		if g.Units > 0 {
			alive = append(alive, g)
		}
	}

	return alive
}

func CountGroups(groups []*Group) (int, int) {
	immuneSystem, infection := 0, 0
	for _, g := range groups {
		if g.Team == ImmuneSystem {
			immuneSystem += g.Units
		} else {
			infection += g.Units
		}
	}

	return immuneSystem, infection
}

func copyGroups(groups []Group) []*Group {
	c := []*Group{}
	for _, g := range groups {
		gg := g
		c = append(c, &gg)
	}

	return c
}

func Day24(reader *bufio.Scanner) (string, string) {
	var part1, part2 int

	groups := ParseTeams(reader)

	// Part 1
	_, units := Fight(groups)
	part1 = units

	// // Part 2
	boost := 1
	for {
		boosted := []Group{}
		for _, g := range groups {
			if g.Team == ImmuneSystem {
				g.Damage += boost
			}
			boosted = append(boosted, g)
		}

		team, units := Fight(boosted)
		if team == ImmuneSystem {
			part2 = units
			break
		}

		boost++
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}
