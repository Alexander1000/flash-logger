package main

import (
	"net/http"
	"log"

	"flash-logger/api/v1/event"
)

func main() {
	log.Println("Starting application ...")

	http.Handle("/event", event.New())

	if err := http.ListenAndServe(":42234", nil); err != nil {
		log.Fatalf("error in start application: %v", err)
	}

	log.Printf("application terminated")
}
