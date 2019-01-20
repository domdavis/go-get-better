package main

// Implement stateMachine.
type state struct {
	name string
	zero *state
	one  *state
}

func (s *state) listen(msg chan int, update chan *state) {
	switch <-msg {
	case 0:
		update <- s.zero
	case 1:
		update <- s.one
	}
}

type stateMachine struct {
	states  []*state
	current *state
	msg     chan int
	update  chan *state
}

func (s *stateMachine) send(i int) {
	s.msg <- i
	s.current = <-s.update
	go s.current.listen(s.msg, s.update)
}

func (s *stateMachine) state() string {
	return s.current.name
}

func newStateMachine() *stateMachine {
	a := &state{name: "A"}
	b := &state{name: "B"}
	c := &state{name: "C"}

	a.zero = a
	a.one = b
	b.zero = c
	b.one = a
	c.zero = b
	c.one = a

	msg := make(chan int)
	update := make(chan *state)

	go a.listen(msg, update)

	s := &stateMachine{
		states:  []*state{a, b, c},
		current: a,
		msg:     msg,
		update:  update,
	}

	return s
}

func main() {
	sm := newStateMachine()
	sm.send(1)          // "state A + 1 => state B"
	sm.send(0)          // "state B + 0 => state C"
	println(sm.state()) // "state C"
}
