package training_test

import (
	"fmt"
	"testing"
	"training"
)

func ExampleFizzBuzz() {
	fmt.Println(training.FizzBuzz(15))

	// Output:
	// 1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz
}

func TestFizzBuzz(t *testing.T) {
	r := training.FizzBuzz(-1)
	if r != "" {
		t.Errorf("unexpected FizzBuzz sequence: %s", r)
	}
}
