package handlers

import repository "fx-bank/internal/domain/repositories"

type Handler struct {
	AccountRepository     *repository.AccountRepository
	TransactionRepository *repository.TransactionRepository
}

func NewHandler(accountRepo *repository.AccountRepository,
	transactionRepo *repository.TransactionRepository) *Handler {
	return &Handler{
		AccountRepository:     accountRepo,
		TransactionRepository: transactionRepo,
	}
}
