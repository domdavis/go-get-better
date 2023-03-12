# Go Get Better

This repository contains the solutions for the exercises contained in the
`go get better` training course. The solutions are found under the relevant
branch (for example `git checkout exercise-2`, or just select the branch from
within GitHub).

## Exercise 2

Convert your solution from Exercise 1 into a function with the signature:

```
func FizzBuzz(int n)
```

The function should perform the following actions:

  * For the numbers 1 to n
  * If a number is divisible by 3 output Fizz
  * If a number is divisible by 5 output Buzz
  * If a number is divisible by 3 and 5 output FizzBuzz
  * Otherwise, output the number

Write an Example style test to check the code.
  
## Hints

  * The test will be in a file with an `_test.go` suffix
  * The function signature for the test is `ExampleFizzBuzz()`
  * The special comment `// Output:` is used to check what is sent to `STDOUT`

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
