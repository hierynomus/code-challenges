package aoc

type StringStack []string

func (s *StringStack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *StringStack) Push(str string) {
	*s = append(*s, str)
}

func (s *StringStack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

type RuneStack []rune

func (s *RuneStack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *RuneStack) Push(r rune) {
	*s = append(*s, r)
}

func (s *RuneStack) Pop() (rune, bool) {
	if s.IsEmpty() {
		return rune(0), false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func (s *RuneStack) Peek() (rune, bool) {
	if s.IsEmpty() {
		return rune(0), false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		return element, true
	}
}
