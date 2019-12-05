package intcode

type InputOutput struct {
	Input  chan int
	Output chan int
}

func NewInputOutput() *InputOutput {
	return &InputOutput{
		Input:  make(chan int, 10),
		Output: make(chan int, 10),
	}
}

func (io *InputOutput) Close() {
	close(io.Output)
	close(io.Input)
}

func (io *InputOutput) Reset() {
	io.Input = make(chan int, 10)
	io.Output = make(chan int, 10)
}

func (io *InputOutput) Read() int {
	return <-io.Input
}

func (io *InputOutput) Write(v int) {
	io.Output <- v
}
