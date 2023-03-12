package solution

import (
	"errors"
	"strconv"
)

// ErrNegativeRange is returned if a negative value is passed to FizzBuzz.
var ErrNegativeRange = errors.New("cannot produce negative amounts of FizzBuzz")

// FizzBuzz Generates the first n values of the FizzBuzz sequence. If n < 0 then
// this will return an error.
func FizzBuzz(n int) ([]string, error) {
	if n < 0 {
		return []string{}, ErrNegativeRange
	}

	s := make([]string, n)

	for i := 1; i <= n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			s[i-1] = "FizzBuzz"
		case i%3 == 0:
			s[i-1] = "Fizz"
		case i%5 == 0:
			s[i-1] = "Buzz"
		default:
			s[i-1] = strconv.Itoa(i)
		}
	}

	return s, nil
}

// Run the given function passing in n.
func Run(n int, f func(int) ([]string, error)) ([]string, error) {
	return f(n)
}
