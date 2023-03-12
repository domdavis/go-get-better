# Go Get Better

This repository contains the solutions for the exercises contained in the
`go get better` training course. The solutions are found under the relevant
branch (for example `git checkout exercise-1`, or just select the branch from
within GitHub).

## Exercise 8

Use the `template` branch and your code from Exercise 7 to make all the tests
pass.
      
## Hints

  * You'll need up update your FizzBuzz generator to return "FizzBuzz" as the
    first element of the sequence

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
