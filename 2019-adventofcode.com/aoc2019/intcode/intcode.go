package intcode

import "fmt"

var Debug bool = false //nolint:gochecknoglobals

type IntCodeMachine struct { //nolint:golint
	Mem        Memory
	initialMem Memory
	opCodes    map[int]*Instruction
	IO         *InputOutput
	Closed     bool
	ClosedCh   chan struct{}
	State      *State
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
		IO:       io,
		Closed:   false,
		ClosedCh: make(chan struct{}, 1),
		State:    &State{0, 0},
	}
}

func (icm *IntCodeMachine) Reset() {
	icm.Mem = icm.initialMem.Copy()
	icm.State = &State{0, 0}
	icm.IO.Reset()
	icm.Closed = false
}

func (icm *IntCodeMachine) Run() int {
	for icm.State.IP < len(icm.Mem) {
		opc := icm.Mem[icm.State.IP] % 100
		i, ok := icm.opCodes[opc]

		if !ok {
			panic(fmt.Errorf("unknown opcode %d", opc))
		}

		if Debug {
			fmt.Printf("%d: %s %v\n", icm.State.IP, i.str, icm.Mem[icm.State.IP:icm.State.IP+i.opCodeLen])
		}

		err := i.f(icm.State, icm.Mem)
		if err != nil {
			switch e := err.(type) {
			case *HaltError:
				icm.IO.Close()
				icm.ClosedCh <- struct{}{}
				icm.Closed = true

				return e.code
			case *Jump:
				icm.State.IP = e.newLocation
				continue
			}
		}

		icm.State.IP += i.opCodeLen
	}

	return 0
}
