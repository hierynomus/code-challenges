package intcode

type Memory []int

func (m Memory) Get(address int, mode int) int {
	if mode == 0 {
		return m[m[address]]
	} else {
		return m[address]
	}
}

func (m Memory) Set(address int, val int) {
	m[m[address]] = val
}

func (m Memory) Copy() Memory {
	newMem := make([]int, len(m))
	copy(newMem, m)
	return newMem
}
