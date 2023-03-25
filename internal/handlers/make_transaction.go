package handlers

import (
	"fx-bank/internal/domain/models"
	repository "fx-bank/internal/domain/repositories"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TransferRequest struct {
	ProviderName      string `json:"provider_name" binding:"required"`
	SenderAccountID   string `json:"sender_account_id" binding:"required"`
	ReceiverAccountID string `json:"receiver_account_id" binding:"required"`
	Amount            int64  `json:"amount" binding:"required,gt=0"`
	Rate              int64  `json:"rate" binding:"required"`
}

func (h *Handler) TransferToAccount(c *gin.Context) {
	var transferRequest TransferRequest
	bindingErr := c.BindJSON(&transferRequest)
	if bindingErr != nil {
		log.Println(bindingErr)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse request, binding error",
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
				"message": "account sender specified was not found",
			})
			return
		}

		if notFoundErr == repository.ErrAccountReceiverNotFound {
			log.Println(notFoundErr)
			c.JSON(http.StatusNotFound, gin.H{
				"message": "account receiver specified was not found",
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError,
				gin.H{
					"message": "could not get user from database",
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
				gin.H{"message": "failed to create transaction"},
			)
			return
		}

		if transferErr == models.ErrInsufficientBalance {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "could not withdraw amount due to insufficient balance",
			})
			return
		} else {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{"message": "settlement failed on both accounts"},
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
			"message": "could not create transaction",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "transfer successful",
		"data":    transferRequest,
	})
}

func GetTransaction(transferRequest *TransferRequest, status models.Status,
	senderAccount *models.Account, receiverAccount *models.Account,
) *models.Transaction {
	return &models.Transaction{
		Credit:              transferRequest.Amount,
		Debit:               -(transferRequest.Amount),
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
