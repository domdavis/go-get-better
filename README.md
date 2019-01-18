# Go Get Better

This repository contains the solutions for the exercises contained in the
`go get better` training course. The solutions are found under the relevant
branch (for example `git checkout exercise-1`, or just select the branch from
within GitHub).

## Exercise 2

Create a binary that will perform 'FizzBuzz'. The rules for this are:

  * For the numbers 1 to 100
  * If a number is divisible by 3 output Fizz
  * If a number is divisible by 5 output Buzz
  * If a number is divisible by 3 and 5 output FizzBuzz
  * Otherwise output the number
  
>Feel free to reuse your code from exercise 1. The goal is to get a working 
environment.
  
## Hints

  * You'll want to run `go mod init` (for example: 
    `go mod init training/exercise2`)
  * `go build` compile the code
  * `go build -o fizzbuzz` will compile to a binary called `fizzbuzz`
  * you can quickly run the code to test it using `go run solution.go` (replace
    `solution.go` with whatever you called your source file)
