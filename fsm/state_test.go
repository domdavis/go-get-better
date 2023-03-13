package fsm_test

import (
	"fmt"
	"testing"

	"github.com/domdavis/solution/fsm"
	"github.com/stretchr/testify/assert"
)

func ExampleState_Listen() {
	initial := &fsm.State{Name: "Initial", Zero: &fsm.State{Name: "Zero"}}

	msg := make(chan int)
	update := make(chan *fsm.State)

	go initial.Listen(msg, update)

	msg <- 0
	state := <-update

	fmt.Println(state.Name)

	// Output:
	// Zero
}

func TestState_Listen(t *testing.T) {
	t.Run("A message of 0 yields the zero state", func(t *testing.T) {
		t.Parallel()

		initial := &fsm.State{
			Name: "Initial",
			Zero: &fsm.State{Name: "Zero"},
			One:  &fsm.State{Name: "One"},
		}

		msg := make(chan int)
		update := make(chan *fsm.State)

		go initial.Listen(msg, update)

		msg <- 0
		state := <-update

		assert.Equal(t, "Zero", state.Name)
	})

	t.Run("A message of 1 yields the one state", func(t *testing.T) {
		t.Parallel()

		initial := &fsm.State{
			Name: "Initial",
			Zero: &fsm.State{Name: "Zero"},
			One:  &fsm.State{Name: "One"},
		}

		msg := make(chan int)
		update := make(chan *fsm.State)

		go initial.Listen(msg, update)

		msg <- 1
		state := <-update

		assert.Equal(t, "One", state.Name)
	})

	t.Run("An invalid message yields the current state", func(t *testing.T) {
		t.Parallel()

		initial := &fsm.State{
			Name: "Initial",
			Zero: &fsm.State{Name: "Zero"},
			One:  &fsm.State{Name: "One"},
		}

		msg := make(chan int)
		update := make(chan *fsm.State)

		go initial.Listen(msg, update)

		msg <- 2
		state := <-update

		assert.Equal(t, "Initial", state.Name)
	})

	t.Run("The state only listens for a single message", func(t *testing.T) {
		t.Parallel()

		initial := &fsm.State{
			Name: "Initial",
			Zero: &fsm.State{Name: "Zero"},
			One:  &fsm.State{Name: "One"},
		}

		msg := make(chan int)
		update := make(chan *fsm.State)

		go initial.Listen(msg, update)

		msg <- 0
		<-update

		select {
		case msg <- 1:
			t.Error("state should not be listening")
		default:
			// All is good.
		}
	})
}
