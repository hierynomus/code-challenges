package aoc2019

import (
	"bufio"
	"fmt"
	"strconv"
	"sync"

	"github.com/hierynomus/code-challenges/adventofcode/internal/aoc2019/intcode"
	"github.com/hierynomus/code-challenges/adventofcode/pkg/aoc"
)

type Computer struct {
	machine *intcode.IntCodeMachine
	input   *PacketInputChannel
	output  *PacketOutputChannel
}

func NewComputer(program []int, router *Router) *Computer {
	i := NewInputChannel()
	o := NewOutputChannel(router)
	machine := intcode.NewMachine(program, i, o)

	return &Computer{machine: machine, input: i, output: o}
}

type Router struct {
	computers []*Computer
	nat       Nat
}

func NewRouter(c []*Computer, n Nat) *Router {
	return &Router{computers: c, nat: n}
}

func (r *Router) Route(p *Packet) {
	d := p.destination
	if d < len(r.computers) {
		r.computers[d].input.WritePacket(p)
	} else {
		r.nat.ReceivePacket(p)
	}
}

type Nat interface {
	ReceivePacket(*Packet)
}

type BreakingNat struct {
	received       chan int
	currentPacket  *Packet
	deliverdPart1  bool
	lastDeliveredY int
	lock           *sync.Mutex
}

func (n *BreakingNat) ReceivePacket(p *Packet) {
	// n.lock.Lock()
	// defer n.lock.Unlock()

	n.currentPacket = p

	if !n.deliverdPart1 {
		n.received <- p.y
		n.deliverdPart1 = true
	}
}

func (n *BreakingNat) Monitor(computers []*Computer) {
	for {
		// n.lock.Lock()

		allIdle := true

		for _, c := range computers {
			if !c.input.idle {
				allIdle = false
				break
			}
		}

		if allIdle && n.currentPacket != nil {
			computers[0].input.WritePacket(n.currentPacket)
			fmt.Println(n.currentPacket.y)

			if n.currentPacket.y == n.lastDeliveredY {
				n.received <- n.lastDeliveredY
			}

			n.lastDeliveredY = n.currentPacket.y
		}
		// n.lock.Unlock()
		// time.Sleep(1000)
	}
}

type Packet struct {
	destination int
	x, y        int
}

type PacketInputChannel struct {
	queue []int
	lock  *sync.Mutex
	idle  bool
}

func NewInputChannel() *PacketInputChannel {
	return &PacketInputChannel{queue: []int{}, lock: &sync.Mutex{}}
}

var _ intcode.IOChannel = (*PacketInputChannel)(nil) // compile-time check

func (pc *PacketInputChannel) Close() {
	pc.lock.Lock()
	defer pc.lock.Unlock()
	pc.queue = []int{}
	pc.idle = false
}

func (pc *PacketInputChannel) Reset() {
	pc.lock.Lock()
	defer pc.lock.Unlock()
	pc.queue = []int{}
	pc.idle = false
}

func (pc *PacketInputChannel) Read() int {
	pc.lock.Lock()
	defer pc.lock.Unlock()

	retval := -1
	if len(pc.queue) > 0 {
		retval = pc.queue[0]
		pc.queue = pc.queue[1:]
	} else {
		pc.idle = true
	}

	return retval
}

func (pc *PacketInputChannel) Write(i int) {
	pc.lock.Lock()
	defer pc.lock.Unlock()
	pc.queue = append(pc.queue, i)
	pc.idle = false
}

func (pc *PacketInputChannel) WritePacket(p *Packet) {
	pc.lock.Lock()
	defer pc.lock.Unlock()
	pc.queue = append(pc.queue, p.x)
	pc.queue = append(pc.queue, p.y)
	pc.idle = false
}

type PacketOutputChannel struct {
	partial []int
	router  *Router
	lock    *sync.Mutex
}

func NewOutputChannel(router *Router) *PacketOutputChannel {
	return &PacketOutputChannel{
		partial: []int{},
		router:  router,
		lock:    &sync.Mutex{},
	}
}

var _ intcode.IOChannel = (*PacketOutputChannel)(nil)

func (pc *PacketOutputChannel) Read() int {
	panic("BooM")
}

func (pc *PacketOutputChannel) Write(i int) {
	pc.lock.Lock()
	defer pc.lock.Unlock()

	pc.partial = append(pc.partial, i)
	if len(pc.partial) == 3 {
		p := &Packet{destination: pc.partial[0], x: pc.partial[1], y: pc.partial[2]}
		pc.partial = []int{}

		go pc.router.Route(p)
	}
}

func (pc *PacketOutputChannel) Close() {
	pc.lock.Lock()
	defer pc.lock.Unlock()
	pc.partial = []int{}
}

func (pc *PacketOutputChannel) Reset() {
	pc.lock.Lock()
	defer pc.lock.Unlock()
	pc.partial = []int{}
}

func Day23(scanner *bufio.Scanner) (string, string) {
	if !scanner.Scan() {
		panic(fmt.Errorf("boom"))
	}

	program := aoc.AsIntArray(scanner.Text())

	computers := make([]*Computer, 50)
	nat1 := &BreakingNat{received: make(chan int, 1), lock: &sync.Mutex{}}
	router := &Router{computers: computers, nat: nat1}

	for i := 0; i < 50; i++ {
		computers[i] = NewComputer(program, router)
		computers[i].input.Write(i)
	}

	for _, c := range computers {
		go c.machine.Run()
	}

	go nat1.Monitor(computers)
	part1 := <-nat1.received

	part2 := <-nat1.received

	return strconv.Itoa(part1), strconv.Itoa(part2)
}
