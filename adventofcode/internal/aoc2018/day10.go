package aoc2018

import (
	"bufio"
	"strconv"
	"strings"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type PoleLights []*PoleLight

type PoleLight struct {
	X, Y   int
	DX, DY int
}

func (l *PoleLight) AtTime(t int) *PoleLight {
	return &PoleLight{
		X:  l.X + t*l.DX,
		Y:  l.Y + t*l.DY,
		DX: l.DX,
		DY: l.DY,
	}
}

func (ls PoleLights) MinX() int {
	x := ls[0].X
	for _, l := range ls {
		if l.X < x {
			x = l.X
		}
	}

	return x
}

func (ls PoleLights) MinY() int {
	y := ls[0].Y
	for _, l := range ls {
		if l.Y < y {
			y = l.Y
		}
	}

	return y
}

func (ls PoleLights) MaxX() int {
	x := ls[0].X
	for _, l := range ls {
		if l.X > x {
			x = l.X
		}
	}

	return x
}

func (ls PoleLights) MaxY() int {
	y := ls[0].Y
	for _, l := range ls {
		if l.Y > y {
			y = l.Y
		}
	}

	return y
}

func (ls PoleLights) AtTime(t int) PoleLights {
	nl := PoleLights{}
	for _, l := range ls {
		nl = append(nl, l.AtTime(t))
	}
	return nl
}

func (ls PoleLights) BoundingBoxSize() int {
	return (ls.MaxX() - ls.MinX()) * (ls.MaxY() - ls.MinY())
}

func Day10(reader *bufio.Scanner) (string, string) {
	lights := PoleLights{}
	for reader.Scan() {
		l := reader.Text()
		xy := strings.Split(l[10:24], ",")
		dxy := strings.Split(l[len(l)-7:len(l)-1], ",")
		lights = append(lights, &PoleLight{
			X:  aoc.ToInt(xy[0]),
			Y:  aoc.ToInt(xy[1]),
			DX: aoc.ToInt(dxy[0]),
			DY: aoc.ToInt(dxy[1]),
		})
	}

	minBoxArea := lights.BoundingBoxSize()
	minBoxTime := 0
	t := 0

	for true {
		t += 1
		nl := lights.AtTime(t)
		mb := nl.BoundingBoxSize()
		if mb < minBoxArea {
			minBoxArea = mb
			minBoxTime = t
		} else {
			break
		}
	}

	lightsAt := lights.AtTime(minBoxTime)
	rendering := make([][]rune, lightsAt.MaxY()-lightsAt.MinY()+1)
	for y := 0; y < len(rendering); y++ {
		rendering[y] = make([]rune, lightsAt.MaxX()-lightsAt.MinX()+1)
		for x := 0; x < len(rendering[y]); x++ {
			rendering[y][x] = '.'
		}
	}

	for _, l := range lightsAt {
		rendering[l.Y-lightsAt.MinY()][l.X-lightsAt.MinX()] = '#'
	}

	render := ""
	for _, r := range rendering {
		render += string(r) + "\n"
	}

	return strings.TrimSpace(render), strconv.Itoa(minBoxTime)
}
