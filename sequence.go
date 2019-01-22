package training

import (
	"encoding/json"
	"errors"
	"strconv"
)

// Sequence holds the results of a sequence generator. The entry 0 holds the
// name of the sequence. Entry `n` holds the output for `n` on that sequence.
type Sequence []string

// Generator can generate a Sequence.
type Generator interface {
	Run(n int) (*Sequence, error)
}

type Simple struct{}
type FizzBuzz struct{}
type DeferredReverse struct{}

// ErrNegativeRange is returned if a sequence is given a range of less than 0.
var ErrNegativeRange = errors.New("cannot produce negative sequence")

// MarshalJSON renders a sequence as a JSON object with the sequence name as
// the key
func (s Sequence) MarshalJSON() ([]byte, error) {
	o := map[string][]string{}

	if len(s) > 1 {
		o[s[0]] = s[1:]
	}

	return json.Marshal(o)
}

// Simple generator that returns the sequence 1 to n. If n < 0 then this will
// return an error.
func (_ Simple) Run(n int) (*Sequence, error) {
	if n < 0 {
		return &Sequence{}, ErrNegativeRange
	}

	s := make(Sequence, n+1)
	s[0] = "Simple"

	for i := 1; i <= int(n); i++ {
		s[i] = strconv.Itoa(i)
	}

	return &s, nil
}

// FizzBuzz returns the first n values of the FizzBuzz sequence. If n < 0 then
// this will return an error.
func (_ FizzBuzz) Run(n int) (*Sequence, error) {
	if n < 0 {
		return &Sequence{}, ErrNegativeRange
	}

	s := make(Sequence, n+1)
	s[0] = "FizzBuzz"

	for i := 1; i <= int(n); i++ {
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
	return &s, nil
}

func (_ DeferredReverse) Run(n int) (*Sequence, error) {
	if n < 0 {
		return &Sequence{}, ErrNegativeRange
	}

	s := &Sequence{"DeferredReverse"}

	for i := 1; i <= int(n); i++ {
		defer func(i int) { *s = append(*s, strconv.Itoa(i)) }(i)
	}

	return s, nil
}

// Run the sequence g from 1 to n.
func Run(g Generator, n int) (Sequence, error) {
	p, err := g.Run(n)
	return *p, err
}
