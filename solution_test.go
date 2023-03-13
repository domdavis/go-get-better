package solution_test

import (
	"fmt"
	"testing"

	"github.com/domdavis/solution"
)

func ExampleSequence() {
	if r, err := solution.Run(solution.Simple{}, 5); err != nil {
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
		generator solution.Generator

		sequence solution.Sequence
	}{
		{
			name:      "simple",
			generator: solution.Simple{},
			sequence:  solution.Sequence{"Simple", "1", "2", "3", "4", "5"},
		},
		{
			name:      "FizzBuzz",
			generator: solution.FizzBuzz{},
			sequence:  solution.Sequence{"FizzBuzz", "1", "2", "Fizz", "4", "Buzz"},
		},
		{
			name:      "DeferredReverse",
			generator: solution.DeferredReverse{},
			sequence:  solution.Sequence{"DeferredReverse", "5", "4", "3", "2", "1"},
		},
	} {
		t.Run(fmt.Sprintf("%s 1-5", test.name), func(t *testing.T) {
			r, err := solution.Run(test.generator, 5)

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
			r, err := solution.Run(test.generator, 0)

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
			_, err := solution.Run(test.generator, -1)

			if err != solution.ErrNegativeRange {
				t.Errorf("unexptected error for %s(-1): %s", test.name, err)
			}
		})
	}
}
