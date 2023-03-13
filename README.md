# Go Get Better

This repository contains the solutions for the exercises contained in the
`go get better` training course. The solutions are found under the relevant
branch (for example `git checkout exercise-1`, or just select the branch from
within GitHub).

## Exercise 15

Write an implementation off FizzBuzz that will respond to a HTTP request to 
`fizzbuzz/n`

## Hints

  * your handler will have the signature 
    `func(w http.ResponseWriter, r *http.Request)`
  * You will probably want to use regular expressions to get the url parameter

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
