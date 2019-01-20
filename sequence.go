package exercise8

import (
	"errors"
	"strconv"
)

// Sequence holds the results of a sequence generator. The entry 0 holds the
// name of the sequence. Entry `n` holds the output for `n` on that sequence.
type Sequence []string

// Generator defines a function that can generate a Sequence.
type Generator func(int) (Sequence, error)

// ErrNegativeRange is returned if a sequence is given a range of less than 0.
var ErrNegativeRange = errors.New("cannot produce negative sequence")

// Simple generator that returns the sequence 1 to n. If n < 0 then this will
// return an error.
func Simple(n int) (Sequence, error) {
	if n < 0 {
		return Sequence{}, ErrNegativeRange
	}

	r := make(Sequence, n+1)
	r[0] = "Simple"

	for i := 1; i <= n; i++ {
		r[i] = strconv.Itoa(i)
	}

	return r, nil
}

// FizzBuzz returns the first n values of the FizzBuzz sequence. If n < 0 then
// this will return an error.
func FizzBuzz(n int) (Sequence, error) {
	if n < 0 {
		return Sequence{}, ErrNegativeRange
	}

	s := make(Sequence, n+1)
	s[0] = "FizzBuzz"

	for i := 1; i <= n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			s[i] = "FizzBuzz"
		case i%3 == 0:
			s[i] = "Fizz"
		case i%5 == 0:
			s[i] = "Buzz"
		default:
			s[i] = strconv.Itoa(i)
		}
	}
	return s, nil
}

// Run the sequence g from 1 to n.
func Run(g Generator, n int) (Sequence, error) {
	return g(n)
}
