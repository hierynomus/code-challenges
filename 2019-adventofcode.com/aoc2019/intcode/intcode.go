package intcode

type IntCodeMachine struct {
	Mem        Memory
	initialMem Memory
	opCodes    map[int]*Instruction
	IO         *InputOutput
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
		IO: io,
	}
}

func (icm *IntCodeMachine) Reset() {
	icm.Mem = icm.initialMem.Copy()
	icm.IO.Reset()
}

func (icm *IntCodeMachine) Run() int {
	pointer := 0
	for pointer < len(icm.Mem) {
		i := icm.opCodes[icm.Mem[pointer]%100]
		err := i.f(pointer, icm.Mem)
		if err != nil {
			switch e := err.(type) {
			case *HaltError:
				icm.IO.Close()
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
