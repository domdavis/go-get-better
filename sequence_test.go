package exercise9_test

import (
	"fmt"
	"testing"
	"training/exercise9"
)

func ExampleSequence() {
	if r, err := exercise9.Run(exercise9.Simple, 5); err != nil {
		fmt.Println(err)
	} else {
		for _, v := range r {
			fmt.Print(v, " ")
		}
	}

	// Output:
	// Simple 1 2 3 4 5
}

func TestSequences(t *testing.T) {
	for _, test := range []struct {
		name      string
		generator exercise9.Generator

		sequence exercise9.Sequence
	}{
		{
			name:      "simple",
			generator: exercise9.Simple,
			sequence:  exercise9.Sequence{"Simple", "1", "2", "3", "4", "5"},
		},
		{
			name:      "FizzBuzz",
			generator: exercise9.FizzBuzz,
			sequence:  exercise9.Sequence{"FizzBuzz", "1", "2", "Fizz", "4", "Buzz"},
		},
		{
			name:      "DeferredReverse",
			generator: exercise9.DeferredReverse,
			sequence:  exercise9.Sequence{"DeferredReverse", "5", "4", "3", "2", "1"},
		},
	} {
		t.Run(fmt.Sprintf("%s 1-5", test.name), func(t *testing.T) {
			r, err := exercise9.Run(test.generator, 5)

			if err != nil {
				t.Errorf("unexpected error for sequence %s: %s", test.name, err)
			}

			if len(r) != len(test.sequence) {
				t.Errorf("unexpected sequence for %s: %v", test.name, r)
			}

			for i, v := range r {
				if v != test.sequence[i] {
					t.Errorf("Expected %s, got %s for %s",
						test.sequence[i], v, test.name)
				}
			}
		})

		t.Run(fmt.Sprintf("%s 0", test.name), func(t *testing.T) {
			r, err := exercise9.Run(test.generator, 0)

			if err != nil {
				t.Errorf("unexpected error for sequence %s: %s", test.name, err)
			}

			if len(r) != 1 {
				t.Errorf("unexpected sequence for %s: %v", test.name, r)
			} else if r[0] != test.sequence[0] {
				t.Errorf("unexpected sequence name for %s: %v", test.name, r)
			}
		})

		t.Run(fmt.Sprintf("%s -1", test.name), func(t *testing.T) {
			_, err := exercise9.Run(test.generator, -1)

			if err != exercise9.ErrNegativeRange {
				t.Errorf("unexptected error for %s(-1): %s", test.name, err)
			}
		})
	}
}
