# Go Get Better

This repository contains the solutions for the exercises contained in the
`go get better` training course. The solutions are found under the relevant
branch (for example `git checkout exercise-4`, or just select the branch from
within GitHub).

## Exercise 4

Turn your 'FizzBuzz' solution into one that returns an error if it's passed a
negative number. 

  * Your Function should take a number for the number of iterations
  * Your function should return an error if that number is less than 0
  * Your function should return a string containing the FizzBuzz sequence
  
> Feel free to reuse your code from exercise 3.
  
## Hints

  * The function signature will be `FizzBuzz(n int) (string, error)`

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
