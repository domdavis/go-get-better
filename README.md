# Go Get Better

This repository contains the solutions for the exercises contained in the
`go get better` training course. The solutions are found under the relevant
branch (for example `git checkout exercise-10`, or just select the branch from
within GitHub).

## Exercise 10

Turn the Run function from exercise 9 into a type function on Generator. You'll
need to update the example and the tests to call the function correctly.
      
## Hints

  * The signature becomes: `func (g Generator) Run(n int) (Sequence, error)`
  * The tests use something like: `g := Generator(Simple); r, err := g.Run(0)`

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
