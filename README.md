# Go Get Better

This repository contains the solutions for the exercises contained in the
`go get better` training course. The solutions are found under the relevant
branch (for example `git checkout exercise-1`, or just select the branch from
within GitHub).

## Exercise 11

Change Run back to a normal function, and change it to take a `Generator` 
interface.

The interface is: 

```go
package training

type Sequence []string

type Generator interface {
	Run(n int) (*Sequence, error)
}
```

## Hints

  * You can use `type name struct{}` to create a type that does not need to 
    store data
  * `func (_ name) Name() {}` can be used if the type function does not need
    access to the type
