package intcode

import "fmt"

type HaltError struct {
	code int
}

func (e *HaltError) Error() string {
	return fmt.Sprintf("Halt %d", e.code)
}

type Instruction struct {
	opCode    int
	opCodeLen int
	f         func(ptr int, mem []int) ([]int, error)
}

func Addition() *Instruction {
	return &Instruction{
		opCode:    1,
		opCodeLen: 4,
		f: func(ptr int, mem []int) ([]int, error) {
			l, r := mem[ptr+1], mem[ptr+2]
			out := mem[ptr+3]
			mem[out] = mem[l] + mem[r]
			return mem, nil
		},
	}
}

func Multiplication() *Instruction {
	return &Instruction{
		opCode:    2,
		opCodeLen: 4,
		f: func(ptr int, mem []int) ([]int, error) {
			l, r := mem[ptr+1], mem[ptr+2]
			out := mem[ptr+3]
			mem[out] = mem[l] * mem[r]
			return mem, nil
		},
	}
}

func Halt() *Instruction {
	return &Instruction{
		opCode:    99,
		opCodeLen: 1,
		f: func(ptr int, mem []int) ([]int, error) {
			return mem, &HaltError{mem[0]}
		},
	}
}
