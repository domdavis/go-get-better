package solution_test

import (
	"fmt"
	"testing"

	"github.com/domdavis/solution"
)

func ExampleRun() {
	r, err := solution.Run(15, solution.FizzBuzz)

	if err != nil {
		fmt.Println(err)
	}

	for _, v := range r {
		fmt.Print(v, " ")
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
