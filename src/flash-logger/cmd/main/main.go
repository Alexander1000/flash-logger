package main

import (
	"net/http"
	"log"

	"flash-logger/api/v1/event"
	"flash-logger/api/v1/logs"
	"flash-logger/storage/memory"
)

func main() {
	log.Println("Starting application ...")

	// @todo загрузка ключей для валидации авторизаций (Bearer)

	storage := memory.New()

	http.Handle("/1/event", event.New(storage))

	http.Handle("/1/logs", logs.New(storage))

	if err := http.ListenAndServe(":42234", nil); err != nil {
		log.Fatalf("error in start application: %v", err)
	}

	log.Println("application terminated")
}
