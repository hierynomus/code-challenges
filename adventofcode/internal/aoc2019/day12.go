package aoc2019

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type JupiterSystem struct {
	Moons []*Moon
}

func (s *JupiterSystem) Copy() *JupiterSystem {
	newMoons := []*Moon{}
	for _, m := range s.Moons {
		newMoons = append(newMoons, NewMoon(m.L))
	}

	return &JupiterSystem{newMoons}
}

func (s *JupiterSystem) ApplyGravity() {
	for _, x := range s.Moons {
		for _, y := range s.Moons {
			if x == y {
				continue
			}

			x.ApplyGravity(y)
		}
	}
}

func (s *JupiterSystem) Move() {
	for _, m := range s.Moons {
		m.Move()
	}
}

// func (s *JupiterSystem) AllCycled() bool {
// 	for _, m := range s.moons {
// 		if !m.Cycled() {
// 			return false
// 		}
// 	}
// 	return true
// }

type Moon struct {
	L      aoc.Point3D
	V      aoc.Point3D
	cycles aoc.Point3D
}

func NewMoon(p aoc.Point3D) *Moon {
	return &Moon{
		L: p,
		V: aoc.ZeroPoint3D(),
	}
}

func (m *Moon) String() string {
	return fmt.Sprintf("Moon(loc: %v, vel: %v, cycle: %v)", m.L, m.V, m.cycles)
}

func (m *Moon) ApplyGravity(o *Moon) {
	m.V = aoc.Point3D{
		X: m.V.X + m.adj(m.L.X, o.L.X),
		Y: m.V.Y + m.adj(m.L.Y, o.L.Y),
		Z: m.V.Z + m.adj(m.L.Z, o.L.Z),
	}
}

func (m *Moon) Move() {
	m.L = aoc.Point3D{
		X: m.L.X + m.V.X,
		Y: m.L.Y + m.V.Y,
		Z: m.L.Z + m.V.Z,
	}
}

func (m *Moon) Energy() int64 {
	potential := m.energy(m.L)
	kinetic := m.energy(m.V)

	return potential * kinetic
}

func (m *Moon) energy(p aoc.Point3D) int64 {
	return aoc.Abs64(p.X) + aoc.Abs64(p.Y) + aoc.Abs64(p.Z)
}

func (m *Moon) adj(self, other int64) int64 {
	if self > other {
		return -1
	} else if self < other {
		return 1
	}

	return 0
}

func ParseMoon(line string) *Moon {
	l := strings.Trim(line, "<>")
	xyz := strings.Split(l, ", ")
	coords := []int64{}

	for _, i := range xyz {
		j, _ := strconv.Atoi(strings.SplitAfter(i, "=")[1])
		coords = append(coords, int64(j))
	}

	return NewMoon(aoc.Point3D{X: coords[0], Y: coords[1], Z: coords[2]})
}

func Day12(scanner *bufio.Scanner) (string, string) {
	system := &JupiterSystem{Moons: []*Moon{}}
	for scanner.Scan() {
		system.Moons = append(system.Moons, ParseMoon(scanner.Text()))
	}

	part1 := system.Copy()

	for i := 0; i < 1000; i++ {
		part1.ApplyGravity()
		part1.Move()
	}

	var totalEnergy int64 = 0
	for _, m := range part1.Moons {
		totalEnergy += m.Energy()
	}

	part2 := system.Copy()
	cycled := make([]int64, 3)
	for time := int64(1); cycled[0] == 0 || cycled[1] == 0 || cycled[2] == 0; time++ {
		part2.ApplyGravity()
		part2.Move()

		foundX, foundY, foundZ := true, true, true
		for i, m := range part2.Moons {
			initial := system.Moons[i]
			if cycled[0] == 0 && foundX {
				if m.L.X != initial.L.X || m.V.X != initial.V.X {
					foundX = false
				}
			}

			if cycled[1] == 0 && foundY {
				if m.L.Y != initial.L.Y || m.V.Y != initial.V.Y {
					foundY = false
				}
			}

			if cycled[2] == 0 && foundZ {
				if m.L.Z != initial.L.Z || m.V.Z != initial.V.Z {
					foundZ = false
				}
			}

		}

		if foundX && cycled[0] == 0 {
			cycled[0] = time
		}

		if foundY && cycled[1] == 0 {
			cycled[1] = time
		}

		if foundZ && cycled[2] == 0 {
			cycled[2] = time
		}
	}

	lcm := aoc.LcmArray(cycled)

	return aoc.Int64ToString(totalEnergy), aoc.Int64ToString(lcm)
}
