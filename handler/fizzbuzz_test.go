package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"training/handler"
)

func TestFizzBuzz(t *testing.T) {
	req, err := http.NewRequest("GET", "/fizzbuzz/3", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	h := http.HandlerFunc(handler.FizzBuzz)


	h.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
