package training_test

import (
	"fmt"
	"testing"
	"training"
)

func ExampleRun() {
	r, err := training.Run(15, training.FizzBuzz)

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
	_, err := training.FizzBuzz(-1)
	if err != training.ErrNegativeRange {
		t.Errorf("unexpected error: %s", err)
	}
}
