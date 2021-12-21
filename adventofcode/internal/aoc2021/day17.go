package aoc2021

import (
	"bufio"
	"fmt"
	"math"
	"strconv"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type ProbeVector struct {
	vx, vy int
}

func (pv ProbeVector) String() string {
	return fmt.Sprintf("(%d, %d)", pv.vx, pv.vy)
}

func Day17(reader *bufio.Scanner) (string, string) {
	var part1, part2 int
	var xmin, xmax, ymin, ymax int

	_, err := fmt.Sscanf(aoc.Read(reader), "target area: x=%d..%d, y=%d..%d", &xmin, &xmax, &ymin, &ymax)
	if err != nil {
		panic(err)
	}

	vymax := -ymin // i.e. it's a parabolic curve, which passes at time `t` through the y=0 line. So the max Y velocity at the end which 'just' passes through the target area should be at `t+1`
	part1 = (vymax * (vymax - 1) / 2)

	// xmax is the max X velocity which hits the target area at `t=1`, so that is the upper bound.
	vxmax := xmax
	// ymin is also the min Y velocity which hits the target area at `t=1`, so that is the lower y bound.
	vymin := ymin

	// The minimum X velocity reaches xmin at time = t and is then '0', so
	// because a = -1 and d = xmin, so:
	// d = vxmin * t + 1/2 * a * t^2 ==
	// xmin = vxmin * t + 1/2 * -1 * t^2 ==
	// v * v - 1/2*v^2 == xmin
	vxmin := int(math.Floor(math.Sqrt(float64(2 * xmin))))

	ty := ProbeTimeY(ymin, ymax, vymin, vymax)
	velocities := map[int]aoc.IntSet{}
	for t, vys := range ty {
		for vx := vxmin; vx <= vxmax; vx++ {
			dx := ProbeXDistance(vx, t)
			if dx >= xmin && dx <= xmax {
				for vy := range aoc.NewIntSet(vys) {
					if _, ok := velocities[vx]; !ok {
						velocities[vx] = aoc.NewIntSet([]int{})
					}
					velocities[vx].Add(vy)
				}
			}
		}
	}

	for _, vys := range velocities {
		part2 += len(vys)
	}

	return strconv.Itoa(part1), strconv.Itoa(part2)
}

func ProbeTimeY(ymin, ymax, vymin, vymax int) map[int][]int {
	ty := map[int][]int{}
	vy := vymax
	t := 1
	for {
		if vy < vymin {
			break
		}

		dy := ProbeYDistance(vy, t)
		if dy >= ymin && dy <= ymax {
			if _, ok := ty[t]; !ok {
				ty[t] = []int{}
			}

			ty[t] = append(ty[t], vy)
		} else if dy < ymin {
			t = 1
			vy -= 1
			continue
		}
		t += 1
	}

	return ty
}

func ProbeXDistance(vx, t int) int {
	s := 0
	for i := 0; i < t; i++ {
		if vx-i <= 0 {
			break
		}
		s += vx - i
	}
	return s
}

func ProbeYDistance(vx, t int) int {
	return vx*t - ((t-1)*t)/2
}
