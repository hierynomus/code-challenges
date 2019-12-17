package intcode

import "fmt"

var Debug bool = false //nolint:gochecknoglobals

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
			9:  SetRelativeBase(),
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
	state := &State{0, 0}
	for state.IP < len(icm.Mem) {
		opc := icm.Mem[state.IP] % 100
		i, ok := icm.opCodes[opc]

		if !ok {
			panic(fmt.Errorf("unknown opcode %d", opc))
		}

		if Debug {
			fmt.Printf("%d: %s %v\n", state.IP, i.str, icm.Mem[state.IP:state.IP+i.opCodeLen])
		}

		err := i.f(state, icm.Mem)
		if err != nil {
			switch e := err.(type) {
			case *HaltError:
				icm.IO.Close()
				icm.Closed = true

				return e.code
			case *Jump:
				state.IP = e.newLocation
				continue
			}
		}

		state.IP += i.opCodeLen
	}

	return 0
}
