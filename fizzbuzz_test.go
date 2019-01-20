package exercise6_test
import (
	"fmt"
	"testing"
	"training/exercise6"
)

func ExampleFizzBuzz() {
	n := 15

	r, err := exercise6.FizzBuzz(n)

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
	_, err := exercise6.FizzBuzz(-1)
	if err != exercise6.ErrNegativeRange {
		t.Errorf("unexpected error: %s", err)
	}
}
