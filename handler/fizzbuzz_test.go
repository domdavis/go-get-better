package handler_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/domdavis/solution/handler"
)

func TestFizzBuzz(t *testing.T) {
	for input, output := range map[int]string{
		1:  "1",
		3:  "Fizz",
		5:  "Buzz",
		15: "FizzBuzz",
	} {
		t.Run(fmt.Sprintf("Check %d", input), func(t *testing.T) {
			t.Parallel()

			url := fmt.Sprintf("/fizzbuzz/%d", input)
			req, err := http.NewRequest("GET", url, nil)

			if err != nil {
				t.Fatal(err)
			}

			recorder := httptest.NewRecorder()

			http.HandlerFunc(handler.FizzBuzz).ServeHTTP(recorder, req)

			responses := strings.Split(
				strings.Trim(recorder.Body.String(), "[]\""), "\",\"")

			switch {
			case recorder.Code != http.StatusOK:
				t.Errorf("handler returned wrong status code: got %v want %v",
					recorder.Code, http.StatusOK)
			case len(responses) != input:
				t.Errorf("Expected %d responses, got %d", input, len(responses))
			case responses[0] != "1":
				t.Errorf("handler returned wrong output: got %v want 1",
					responses[0])
			case responses[input-1] != output:
				t.Errorf("handler returned wrong output: got %v want %v",
					responses[input-1], output)
			}
		})
	}

	t.Run("An invalid request will return 404", func(t *testing.T) {
		t.Parallel()

		req, err := http.NewRequest("GET", "/fizzbuzz/fail", nil)

		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()

		http.HandlerFunc(handler.FizzBuzz).ServeHTTP(recorder, req)

		if recorder.Code != http.StatusNotFound {
			t.Errorf("Expected %d, got %d", http.StatusNotFound, recorder.Code)
		}

	})
}
