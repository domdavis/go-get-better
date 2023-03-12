package solution_test

import (
	"fmt"
	"testing"

	"github.com/domdavis/solution"
)

func ExampleFizzBuzz() {
	n := 15

	r, err := solution.FizzBuzz(n)

	if err != nil {
		fmt.Println(err)
	}

	for i := 1; i <= n; i++ {
		fmt.Print(r[i], " ")
	}

	// Output:
	// 1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz
}

func TestFizzBuzz(t *testing.T) {
	_, err := solution.FizzBuzz(-1)
	if err != solution.ErrNegativeRange {
		t.Errorf("unexpected error: %s", err)
	}
}
