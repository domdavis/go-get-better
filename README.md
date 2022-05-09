# Go Get Better

This repository contains the solutions for the exercises contained in the
`go get better` training course. The solutions are found under the relevant
branch (for example `git checkout exercise-1`, or just select the branch from
within GitHub).

## Exercise 9

Use the code from exercise 9 implement a 'deferred reverse` sequence
      
## Hints

  * `&` means pointer to, and `*` means get me the value pointed to
  * You'll want to change `Generator` to `func(n int) (*Sequence, error)`, and
    update `Simple` and `FizzBuzz`    
  * your defer func will want to be `defer func(i int) { ... }(i)`

  * You'll want to update `Run` to dereference the pointer to the sequence
