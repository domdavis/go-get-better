package solution

import (
	"strconv"
	"strings"
)

// FizzBuzz generates the first n values of the FizzBuzz sequence. If n < 1
// then this will return the empty string.
func FizzBuzz(n int) string {
	var b strings.Builder
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

	return strings.TrimSpace(b.String())
}
