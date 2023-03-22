package main

import (
	"fx-bank/config"
	"fx-bank/internal/postgres"
	"fx-bank/server"
	"log"
)

func main() {
	err := config.LoadNormalConfig()
	if err != nil {
		log.Fatal(err)
	}

	_, err = postgres.Init()
	if err != nil {
		log.Fatal(err)
	}

	httpServer := server.New()
	httpServer.Start()
}
