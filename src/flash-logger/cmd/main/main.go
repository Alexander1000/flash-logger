package main

import (
	"net/http"
	"log"
	"flag"
	"fmt"
	"os"
	"syscall"
	"os/signal"
	"context"

	"flash-logger/api/v1/event"
	"flash-logger/storage/memory"
	"flash-logger/config"
	"flash-logger/handlers/auth"
	"flash-logger/api/v1/logs"
)

func main() {
	log.Println("Starting application ...")

	trap := make(chan os.Signal, 1)
	signal.Notify(trap, syscall.SIGINT, os.Interrupt, syscall.SIGTERM)

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

	storage := memory.New(cfg.Projects)

	http.Handle("/1/event", auth.NewAuthHandler(event.New(storage), cfg.Projects))

	http.Handle("/1/logs", auth.NewAuthHandler(logs.New(storage), cfg.Projects))

	go func() {
		if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), nil); err != nil {
			log.Fatalf("error in start application: %v", err)
		}
	}()

	ctx := context.Background()

	err = nil
	select {
	case <-trap:
		log.Println("termination signal caught")
	case <-ctx.Done():
		err = ctx.Err()
	}

	if err != nil {
		log.Printf("error in caught signal: %v", err)
	}

	log.Println("application terminated")
}
