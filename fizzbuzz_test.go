package exercise3_test

import (
	"fmt"
	"testing"
	"training/exercise3"
)

func ExampleFizzBuzz() {
	fmt.Println(exercise3.FizzBuzz(15))

	// Output:
	// 1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz
}

func TestFizzBuzz(t *testing.T) {
	r := exercise3.FizzBuzz(-1)
	if r != "" {
		t.Errorf("unexpected FizzBuzz sequence: %s", r)
	}
}
