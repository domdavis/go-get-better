package exercise7_test
import (
	"fmt"
	"testing"
	"training/exercise7"
)

func ExampleSequence() {
	r, err := exercise7.Run(15, exercise7.FizzBuzz)

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
	_, err := exercise7.FizzBuzz(-1)
	if err != exercise7.ErrNegativeRange {
		t.Errorf("unexpected error: %s", err)
	}
}
