package exercise6

import (
	"errors"
	"strconv"
)

var ErrNegativeRange = errors.New("cannot produce negative amounts of FizzBuzz")

// Generate the first n values of the FizzBuzz sequence. If n < 0 then this will
// return an error.
func FizzBuzz(n int) (map[int]string, error) {
	if n < 0 {
		return nil, ErrNegativeRange
	}

	m := make(map[int]string, n)

	for i := 1; i <= n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			m[i] = "FizzBuzz"
		case i%3 == 0:
			m[i] = "Fizz"
		case i%5 == 0:
			m[i] = "Buzz"
		default:
			m[i] = strconv.Itoa(i)
		}
	}

	return m, nil
}
