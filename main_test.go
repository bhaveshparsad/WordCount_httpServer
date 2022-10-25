package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWordCount(t *testing.T) {
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong Status Code: got %v want %v", status, http.StatusOK)
	}

	expected := "Success"

	if rr.Body.String() != expected {
		t.Errorf("Handler returned Unexpected body : got %v want %v", rr.Body.String(), expected)
	}
}