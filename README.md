# Go Get Better

This repository contains the solutions for the exercises contained in the
`go get better` training course. The solutions are found under the relevant
branch (for example `git checkout exercise-13`, or just select the branch from
within GitHub).

## Exercise 13

Implement the following using channels:

````go
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
````

## Hints

  * You will need a package `fsm`
  * Use the state machine definition on the exercise slide
  * NewMachine will yield some type that has `Send` and `State` functions
  * You might want two channels, one to listen for events, one to return states

## Running The Code

To run the code from the command line:

```bash
go run solution.go
```

Tests can be run using:

```bash
go test ./...
```

Or, for more verbose output and coverage information:

```bash
go test -v -covermode=count ./...
```

A [Makefile](Makefile) is included with some basic targets (mainly `run` and
`test`) to do the above, as well as perform a check of the code using tools like
[golangci-lint][linter] for full code analysis. These are not required for the
course and can be ignored.

[linter]: https://golangci-lint.run
