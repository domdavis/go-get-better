# Go Get Better

This repository contains the solutions for the exercises contained in the
`go get better` training course. The solutions are found under the relevant
branch (for example `git checkout exercise-14`, or just select the branch from
within GitHub).

## Exercise 14

Complete the implementation for the branch `pipelines`.

## Hints

* You will need `fizz`, `buzz` and `number` steps.
* Look at `pipe` which is used by the `Simple` generator.

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
