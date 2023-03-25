package handlers

import repository "fx-bank/internal/domain/repositories"

type Handler struct {
	AccountRepository     *repository.AccountRepository
	TransactionRepository *repository.TransactionRepository
	UserRepository        *repository.UserRepository
}

func NewHandler(accountRepo *repository.AccountRepository,
	transactionRepo *repository.TransactionRepository, userRepo *repository.UserRepository) *Handler {
	return &Handler{
		AccountRepository:     accountRepo,
		TransactionRepository: transactionRepo,
		UserRepository:        userRepo,
	}
}
