package handlers

import (
	"fx-bank/config"
	"fx-bank/internal/domain/repositories"
	"fx-bank/internal/postgres"
	"fx-bank/server"
	"log"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

var routeHandlers *gin.Engine

func TestMain(m *testing.M) {
	err := config.LoadTestConfig("../../.env.test")
	if err != nil {
		panic(err)
	}

	db, err := postgres.Init()
	if err != nil {
		log.Fatal(err)
	}

	userRepo := repository.NewUserRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)
	accountRepo := repository.NewAccountRepository(db)

	handler := NewHandler(accountRepo, transactionRepo, userRepo)
	srv := server.New(handler)
	routeHandlers = srv.SetupRoutes()
	os.Exit(m.Run())
}
