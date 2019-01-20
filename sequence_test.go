package training_test

import (
	"fmt"
	"testing"
	"training"
)

func ExampleSequence() {
	simple := training.Generator(training.Simple)
	if r, err := simple.Run(5); err != nil {
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
		generator training.Generator

		sequence training.Sequence
	}{
		{
			name:      "simple",
			generator: training.Simple,
			sequence:  training.Sequence{"Simple", "1", "2", "3", "4", "5"},
		},
		{
			name:      "FizzBuzz",
			generator: training.FizzBuzz,
			sequence:  training.Sequence{"FizzBuzz", "1", "2", "Fizz", "4", "Buzz"},
		},
		{
			name:      "DeferredReverse",
			generator: training.DeferredReverse,
			sequence:  training.Sequence{"DeferredReverse", "5", "4", "3", "2", "1"},
		},
	} {
		t.Run(fmt.Sprintf("%s 1-5", test.name), func(t *testing.T) {
			g := training.Generator(test.generator)
			r, err := g.Run(5)

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
			g := training.Generator(test.generator)
			r, err := g.Run(0)

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
			g := training.Generator(test.generator)
			_, err := g.Run(-1)

			if err != training.ErrNegativeRange {
				t.Errorf("unexptected error for %s(-1): %s", test.name, err)
			}
		})
	}
}
