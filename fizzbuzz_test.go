package exercise5_test
import (
	"fmt"
	"testing"
	"training/exercise5"
)

func ExampleFizzBuzz() {
	if r, err := exercise5.FizzBuzz(15); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}

	// Output:
	// [1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz]
}

func TestFizzBuzz(t *testing.T) {
	_, err := exercise5.FizzBuzz(-1)
	if err != exercise5.ErrNegativeRange {
		t.Errorf("unexpected error: %s", err)
	}
}
