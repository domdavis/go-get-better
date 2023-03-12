# Go Get Better

This repository contains the solutions for the exercises contained in the
`go get better` training course. The solutions are found under the relevant
branch (for example `git checkout exercise-3`, or just select the branch from
within GitHub).

## Exercise 3

Turn your 'FizzBuzz' solution into a function in a package. Test the function. 

  * Your Function should take a number for the number of iterations
  * Your function should return a string containing the FizzBuzz sequence

  
> Feel free to reuse your code from previous exercises.
  
## Hints

  * Your test file needs the suffix `_test.go`
  * Your tests should live in the `_test` package

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
