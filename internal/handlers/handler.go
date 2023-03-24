package handlers

import repository "fx-bank/internal/domain/repositories"

type Handler struct {
	AccountRepository *repository.AccountRepository
}

func NewHandler(accountRepo *repository.AccountRepository) *Handler {
	return &Handler{
		AccountRepository: accountRepo,
	}
}
