package days

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Day06 struct{}

type SpaceObject struct {
	object   string
	orbiting *SpaceObject
}

func (s *SpaceObject) String() string {
	return fmt.Sprintf("%s->%s", s.object, s.orbiting)
}

type SpaceMap map[string]*SpaceObject

func (s SpaceMap) GetOrCreate(object string) *SpaceObject {
	var o *SpaceObject
	if v, ok := s[object]; ok {
		o = v
	} else {
		o = &SpaceObject{object, nil}
		s[object] = o
	}
	return o
}

func (d *Day06) Solve(scanner *bufio.Scanner) (string, string) {
	spaceMap := make(SpaceMap)
	for scanner.Scan() {
		orbit := strings.Split(scanner.Text(), ")")
		o1, o2 := spaceMap.GetOrCreate(orbit[0]), spaceMap.GetOrCreate(orbit[1])
		if o2.orbiting != nil {
			panic(fmt.Errorf("BooM!"))
		}
		o2.orbiting = o1
	}

	orbits := 0
	for _, v := range spaceMap {
		for v.orbiting != nil {
			orbits++
			v = v.orbiting
		}
	}

	me := spaceMap["YOU"]
	santa := spaceMap["SAN"]

	transfers := 0
	found := false
	for mDist, m := 0, me.orbiting; !found && m.orbiting != nil; mDist, m = mDist+1, m.orbiting {
		for sDist, s := 0, santa.orbiting; !found && s.orbiting != nil; sDist, s = sDist+1, s.orbiting {
			if s == m {
				transfers = mDist + sDist
				found = true
			}
		}
	}

	return strconv.Itoa(orbits), strconv.Itoa(transfers)
}
