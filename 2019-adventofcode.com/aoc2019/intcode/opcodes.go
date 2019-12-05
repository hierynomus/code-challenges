package intcode

import "fmt"

type HaltError struct {
	code int
}

type Jump struct {
	newLocation int
}

func (e *HaltError) Error() string {
	return fmt.Sprintf("Halt %d", e.code)
}

func (e *Jump) Error() string {
	return fmt.Sprintf("Jump %d", e.newLocation)
}

type Instruction struct {
	opCode    int
	opCodeLen int
	f         func(ptr int, mem Memory) error
}

func Addition() *Instruction {
	return &Instruction{
		opCode:    1,
		opCodeLen: 4,
		f: func(ptr int, mem Memory) error {
			modeOp := mem[ptr]
			l := mem.Get(ptr+1, modeOp/100%10)
			r := mem.Get(ptr+2, modeOp/1000%10)
			mem.Set(ptr+3, l+r)
			return nil
		},
	}
}

func Multiplication() *Instruction {
	return &Instruction{
		opCode:    2,
		opCodeLen: 4,
		f: func(ptr int, mem Memory) error {
			modeOp := mem[ptr]
			l := mem.Get(ptr+1, modeOp/100%10)
			r := mem.Get(ptr+2, modeOp/1000%10)
			mem.Set(ptr+3, l*r)
			return nil
		},
	}
}

func Halt() *Instruction {
	return &Instruction{
		opCode:    99,
		opCodeLen: 1,
		f: func(ptr int, mem Memory) error {
			return &HaltError{mem[0]}
		},
	}
}

func Input(io *InputOutput) *Instruction {
	return &Instruction{
		opCode:    3,
		opCodeLen: 2,
		f: func(ptr int, mem Memory) error {
			mem.Set(ptr+1, io.Read())
			return nil
		},
	}
}

func Output(io *InputOutput) *Instruction {
	return &Instruction{
		opCode:    4,
		opCodeLen: 2,
		f: func(ptr int, mem Memory) error {
			modeOp := mem[ptr]
			io.Write(mem.Get(ptr+1, modeOp/100%10))
			return nil
		},
	}
}

func JumpIfTrue() *Instruction {
	return &Instruction{
		opCode:    5,
		opCodeLen: 3,
		f: func(ptr int, mem Memory) error {
			modeOp := mem[ptr]
			jump := mem.Get(ptr+1, modeOp/100%10)
			if jump != 0 {
				return &Jump{mem.Get(ptr+2, modeOp/1000%10)}
			}
			return nil
		},
	}
}

func JumpIfFalse() *Instruction {
	return &Instruction{
		opCode:    5,
		opCodeLen: 3,
		f: func(ptr int, mem Memory) error {
			modeOp := mem[ptr]
			jump := mem.Get(ptr+1, modeOp/100%10)
			if jump == 0 {
				return &Jump{mem.Get(ptr+2, modeOp/1000%10)}
			}
			return nil
		},
	}
}

func LessThan() *Instruction {
	return &Instruction{
		opCode:    6,
		opCodeLen: 4,
		f: func(ptr int, mem Memory) error {
			modeOp := mem[ptr]
			l := mem.Get(ptr+1, modeOp/100%10)
			r := mem.Get(ptr+2, modeOp/1000%10)
			if l < r {
				mem.Set(ptr+3, 1)
			} else {
				mem.Set(ptr+3, 0)
			}
			return nil
		},
	}
}
func Equals() *Instruction {
	return &Instruction{
		opCode:    8,
		opCodeLen: 4,
		f: func(ptr int, mem Memory) error {
			modeOp := mem[ptr]
			l := mem.Get(ptr+1, modeOp/100%10)
			r := mem.Get(ptr+2, modeOp/1000%10)
			if l == r {
				mem.Set(ptr+3, 1)
			} else {
				mem.Set(ptr+3, 0)
			}
			return nil
		},
	}
}
