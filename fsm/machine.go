package fsm

// Machine holds a definition and implementation of a finite state machine.
// Machine should be obtained using NewMachine
type Machine struct {
	States  []*State
	Current *State
	Msg     chan int
	Update  chan *State
}

// NewMachine returns a new Machine configured for Exercise 13.
func NewMachine() *Machine {
	a := &State{Name: "A"}
	b := &State{Name: "B"}
	c := &State{Name: "C"}

	a.Zero = a
	a.One = b
	b.Zero = c
	b.One = a
	c.Zero = b
	c.One = a

	msg := make(chan int)
	update := make(chan *State)

	go a.Listen(msg, update)

	return &Machine{
		States:  []*State{a, b, c},
		Current: a,
		Msg:     msg,
		Update:  update,
	}
}

// Send the given event to the Machine. The current State will be updated based
// on the definition of the Machine.
func (s *Machine) Send(i int) {
	s.Msg <- i
	s.Current = <-s.Update
	go s.Current.Listen(s.Msg, s.Update)
}

// State the Machine is currently in. Only the State name is returned.
func (s *Machine) State() string {
	return s.Current.Name
}
