package main

import (
	"net/http"

	"github.com/domdavis/solution/handler"
)

func main() {
	http.HandleFunc("/", handler.FizzBuzz)

	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
