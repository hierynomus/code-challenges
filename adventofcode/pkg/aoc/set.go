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

type RuneSet map[rune]struct{}

func NewRuneSet(rs []rune) RuneSet {
	s := RuneSet{}
	s.Adds(rs)
	return s
}

func (s RuneSet) Add(i rune) {
	s[i] = exists
}

func (s RuneSet) Adds(rs []rune) {
	for _, r := range rs {
		s[r] = exists
	}
}

func (s RuneSet) Intersect(rs []rune) {
	for c := range s {
		f := false
		for _, r := range rs {
			f = f || c == r
		}

		if !f {
			s.Delete(c)
		}
	}
}

func (s RuneSet) Contains(i rune) bool {
	_, ok := s[i]
	return ok
}

func (s RuneSet) Delete(i rune) {
	delete(s, i)
}

func (s RuneSet) Deletes(rs []rune) {
	for _, r := range rs {
		delete(s, r)
	}
}

func (s RuneSet) Copy() RuneSet {
	n := RuneSet{}
	for k := range s {
		n[k] = exists
	}

	return n
}

func (s RuneSet) Min() rune {
	ks := s.Keys()
	if len(ks) < 1 {
		return 0
	}

	m := ks[0]
	for i := 1; i < len(ks); i++ {
		if ks[i] < m {
			m = ks[i]
		}
	}

	return m
}

func (s RuneSet) Keys() []rune {
	ks := []rune{}
	for k := range s {
		ks = append(ks, k)
	}

	return ks
}
