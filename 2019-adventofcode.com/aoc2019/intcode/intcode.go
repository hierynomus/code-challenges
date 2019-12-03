package intcode

type IntCodeMachine struct {
	mem     []int
	opCodes map[int]*Instruction
}

func NewIntCodeMachine(initialMem []int) *IntCodeMachine {
	return &IntCodeMachine{
		mem: initialMem,
		opCodes: map[int]*Instruction{
			1:  Addition(),
			2:  Multiplication(),
			99: Halt(),
		},
	}
}

func (icm *IntCodeMachine) SetNounVerb(noun int, verb int) {
	icm.mem[1] = noun
	icm.mem[2] = verb
}

func (icm *IntCodeMachine) Run() int {
	mem := make([]int, len(icm.mem))
	copy(mem, icm.mem)
	pointer := 0
	for pointer < len(mem) {
		i := icm.opCodes[mem[pointer]]
		newMem, err := i.f(pointer, mem)
		if err != nil {
			return err.(*HaltError).code
		}
		mem = newMem
		pointer += i.opCodeLen
	}
	return 0
}
