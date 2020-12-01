package aoc

var exists = struct{}{}

type IntSet map[int]struct{}

func (s IntSet) Add(i int) {
	s[i] = exists
}

func (s IntSet) Contains(i int) bool {
	_, ok := s[i]
	return ok
}

func (s IntSet) Delete(i int) {
	delete(s, i)
}
