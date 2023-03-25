package main

import (
	"fx-bank/config"
	repository "fx-bank/internal/domain/repositories"
	"fx-bank/internal/handlers"
	"fx-bank/internal/postgres"
	"fx-bank/server"
	"log"
)

func main() {
	err := config.LoadNormalConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgres.Init()
	if err != nil {
		log.Fatal(err)
	}

	accountRepo := repository.NewAccountRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)
	userRepo := repository.NewUserRepository(db)
	handler := handlers.NewHandler(accountRepo, transactionRepo, userRepo)

	httpServer := server.New(handler)
	httpServer.Start()
}
