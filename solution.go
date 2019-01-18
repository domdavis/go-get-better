package main

import "fmt"

const (
	fizz = "Fizz"
	buzz = "Buzz"
)

func main() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Printf("%-5d - %s%s", i, fizz, buzz)
		case i%3 == 0:
			fmt.Printf("%-5d - %s", i, fizz)
		case i%5 == 0:
			fmt.Printf("%-5d - %s", i, buzz)
		default:
			fmt.Print(i)
		}

		fmt.Println()
	}
}
