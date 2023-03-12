# Go Get Better

This repository contains the solutions for the exercises contained in the
`go get better` training course. The solutions are found under the relevant
branch (for example `git checkout exercise-1`, or just select the branch from
within GitHub).

## Exercise 1

Create a program that will perform 'FizzBuzz'. The rules for this are:

  * For the numbers 1 to 100
  * If a number is divisible by 3 output Fizz
  * If a number is divisible by 5 output Buzz
  * If a number is divisible by 3 and 5 output FizzBuzz
  * Otherwise, output the number
  
## Alternative solutions

A couple of _joke_ alternative solutions are also provided. These are runnable
tests and highlight how to write testable examples.

## Running The Code

To run the code from the command line:

```bash
go run solution.go
```

Tests can be run using:

```bash
go test -v -covermode=count ./...
```

A [Makefile](Makefile) is included with some basic targets (mainly `run` and 
`test`) to do the above, as well as perform a check of the code using tools like
[golangci-lint][linter] for full code analysis. These are not required for the
course and can be ignored.

[playground]: https://play.golang.org/
[linter]: https://golangci-lint.run
