package days

import (
	"bufio"
	"fmt"
	"github.com/hierynomus/aoc2019/aoc"
	"strconv"
	"strings"
)

type Day06 struct{}

var SpaceObjects map[string]*SpaceObject = map[string]*SpaceObject{}

func GetOrCreate(object string) *SpaceObject {
	var o *SpaceObject
	if v, ok := SpaceObjects[object]; ok {
		o = v
	} else {
		o = &SpaceObject{object, nil, []*SpaceObject{}}
		SpaceObjects[object] = o
	}
	return o
}

type SpaceObject struct {
	object   string
	orbiting *SpaceObject
	orbiters []*SpaceObject
}

func (s *SpaceObject) String() string {
	return fmt.Sprintf("%s->%s", s.object, s.orbiting)
}

func (d *Day06) Solve(scanner *bufio.Scanner) (string, string) {
	for scanner.Scan() {
		orbit := strings.Split(scanner.Text(), ")")
		planet1, planet2 := GetOrCreate(orbit[0]), GetOrCreate(orbit[1])
		if planet2.orbiting != nil {
			panic(fmt.Errorf("BooM!"))
		}
		planet2.orbiting = planet1
		planet1.orbiters = append(planet1.orbiters, planet2)
	}

	orbits := 0
	for _, v := range SpaceObjects {
		for v.orbiting != nil {
			orbits++
			v = v.orbiting
		}
	}

	me := SpaceObjects["YOU"]
	santa := SpaceObjects["SAN"]

	fmt.Printf("%s\n", me)
	fmt.Printf("%s\n", santa)

	orbitListSanta := []string{}
	o := santa
	for o.orbiting != nil {
		orbitListSanta = append(orbitListSanta, o.orbiting.object)
		o = o.orbiting
	}

	m := me
	transfers := 0
	for m.orbiting != nil {
		i := aoc.StringArrayIndex(orbitListSanta, m.orbiting.object)
		if i != -1 {
			transfers += i
			break
		} else {
			transfers++
			m = m.orbiting
		}
	}

	return strconv.Itoa(orbits), strconv.Itoa(transfers)
}
