package event

import (
	"testing"
	"net/http"
	"bytes"
	"net/http/httptest"

	"flash-logger/storage/memory"
	"flash-logger/config"
)

func TestHandler_GetRequest_MethodNotAllowed(t *testing.T) {
	handler := New(memory.New(make([]config.Project, 0, 0)))
	req, err := http.NewRequest("GET", "/1/events", bytes.NewBuffer([]byte(`{}`)))
	if err != nil {
		t.Fatalf("Request build error: %v.", err)
	}

	resp := httptest.NewRecorder()
	handler.ServeHTTP(resp, req)

	if resp.HeaderMap.Get("Content-Type") != "application/json; charset=utf-8" {
		t.Fatalf("unexpected content-type header")
	}

	if resp.Code != http.StatusMethodNotAllowed {
		t.Fatalf("unexpected status code")
	}
}
