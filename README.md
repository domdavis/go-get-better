# Go Get Better

This repository contains the solutions for the exercises contained in the
`go get better` training course. The solutions are found under the relevant
branch (for example `git checkout exercise-1`, or just select the branch from
within GitHub).

## Exercise 11

Update the sequence type so when it's rendered as JSON it's output in the format 
`{"name": ["1", "2",...]}`. An empty sequence should simply be `{}`

## Hints

  * You'll need to implement `func (s Sequence) MarshalJSON() ([]byte, error)`
  * The output is basically a map
