package main

import "fmt"

const (
	fizz = "Fizz"
	buzz = "Buzz"
)

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("%-5d - %s%s", i, fizz, buzz)
		} else if i%3 == 0 {
			fmt.Printf("%-5d - %s", i, fizz)
		} else if i%5 == 0 {
			fmt.Printf("%-5d - %s", i, buzz)
		} else {
			fmt.Print(i)
		}
		fmt.Println()
	}
}
