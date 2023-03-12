# Go Get Better

This repository contains the solutions for the exercises contained in the
`go get better` training course. The solutions are found under the relevant
branch (for example `git checkout exercise-1`, or just select the branch from
within GitHub).

## Exercise 11

Using the code from Exercise 9 (not 10), change the `Run` function to be:

```
// Run the sequence g from 1 to n.
func Run(g Generator, n int) (Sequence, error) {
	p, err := g.Run(n)
	return *p, err
}
```

The `Generator` interface is defined as: 

```
type Generator interface {
	Run(n int) (*Sequence, error)
}
```

## Hints

  * You can use `type name struct{}` to create a type that does not need to 
    store data
  * `func (name) Run() {}` can be used if the type function does not need
    access to the type

## Running The Tests

To run the tests from the command line:

```bash
go test ./...
```

Or, for more verbose output and coverage information:

```bash
go test -v -covermode=count ./...
```

A [Makefile](Makefile) is included with some basic targets for testing, as well
as perform a check of the code using tools like [golangci-lint][linter] for full
code analysis. These are not required for the course and can be ignored.

[linter]: https://golangci-lint.run
