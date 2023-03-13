package fsm

// State in the Machine. A State has a name, and pointers to the State instances
// it will transition to on an event.
type State struct {
	Name string
	Zero *State
	One  *State
}

// Listen on the msg channel for a single transition event. In the case of a 0
// or 1 on the msg channel the relevant target state will be put on the update
// channel. Any other update message will result in the current state being
// put on the update channel.
func (s *State) Listen(msg chan int, update chan *State) {
	switch <-msg {
	case 0:
		update <- s.Zero
	case 1:
		update <- s.One
	default:
		update <- s
	}
}
