package exercise4_test
import (
	"fmt"
	"testing"
	"training/exercise4"
)

func ExampleFizzBuzz() {
	if r, err := exercise4.FizzBuzz(15); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}

	// Output:
	// 1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz
}

func TestFizzBuzz(t *testing.T) {
	_, err := exercise4.FizzBuzz(-1)
	if err != exercise4.ErrNegativeRange {
		t.Errorf("unexpected error: %s", err)
	}
}
