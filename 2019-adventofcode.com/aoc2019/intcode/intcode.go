package intcode

import "fmt"

var Debug bool = false

type IntCodeMachine struct {
	Mem        Memory
	initialMem Memory
	opCodes    map[int]*Instruction
	IO         *InputOutput
	Closed     bool
}

func NewIntCodeMachine(initialMem Memory) *IntCodeMachine {
	io := NewInputOutput()
	return &IntCodeMachine{
		Mem:        initialMem.Copy(),
		initialMem: initialMem,
		opCodes: map[int]*Instruction{
			1:  Addition(),
			2:  Multiplication(),
			3:  Input(io),
			4:  Output(io),
			5:  JumpIfTrue(),
			6:  JumpIfFalse(),
			7:  LessThan(),
			8:  Equals(),
			99: Halt(),
		},
		IO:     io,
		Closed: false,
	}
}

func (icm *IntCodeMachine) Reset() {
	icm.Mem = icm.initialMem.Copy()
	icm.IO.Reset()
	icm.Closed = false
}

func (icm *IntCodeMachine) Run() int {
	pointer := 0
	for pointer < len(icm.Mem) {
		i := icm.opCodes[icm.Mem[pointer]%100]
		if Debug {
			fmt.Printf("%d: %s %v\n", pointer, i.str, icm.Mem[pointer:pointer+i.opCodeLen])
		}
		err := i.f(pointer, icm.Mem)
		if err != nil {
			switch e := err.(type) {
			case *HaltError:
				icm.IO.Close()
				icm.Closed = true
				return e.code
			case *Jump:
				pointer = e.newLocation
				continue
			}
		}
		pointer += i.opCodeLen
	}
	return 0
}
