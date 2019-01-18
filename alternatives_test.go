package main

import (
	"fmt"
	"math/rand"
)

// Idiomatically with Go if there are multiple if/elseif branches we use a
// switch statement instead.
//
// We print on one line and only go up to 19 to reduce the length of the output.
func ExampleIdiomaticSolution() {
	for i := 1; i <= 19; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Print("fizzbuzz ")
		case i%3 == 0:
			fmt.Print("fizz ")
		case i%5 == 0:
			fmt.Print("buzz ")
		default:
			fmt.Printf("%d ", i)
		}
	}

	// Output:
	// 1 2 fizz 4 buzz fizz 7 8 fizz buzz 11 fizz 13 14 fizzbuzz 16 17 fizz 19
}

// Utilise a feature of the Go random number generator to produce the sequence.
// This is likely to break if the random number generator changes.
//
// We print on one line and only go up to 19 to reduce the length of the output.
func ExampleBrittleSolution() {
	for i := 1; i <= 19; i++ {
		if i%15 == 1 {
			rand.Seed(int64(176064004))
		}
		fmt.Print([]string{fmt.Sprintf("%d ", i),
			"fizz ", "buzz ", "fizzbuzz "}[rand.Int63()%4])
	}

	// Output:
	// 1 2 fizz 4 buzz fizz 7 8 fizz buzz 11 fizz 13 14 fizzbuzz 16 17 fizz 19
}

// Utilise the fact that FizzBuzz is a repeating pattern which we can use to
// select the correct result from a slice of possible answers. This is a less
// brittle version of the above random number solution, but still relies on
// "magic" which may not be immediately obvious.
//
// We print on one line and only go up to 19 to reduce the length of the output.
func ExampleCompactSolution() {
	pattern := []int{0, 0, 1, 0, 2, 1, 0, 0, 1, 2, 0, 1, 0, 0, 3}

	for i := 1; i <= 19; i++ {
		fmt.Print([]string{fmt.Sprintf("%d ", i),
			"fizz ", "buzz ", "fizzbuzz "}[pattern[(i-1)%15]])
	}

	// Output:
	// 1 2 fizz 4 buzz fizz 7 8 fizz buzz 11 fizz 13 14 fizzbuzz 16 17 fizz 19
}
