package aoc2019

import (
	"bufio"
	"fmt"
	"strconv"
	"unicode"

	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type VaultWalk struct {
	Pos        aoc.Point
	PrevPos    aoc.Point
	WalkLength int
	KeysFound  []rune
	Index      int
}

func NewWalk(pos aoc.Point, prevPos aoc.Point, length int, keys []rune) *VaultWalk {
	// kf := make([]rune, len(keys))
	// copy(kf, keys)
	return &VaultWalk{
		Pos:        pos,
		PrevPos:    prevPos,
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

	// p := v.Pos
	for _, np := range v.Pos.Neighbours4() {
		if 0 <= np.Y && np.Y < len(vault) && 0 <= np.X && np.X < len(vault[np.Y]) {
			if v.PrevPos == np {
				continue
			}

			r := vault[np.Y][np.X]

			switch {
			case r == '.':
				walks = append(walks, NewWalk(np, v.Pos, v.WalkLength+1, v.KeysFound))
			case unicode.IsUpper(r) && v.HaveKey(r):
				walks = append(walks, NewWalk(np, v.Pos, v.WalkLength+1, v.KeysFound))
			case unicode.IsLower(r):
				keys := make([]rune, len(v.KeysFound)+1)
				copy(keys, v.KeysFound)
				keys[len(v.KeysFound)] = r
				walks = append(walks, NewWalk(np, v.Pos, v.WalkLength+1, keys))
			}
		}
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

func Day18(scanner *bufio.Scanner) (string, string) {
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

	head := &VaultWalk{
		Pos:        pos,
		PrevPos:    pos,
		WalkLength: 0,
		KeysFound:  []rune{},
		Index:      0}

	for len(head.KeysFound) < nrKeys {
		fmt.Printf("%v\n", head)

		for _, walk := range head.Walk(vault) {
			pq.Push(walk)
		}

		head = pq.Pop().(*VaultWalk)
	}

	return strconv.Itoa(head.WalkLength), ""
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
