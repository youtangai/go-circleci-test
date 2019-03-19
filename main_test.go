package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSetUpRouter(t *testing.T) {
	router := setUpRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Fatal("code expect:200, actual:", w.Code)
	}
}
