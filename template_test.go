package template_test

import (
	"fmt"
	"testing"

	"github.com/domdavis/template"
)

func ExampleSequence() {
	if r, err := template.Run(template.Simple, 5); err != nil {
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
		generator template.Generator

		sequence template.Sequence
	}{
		{
			name:      "simple",
			generator: template.Simple,
			sequence:  template.Sequence{"Simple", "1", "2", "3", "4", "5"},
		},
		{
			name:      "FizzBuzz",
			generator: template.FizzBuzz,
			sequence:  template.Sequence{"FizzBuzz", "1", "2", "Fizz", "4", "Buzz"},
		},
	} {
		t.Run(fmt.Sprintf("%s 1-5", test.name), func(t *testing.T) {
			r, err := template.Run(test.generator, 5)

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
			r, err := template.Run(test.generator, 0)

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
			_, err := template.Run(test.generator, -1)

			if err != template.ErrNegativeRange {
				t.Errorf("unexptected error for %s(-1): %s", test.name, err)
			}
		})
	}
}
