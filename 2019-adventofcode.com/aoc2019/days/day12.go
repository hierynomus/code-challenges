package days

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/hierynomus/aoc2019/aoc"
)

type Day12 struct{}

type JupiterSystem struct {
	moons []*Moon
}

func (s *JupiterSystem) Copy() *JupiterSystem {
	newMoons := []*Moon{}
	for _, m := range s.moons {
		newMoons = append(newMoons, NewMoon(m.L))
	}
	return &JupiterSystem{newMoons}
}

func (s *JupiterSystem) ApplyGravity() {
	for _, x := range s.moons {
		for _, y := range s.moons {
			if x == y {
				continue
			}
			x.ApplyGravity(y)
		}
	}
}

func (s *JupiterSystem) Move() {
	for _, m := range s.moons {
		m.Move()
	}
}

func (s *JupiterSystem) AllCycled() bool {
	for _, m := range s.moons {
		if !m.Cycled() {
			return false
		}
	}
	return true
}

type Moon struct {
	Initial aoc.Point3D
	L       aoc.Point3D
	V       aoc.Point3D
	time    int64
	cycles  aoc.Point3D
}

func NewMoon(p aoc.Point3D) *Moon {
	return &Moon{
		Initial: p,
		L:       p,
		V:       aoc.ZeroPoint3D(),
		time:    0,
		cycles:  aoc.ZeroPoint3D(),
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

func (m *Moon) Cycled() bool {
	return m.cycles.X > 0 && m.cycles.Y > 0 && m.cycles.Z > 0
}

func (m *Moon) Move() {
	m.L = aoc.Point3D{
		X: m.L.X + m.V.X,
		Y: m.L.Y + m.V.Y,
		Z: m.L.Z + m.V.Z,
	}
	m.time++
	if !m.Cycled() {
		if m.L.X == m.Initial.X && m.V.X == 0 && m.cycles.X == 0 {
			m.cycles = aoc.Point3D{X: m.time, Y: m.cycles.Y, Z: m.cycles.Z}
		}
		if m.L.Y == m.Initial.Y && m.V.Y == 0 && m.cycles.Y == 0 {
			m.cycles = aoc.Point3D{X: m.cycles.X, Y: m.time, Z: m.cycles.Z}
		}
		if m.L.Z == m.Initial.Z && m.V.Z == 0 && m.cycles.Z == 0 {
			m.cycles = aoc.Point3D{X: m.cycles.X, Y: m.cycles.Y, Z: m.time}
		}
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

func (d *Day12) Solve(scanner *bufio.Scanner) (string, string) {
	system := &JupiterSystem{moons: []*Moon{}}
	for scanner.Scan() {
		l := strings.Trim(scanner.Text(), "<>")
		xyz := strings.Split(l, ", ")
		coords := []int64{}
		for _, i := range xyz {
			j, _ := strconv.Atoi(strings.SplitAfter(i, "=")[1])
			coords = append(coords, int64(j))
		}
		system.moons = append(system.moons, NewMoon(aoc.Point3D{X: coords[0], Y: coords[1], Z: coords[2]}))
	}

	part1 := system.Copy()
	for i := 0; i < 1000; i++ {
		part1.ApplyGravity()
		part1.Move()
	}

	var totalEnergy int64 = 0
	for _, m := range part1.moons {
		totalEnergy += m.Energy()
	}

	part2 := system.Copy()
	for !part2.AllCycled() {
		part2.ApplyGravity()
		part2.Move()
	}

	lcms := []int64{}
	for _, m := range part2.moons {
		fmt.Printf("%v\n", m)
		lcm := aoc.Lcm(m.cycles.X, m.cycles.Y)
		lcm = aoc.Lcm(lcm, m.cycles.Z)
		fmt.Printf("%d\n", lcm)
		lcms = append(lcms, lcm)
	}
	return fmt.Sprintf("%d", totalEnergy), fmt.Sprintf("%d", aoc.LcmArray(lcms))
}
