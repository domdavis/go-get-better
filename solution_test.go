package solution_test

import (
	"fmt"
	"testing"

	"github.com/domdavis/solution"
)

func ExampleFizzBuzz() {
	fmt.Println(solution.FizzBuzz(15))

	// Output:
	// 1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz
}

func TestFizzBuzz(t *testing.T) {
	r := solution.FizzBuzz(-1)
	if r != "" {
		t.Errorf("unexpected FizzBuzz sequence: %s", r)
	}
}
