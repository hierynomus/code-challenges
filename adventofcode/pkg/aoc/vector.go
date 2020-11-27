package aoc

import (
	"fmt"
	"math"
)

type Vector struct {
	Radians float64
	Length  float64
}

func (v *Vector) String() string {
	return fmt.Sprintf("V(%f, %f)", v.Radians, v.Length)
}

func CreateVector(a *Point, o *Point) *Vector {
	angle := math.Atan2(float64(o.X-a.X), float64(a.Y-o.Y)) * 180 / math.Pi
	if angle < 0 {
		angle += 360
	}

	return &Vector{
		Radians: angle,
		Length:  math.Sqrt(float64((a.X - o.X) ^ 2 + (a.Y - o.Y) ^ 2)),
	}
}
