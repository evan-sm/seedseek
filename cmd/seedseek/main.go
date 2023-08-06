package main

import (
	"log"

	"seedseek/internal/app"
	"seedseek/internal/config"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
