package handler

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"
)

var re = regexp.MustCompile(`/fizzbuzz/([0-9]+)$`)

// FizzBuzz handler will respond to an appropriately formed URL with the
// FizzBuzz element for the given number.
func FizzBuzz(w http.ResponseWriter, r *http.Request) {
	match := re.FindStringSubmatch(r.URL.EscapedPath())

	if len(match) != 2 {
		w.WriteHeader(404)
		return
	}

	n, _ := strconv.Atoi(match[1])
	b, _ := json.Marshal(fizzbuzz(n))

	if _, err := w.Write(b); err != nil {
		w.WriteHeader(500)
	}
}

func fizzbuzz(n int) []string {
	s := make([]string, n)

	for i := 1; i <= n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			s[i-1] = "FizzBuzz"
		case i%3 == 0:
			s[i-1] = "Fizz"
		case i%5 == 0:
			s[i-1] = "Buzz"
		default:
			s[i-1] = strconv.Itoa(i)
		}
	}
	return s
}
