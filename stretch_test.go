package main

import (
	"fmt"
	"strconv"
)

// Possible alternative solution to FizzBuzz. We use if rather than case because
// it's not an if/else tree and some cases will hit two branches.
//
// We print on one line and only go up to 15 to reduce the length of the output.
func ExampleFizzBuzz() {
	for i := 1; i <= 19; i++ {
		var out string

		if i%3 == 0 {
			out = "Fizz"
		}

		if i%5 == 0 {
			out += "Buzz"
		}

		if i%3 != 0 && i%5 != 0 {
			out = strconv.Itoa(i)
		}

		fmt.Printf("%s ", out)
	}

	// Output:
	// 1 2 Fizz 4 Buzz Fizz 7 8 Fizz Buzz 11 Fizz 13 14 FizzBuzz 16 17 Fizz 19
}
