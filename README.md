# Go Get Better

This repository contains the solutions for the exercises contained in the
`go get better` training course. The solutions are found under the relevant
branch (for example `git checkout exercise-1`, or just select the branch from
within GitHub).

## Exercise 13

Implement the following using channels

````go
package main

// Implement stateMachine.
type stateMachine struct{}

func main() {
    sm := newStateMachine()
    sm.send(1)          // "state A + 1 => state B"
    sm.send(0)          // "state B + 0 => state C"
    println(sm.state()) // "state C"
}
````

## Hints

  * Use the state machine definition on the exercise slide
  * Each state should have a listener which listens for a single message
  * When the state changes a new listener should be launched
