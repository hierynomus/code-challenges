package days

import (
	"bufio"
	"unicode"

	"github.com/hierynomus/aoc2019/aoc"
)

type VaultWalk struct {
	Pos        aoc.Point
	WalkLength int
	KeysFound  []rune
	Index      int
}

func NewWalk(pos aoc.Point, length int, keys []rune) *VaultWalk {
	kf := make([]rune, len(keys))
	copy(kf, keys)
	return &VaultWalk{
		Pos:        pos,
		WalkLength: length,
		KeysFound:  keys,
	}
}

func (v *VaultWalk) HaveKey(door rune) bool {
	for _, k := range v.KeysFound {
		if unicode.ToUpper(k) == door {
			return true
		}
	}

	return false
}

func (v *VaultWalk) Walk(vault [][]rune) []*VaultWalk {
	walks := []*VaultWalk{}
	d := []aoc.Point{
		aoc.Point{-1, 0},
		aoc.Point{1, 0},
		aoc.Point{0, -1},
		aoc.Point{0, 1},
	}

	// p := v.Pos
	for _, p := range d {
		np := v.Pos.Add(p)
		if 0 <= np.Y && np.Y < len(vault) && 0 <= np.X && np.X < len(vault[np.Y]) {
			r := vault[np.Y][np.X]
			switch {
			case r == '.':
				walks = append(walks, NewWalk(np, v.WalkLength+1, v.KeysFound))
			case unicode.IsUpper(r) && v.HaveKey(r):
				walks = append(walks, NewWalk(np, v.WalkLength+1, v.KeysFound))
			case unicode.IsLower(r):
				keys := make([]rune, len(v.KeysFound)+1)
				copy(keys, v.KeysFound)
				keys[len(v.KeysFound)] = r
				walks = append(walks, NewWalk(np, v.WalkLength+1, keys))

			}
		}
		// if vault[]

		// }
	}
	return walks
}

type VaultPriorityQueue []*VaultWalk

func (pq VaultPriorityQueue) Len() int { return len(pq) }

func (pq VaultPriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest based on expiration number as the priority
	// The lower the expiry, the higher the priority
	return pq[i].WalkLength < pq[j].WalkLength
}

// We just implement the pre-defined function in interface of heap.

func (pq *VaultPriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1
	*pq = old[0 : n-1]

	return item
}

func (pq *VaultPriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*VaultWalk)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq VaultPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

type Day18 struct{}

func (d *Day18) Solve(scanner *bufio.Scanner) (string, string) {
	vault := [][]rune{}
	nrKeys := 0
	nrDoors := 0
	x, y := 0, 0

	var pos aoc.Point

	for scanner.Scan() {
		x = 0
		line := scanner.Text()

		vault = append(vault, []rune{})

		for _, c := range line {
			switch {
			case unicode.IsUpper(c):
				nrDoors++
			case unicode.IsLower(c):
				nrKeys++
			case c == '@':
				pos = aoc.Point{X: x, Y: y}
			default:
				vault[y] = append(vault[y], c)
				x++
			}
		}

		y++
	}

	pq := &VaultPriorityQueue{}
	pq.Push(&VaultWalk{
		Pos:        pos,
		WalkLength: 0,
		KeysFound:  []rune{},
		Index:      0})

	// for {
	// 	head := pq.Pop().(*VaultWalk)

	// }
	return "", ""
}

func CopyList(l []rune) []rune {
	c := make([]rune, len(l))
	copy(c, l)

	return c
}

func CopyMap(m map[aoc.Point]struct{}) map[aoc.Point]struct{} {
	c := map[aoc.Point]struct{}{}
	for k := range m {
		c[k] = struct{}{}
	}

	return c
}
