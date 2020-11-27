package intcode

import "fmt"

var Debug bool = false //nolint:gochecknoglobals

type IntCodeMachine struct {
	Mem        Memory
	initialMem Memory
	opCodes    map[int]*Instruction
	Input      IOChannel
	Output     IOChannel
	Closed     bool
	ClosedCh   chan struct{}
	State      *State
}

func NewIntCodeMachine(initialMem Memory) *IntCodeMachine {
	return NewMachine(initialMem, NewBlockingChannel(), NewBlockingChannel())
}

func NewMachine(initialMem Memory, input IOChannel, output IOChannel) *IntCodeMachine {
	return &IntCodeMachine{
		Mem:        initialMem.Copy(),
		initialMem: initialMem,
		opCodes: map[int]*Instruction{
			1:  Addition(),
			2:  Multiplication(),
			3:  Input(input),
			4:  Output(output),
			5:  JumpIfTrue(),
			6:  JumpIfFalse(),
			7:  LessThan(),
			8:  Equals(),
			9:  SetRelativeBase(),
			99: Halt(),
		},
		Input:    input,
		Output:   output,
		Closed:   false,
		ClosedCh: make(chan struct{}, 1),
		State:    &State{0, 0},
	}
}

func (icm *IntCodeMachine) Reset() {
	icm.Mem = icm.initialMem.Copy()
	icm.State = &State{0, 0}
	icm.Input.Reset()
	icm.Output.Reset()
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
				icm.Input.Close()
				icm.Output.Close()
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
