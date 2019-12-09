package intcode

type State struct {
	Ip           int // InstructionPointer
	RelativeBase int
}
