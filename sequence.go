package training

import (
	"errors"
	"strconv"
)

var ErrNegativeRange = errors.New("cannot produce negative amounts of FizzBuzz")

// Generate the first n values of the FizzBuzz sequence. If n < 0 then this will
// return an error.
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


func Run(n int, f func(int) ([]string, error)) ([]string, error) {
	return f(n)
}
