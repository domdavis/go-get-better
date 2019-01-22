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
