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
	str       string
	opCodeLen int
	f         func(state *State, mem Memory) error
}

func Addition() *Instruction {
	return &Instruction{
		opCode:    1,
		str:       "add",
		opCodeLen: 4,
		f: func(state *State, mem Memory) error {
			modeOp := mem[state.IP]
			l := mem.Get(state.IP+1, modeOp/100%10, state.RelativeBase)
			r := mem.Get(state.IP+2, modeOp/1000%10, state.RelativeBase)
			mem.Set(state.IP+3, l+r, modeOp/10000%10, state.RelativeBase)
			return nil
		},
	}
}

func Multiplication() *Instruction {
	return &Instruction{
		opCode:    2,
		str:       "mul",
		opCodeLen: 4,
		f: func(state *State, mem Memory) error {
			modeOp := mem[state.IP]
			l := mem.Get(state.IP+1, modeOp/100%10, state.RelativeBase)
			r := mem.Get(state.IP+2, modeOp/1000%10, state.RelativeBase)
			mem.Set(state.IP+3, l*r, modeOp/10000%10, state.RelativeBase)
			return nil
		},
	}
}

func Halt() *Instruction {
	return &Instruction{
		opCode:    99,
		str:       "halt",
		opCodeLen: 1,
		f: func(state *State, mem Memory) error {
			return &HaltError{mem[0]}
		},
	}
}

func Input(io *InputOutput) *Instruction {
	return &Instruction{
		opCode:    3,
		str:       "inp",
		opCodeLen: 2,
		f: func(state *State, mem Memory) error {
			modeOp := mem[state.IP]
			mem.Set(state.IP+1, io.Read(), modeOp/100%10, state.RelativeBase)
			return nil
		},
	}
}

func Output(io *InputOutput) *Instruction {
	return &Instruction{
		opCode:    4,
		str:       "out",
		opCodeLen: 2,
		f: func(state *State, mem Memory) error {
			modeOp := mem[state.IP]
			io.Write(mem.Get(state.IP+1, modeOp/100%10, state.RelativeBase))
			return nil
		},
	}
}

func JumpIfTrue() *Instruction {
	return &Instruction{
		opCode:    5,
		str:       "jt",
		opCodeLen: 3,
		f: func(state *State, mem Memory) error {
			modeOp := mem[state.IP]
			jump := mem.Get(state.IP+1, modeOp/100%10, state.RelativeBase)
			if jump != 0 {
				return &Jump{mem.Get(state.IP+2, modeOp/1000%10, state.RelativeBase)}
			}
			return nil
		},
	}
}

func JumpIfFalse() *Instruction {
	return &Instruction{
		opCode:    5,
		str:       "jf",
		opCodeLen: 3,
		f: func(state *State, mem Memory) error {
			modeOp := mem[state.IP]
			jump := mem.Get(state.IP+1, modeOp/100%10, state.RelativeBase)
			if jump == 0 {
				return &Jump{mem.Get(state.IP+2, modeOp/1000%10, state.RelativeBase)}
			}
			return nil
		},
	}
}

func LessThan() *Instruction {
	return &Instruction{
		opCode:    6,
		str:       "lt",
		opCodeLen: 4,
		f: func(state *State, mem Memory) error {
			modeOp := mem[state.IP]
			l := mem.Get(state.IP+1, modeOp/100%10, state.RelativeBase)
			r := mem.Get(state.IP+2, modeOp/1000%10, state.RelativeBase)
			if l < r {
				mem.Set(state.IP+3, 1, modeOp/10000%10, state.RelativeBase)
			} else {
				mem.Set(state.IP+3, 0, modeOp/10000%10, state.RelativeBase)
			}
			return nil
		},
	}
}
func Equals() *Instruction {
	return &Instruction{
		opCode:    8,
		str:       "eq",
		opCodeLen: 4,
		f: func(state *State, mem Memory) error {
			modeOp := mem[state.IP]
			l := mem.Get(state.IP+1, modeOp/100%10, state.RelativeBase)
			r := mem.Get(state.IP+2, modeOp/1000%10, state.RelativeBase)
			if l == r {
				mem.Set(state.IP+3, 1, modeOp/10000%10, state.RelativeBase)
			} else {
				mem.Set(state.IP+3, 0, modeOp/10000%10, state.RelativeBase)
			}
			return nil
		},
	}
}

func SetRelativeBase() *Instruction {
	return &Instruction{
		opCode:    9,
		str:       "roff",
		opCodeLen: 2,
		f: func(state *State, mem Memory) error {
			modeOp := mem[state.IP]
			l := mem.Get(state.IP+1, modeOp/100%10, state.RelativeBase)
			state.RelativeBase += l
			return nil
		},
	}
}
