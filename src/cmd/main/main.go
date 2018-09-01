package main

import (
	"net/http"
	"time"
	"fmt"
)

func main() {
	handler := http.HandlerFunc(func (resp http.ResponseWriter, req *http.Request) {
		// nothing
	})

	s := &http.Server{
		Addr: ":42234",
		Handler: handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
	}

	err := s.ListenAndServe()
	if err != nil {
		fmt.Printf("error in start application: %v", err)
	}

	fmt.Printf("application terminated")
}
