package aoc2018

import (
	"testing"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/day"
)

const SampleInput = `Immune System:
17 units each with 5390 hit points (weak to radiation, bludgeoning) with an attack that does 4507 fire damage at initiative 2
989 units each with 1274 hit points (immune to fire; weak to bludgeoning, slashing) with an attack that does 25 slashing damage at initiative 3

Infection:
801 units each with 4706 hit points (weak to radiation) with an attack that does 116 bludgeoning damage at initiative 1
4485 units each with 2961 hit points (immune to radiation; weak to fire, cold) with an attack that does 12 slashing damage at initiative 4
`

func TestDay24_sample(t *testing.T) {
	d := day.TestDay(t, Day24)
	d.WithInput(SampleInput, "5216", "51")
}
