package handlers

import (
	"fx-bank/internal/domain/models"
	repository "fx-bank/internal/domain/repositories"
	"github.com/shopspring/decimal"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TransferRequest struct {
	ProviderName      string          `json:"provider_name" binding:"required"`
	SenderAccountID   string          `json:"sender_account_id" binding:"required"`
	ReceiverAccountID string          `json:"receiver_account_id" binding:"required"`
	Amount            decimal.Decimal `json:"amount" binding:"required,gt=0"`
	Rate              decimal.Decimal `json:"rate" binding:"required"`
}

func (h *Handler) TransferToAccount(c *gin.Context) {
	var transferRequest TransferRequest
	bindingErr := c.BindJSON(&transferRequest)
	if bindingErr != nil {
		log.Println(bindingErr)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not process request. Try again later",
		})
		return
	}

	senderAccount, receiverAccount, notFoundErr := h.AccountRepository.FindAccountByID(
		transferRequest.SenderAccountID,
		transferRequest.ReceiverAccountID)

	if notFoundErr != nil {
		if notFoundErr == repository.ErrAccountSenderNotFound {
			log.Println(notFoundErr)
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Account specified for sender was not found",
			})
			return
		}

		if notFoundErr == repository.ErrAccountReceiverNotFound {
			log.Println(notFoundErr)
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Account specified for receiver was not found",
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError,
				gin.H{
					"message": "Could not process request. Try again later",
				})
			return
		}
	}

	//TODO: Add transaction support across multiple repositories
	transferErr := h.AccountRepository.Transfer(*senderAccount, *receiverAccount,
		transferRequest.Amount, transferRequest.Rate)
	if transferErr != nil {
		transaction := GetTransaction(&transferRequest, models.Failed,
			senderAccount, receiverAccount)
		createAccountTransactionErr := h.TransactionRepository.CreateTransaction(
			transaction,
		)
		if createAccountTransactionErr != nil {
			log.Println("failed to create transaction")
			log.Println(createAccountTransactionErr)
			c.JSON(
				http.StatusInternalServerError,
				gin.H{"message": "Could not process request. Try again later"},
			)
			return
		}

		if transferErr == models.ErrInsufficientBalance {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Could not withdraw amount due to insufficient balance",
			})
			return
		} else {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{"message": "Could not process request. Try again later"},
			)
			return
		}
	}

	transaction := GetTransaction(&transferRequest, models.Completed, senderAccount,
		receiverAccount)
	transferSuccessTransactionRecordErr := h.TransactionRepository.CreateTransaction(
		transaction,
	)

	if transferSuccessTransactionRecordErr != nil {
		log.Println(transferSuccessTransactionRecordErr)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not process request. Try again later",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Transfer successful",
		"data":    transferRequest,
	})
}

func GetTransaction(transferRequest *TransferRequest, status models.Status,
	senderAccount *models.Account, receiverAccount *models.Account,
) *models.Transaction {
	return &models.Transaction{
		Credit:              transferRequest.Amount,
		Debit:               transferRequest.Amount.Neg(),
		SenderAccountName:   senderAccount.Name,
		ReceiverAccountName: receiverAccount.Name,
		Rate:                transferRequest.Rate,
		ProviderName:        transferRequest.ProviderName,
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
		SenderCurrency:      senderAccount.Currency,
		ReceiverCurrency:    receiverAccount.Currency,
		Status:              status,
	}
}
