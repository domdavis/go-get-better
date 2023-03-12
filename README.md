# Go Get Better

This repository contains the solutions for the exercises contained in the
`go get better` training course. The solutions are found under the relevant
branch (for example `git checkout exercise-9`, or just select the branch from
within GitHub).

## Exercise 9

Use the code from exercise 8 and implement a 'deferred reverse` sequence
      
## Hints

  * `&` means pointer to, and `*` means get me the value pointed to
  * You'll want to change `Generator` to `func(n int) (*Sequence, error)`, and
    update `Simple` and `FizzBuzz`    
  * your defer func will want to be `defer func(i int) { ... }(i)`

  * You'll want to update `Run` to dereference the pointer to the sequence

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
