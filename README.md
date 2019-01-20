# Go Get Better

This repository contains the solutions for the exercises contained in the
`go get better` training course. The solutions are found under the relevant
branch (for example `git checkout exercise-1`, or just select the branch from
within GitHub).

## Exercise 3

Turn your 'FizzBuzz' solution into a function in a package. Test the function. 

  * Your Function should take a number for the number of iterations
  * Your function should return a string containing the FizzBuzz sequence

  
> Feel free to reuse your code from exercise 1.
  
## Hints

  * You'll want to run `go mod init` (for example: 
    `go mod init training`)
  * `go test` will test the code
  * Your test file needs the suffix `_test.go`
  * `go test --covermode=count` shows test coverage
