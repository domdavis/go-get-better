package fsm_test

import (
	"fmt"

	"github.com/domdavis/solution/fsm"
)

func ExampleNewMachine() {
	sm := fsm.NewMachine()
	sm.Send(1)
	sm.Send(0)
	fmt.Println(sm.State())

	// Output:
	// C
}
