package main

import (
	"net/http"
	"log"
	"flag"

	"flash-logger/api/v1/event"
	"flash-logger/api/v1/logs"
	"flash-logger/storage/memory"
)

func main() {
	log.Println("Starting application ...")

	configPath := flag.String("c", "", "config file")

	flag.Parse()

	if len(*configPath) == 0 {
		log.Fatalf("unknown config file")
	}

	log.Printf("Config file: %s", *configPath)

	// @todo загрузка ключей для валидации авторизаций (Bearer)

	storage := memory.New()

	http.Handle("/1/event", event.New(storage))

	http.Handle("/1/logs", logs.New(storage))

	// @todo параметризовать порт запуска
	if err := http.ListenAndServe(":42234", nil); err != nil {
		log.Fatalf("error in start application: %v", err)
	}

	log.Println("application terminated")
}
