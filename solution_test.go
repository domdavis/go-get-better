package solution_test

import (
	"fmt"
	"testing"

	"github.com/domdavis/solution"
)

func ExampleSequence() {
	g := solution.Simple
	if r, err := g.Run(5); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(r)
	}

	// Output:
	// {Simple [1 2 3 4 5]}
}

func TestSequences(t *testing.T) {
	for _, test := range []struct {
		generator solution.Generator
		sequence  solution.Sequence
	}{
		{
			generator: solution.Simple,
			sequence: solution.Sequence{
				Name:     "Simple",
				Elements: []string{"1", "2", "3", "4", "5"},
			},
		},
		{
			generator: solution.FizzBuzz,
			sequence: solution.Sequence{
				Name:     "FizzBuzz",
				Elements: []string{"1", "2", "Fizz", "4", "Buzz"},
			},
		},
	} {
		t.Run(fmt.Sprintf("%s 1-5", test.sequence.Name), func(t *testing.T) {
			r, err := test.generator.Run(5)

			if err != nil {
				t.Errorf("unexpected error for sequence %s: %s",
					test.sequence.Name, err)
			}

			if r.Name != test.sequence.Name {
				t.Errorf("unexpected sequence name '%s', expected '%s'",
					r.Name, test.sequence.Name)
			}

			if len(r.Elements) != len(test.sequence.Elements) {
				t.Errorf("unexpected sequence for %s: %v",
					test.sequence.Name, r)
			}

			for i, v := range r.Elements {
				if v != test.sequence.Elements[i] {
					t.Errorf("Expected %s, got %s for %s",
						test.sequence.Elements[i], v, test.sequence.Name)
				}
			}
		})

		t.Run(fmt.Sprintf("%s 0", test.sequence.Name), func(t *testing.T) {
			r, err := test.generator.Run(0)

			if err != nil {
				t.Errorf("unexpected error for sequence %s: %s",
					test.sequence.Name, err)
			}

			if r.Name != test.sequence.Name {
				t.Errorf("unexpected sequence name '%s', expected '%s'",
					r.Name, test.sequence.Name)
			}

			if len(r.Elements) != 0 {
				t.Errorf("unexpected sequence for %s: %v",
					test.sequence.Name, r)
			}
		})

		t.Run(fmt.Sprintf("%s -1", test.sequence.Name), func(t *testing.T) {
			r, err := test.generator.Run(-1)

			if r.Name != test.sequence.Name {
				t.Errorf("unexpected sequence name '%s', expected '%s'",
					r.Name, test.sequence.Name)
			}

			if r.Name != test.sequence.Name {
				t.Errorf("unexpected sequence name '%s', expected '%s'",
					r.Name, test.sequence.Name)
			}

			if err != solution.ErrNegativeRange {
				t.Errorf("unexptected error for %s(-1): %s",
					test.sequence.Name, err)
			}
		})
	}
}
