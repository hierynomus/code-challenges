package intcode

import "fmt"

type Memory []int

const (
	Position  = 0
	Immediate = 1
	Relative  = 2
)

func (m Memory) Get(address int, mode int, relativeBase int) int {
	if mode == Position {
		return m[m[address]]
	} else if mode == Immediate {
		return m[address]
	} else if mode == Relative {
		return m[m[address]+relativeBase]
	}
	panic(fmt.Errorf("Invalid mode %d", mode))
}

func (m Memory) Set(address int, val int, mode int, relativeBase int) {
	if mode == Position {
		m[m[address]] = val
		return
	} else if mode == Relative {
		m[m[address]+relativeBase] = val
		return
	}
	panic(fmt.Errorf("Invalid mode %d", mode))
}

func (m Memory) Copy() Memory {
	newMem := make([]int, len(m)+100000)
	copy(newMem, m)
	return newMem
}
