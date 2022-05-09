package training

import (
	"errors"
	"strconv"
	"strings"
)

var ErrNegativeRange = errors.New("cannot produce negative amounts of FizzBuzz")

// FizzBuzz generates the first n values of the FizzBuzz sequence. If n < 0 then
// this will return an error.
func FizzBuzz(n int) (string, error) {
	var b strings.Builder

	if n < 0 {
		return "", ErrNegativeRange
	}

	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			b.WriteString("Fizz")
		}

		if i%5 == 0 {
			b.WriteString("Buzz")
		}

		if i%3 != 0 && i%5 != 0 {
			b.WriteString(strconv.Itoa(i))
		}

		b.WriteString(" ")
	}

	return strings.TrimSpace(b.String()), nil
}
