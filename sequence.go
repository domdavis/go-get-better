package training

import (
	"errors"
	"strconv"
	"strings"
)

// Sequence holds the results of a sequence generator. The entry 0 holds the
// name of the sequence. Entry `n` holds the output for `n` on that sequence.
type Sequence struct {
	Name     string
	Elements []string
}

type Step func(in <-chan int, out chan<- string)

// Generator can generate a Sequence.
type Generator struct {
	Name  string
	Steps []Step
}

// ErrNegativeRange is returned if a sequence is given a range of less than 0.
var ErrNegativeRange = errors.New("cannot produce negative sequence")

// Simple just returns what was passed into it
var Simple = Generator{
	Name:  "Simple",
	Steps: []Step{passthrough},
}

// FizzBuzz returns the first n values of the FizzBuzz sequence.
var FizzBuzz = Generator{
	Name:  "FizzBuzz",
	Steps: []Step{fizz, buzz, number},
}

// Run the sequence g from 1 to n.
func (g Generator) Run(n int) (Sequence, error) {
	s := Sequence{Name: g.Name}

	if n < 0 {
		return s, ErrNegativeRange
	}

	in := make(chan int)
	out := make(chan string)

	s.Elements = make([]string, n)

	for i := 1; i <= n; i++ {
		b := strings.Builder{}
		for _, step := range g.Steps {
			go step(in, out)
			in <- i
			b.WriteString(<-out)
		}
		s.Elements[i-1] = b.String()
	}

	return s, nil
}

func passthrough(in <-chan int, out chan<- string) {
	out <- strconv.Itoa(<-in)
}

func fizz(in <-chan int, out chan<- string) {
	if <-in%3 == 0 {
		out <- "Fizz"
	} else {
		out <- ""
	}
}

func buzz(in <-chan int, out chan<- string) {
	if <-in%5 == 0 {
		out <- "Buzz"
	} else {
		out <- ""
	}
}

func number(in <-chan int, out chan<- string) {
	i := <-in
	if i%3 != 0 && i%5 != 0 {
		out <- strconv.Itoa(i)
	} else {
		out <- ""
	}
}
