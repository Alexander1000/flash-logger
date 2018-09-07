package main

import (
	"net/http"
	"log"
	"flag"

	"flash-logger/api/v1/event"
	"flash-logger/storage/memory"
	"flash-logger/config"
	"flash-logger/handlers/auth"
	"flash-logger/api/v1/logs"
)

func main() {
	log.Println("Starting application ...")

	configPath := flag.String("c", "", "config file")

	flag.Parse()

	if len(*configPath) == 0 {
		log.Fatalf("unknown config file")
	}

	var err error
	var cfg *config.Config
	if cfg, err = config.LoadFromFile(*configPath); err != nil {
		log.Fatalf("error in load config from file: %v", err)
	}

	if err := cfg.LoadProjects(); err != nil {
		log.Fatalf("error in load projects: %v", err)
	}

	log.Printf("Config file: %s", *configPath)

	for _, project := range cfg.Projects {
		log.Printf("Registered project '%s'", project.Title)
	}

	log.Printf("Starting service on port: %d", cfg.Port)

	// @todo загрузка ключей для валидации авторизаций (Bearer)

	storage := memory.New()

	http.Handle("/1/event", auth.NewAuthHandler(event.New(storage), cfg.Projects))

	http.Handle("/1/logs", auth.NewAuthHandler(logs.New(storage), cfg.Projects))

	// @todo параметризовать порт запуска
	if err := http.ListenAndServe(":42234", nil); err != nil {
		log.Fatalf("error in start application: %v", err)
	}

	log.Println("application terminated")
}
