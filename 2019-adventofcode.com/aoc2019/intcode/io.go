package intcode

type IOChannel interface {
	Read() int
	Write(i int)
	Close()
	Reset()
}

type BlockingChannel struct {
	c chan int
}

func (io *BlockingChannel) Read() int {
	return <-io.c
}

func (io *BlockingChannel) Write(i int) {
	io.c <- i
}

func (io *BlockingChannel) Close() {
	close(io.c)
}

func (io *BlockingChannel) Reset() {
	io.c = make(chan int, 10)
}

func NewBlockingChannel() *BlockingChannel {
	return &BlockingChannel{
		c: make(chan int, 10),
	}
}
