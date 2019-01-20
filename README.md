# Go Get Better

This repository contains the solutions for the exercises contained in the
`go get better` training course. The solutions are found under the relevant
branch (for example `git checkout exercise-1`, or just select the branch from
within GitHub).

## Exercise 6

Turn your 'FizzBuzz' solution into one that returns a map. 

  * Your Function should take a number for the number of iterations
  * Your function should return an error if that number is less than 0
  * Your function should return a map containing an integer as a key, and the
    output for FizzBuzz for that number
    
> Feel free to reuse your code from exercise 5.
  
## Hints

  * The function signature will be `FizzBuzz(n int) (map[int]string, error)`
  * You can use `make(map[int]string, size)`
