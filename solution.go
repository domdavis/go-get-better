package main

import (
	"fmt"
	"strconv"
)

const (
	fizz = "Fizz"
	buzz = "Buzz"
)

func main() {
	FizzBuzz(100)
}

// FizzBuzz will output the first n iteration of FizzBuzz to STDOUT, with one
// element per line.
func FizzBuzz(n int) {
	for i := 1; i <= n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Printf("%s%s", fizz, buzz)
		case i%3 == 0:
			fmt.Print(fizz)
		case i%5 == 0:
			fmt.Print(buzz)
		default:
			fmt.Print(i)
		}

		fmt.Println()
	}
}

// FizzBuzz2 will output the first n iteration of FizzBuzz to STDOUT, with one
// element per line using a slightly different algorithm to FizzBuzz.
func FizzBuzz2(n int) {
	for i := 1; i <= n; i++ {
		var out string
		if i%3 == 0 {
			out = fizz
		}

		if i%5 == 0 {
			out += buzz
		}

		if i%3 != 0 && i%5 != 0 {
			out = strconv.Itoa(i)
		}

		fmt.Println(out)
	}
}
