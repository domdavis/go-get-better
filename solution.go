package main

import (
	"fmt"

	"github.com/domdavis/solution/fsm"
)

func main() {
	sm := fsm.NewMachine()
	sm.Send(1)              // "state A + 1 => state B"
	sm.Send(0)              // "state B + 0 => state C"
	fmt.Println(sm.State()) // "state C"
}
