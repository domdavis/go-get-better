# Go Get Better

This repository contains the solutions for the exercises contained in the
`go get better` training course. The solutions are found under the relevant
branch (for example `git checkout exercise-12`, or just select the branch from
within GitHub).

## Exercise 11

Update the sequence type from Exercise 10 so when it's marshalled to JSON the 
output in the format: `{"name": ["1", "2",...]}`. An empty sequence should 
simply be `{}`.

## Hints

  * You'll need to implement `func (s Sequence) MarshalJSON() ([]byte, error)`
  * The output is basically a map

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
