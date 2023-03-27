package handlers

import (
	"fx-bank/internal/domain/models"
	repository "fx-bank/internal/domain/repositories"
	"github.com/shopspring/decimal"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateAccountRequest struct {
	Name     string          `json:"name" binding:"required"`
	Currency models.Currency `json:"currency" binding:"required"`
	UserID   string          `json:"user_id" binding:"required"`
}

func (h *Handler) CreateAccount(c *gin.Context) {
	var account models.Account
	var createAccountRequest CreateAccountRequest
	bindingErr := c.BindJSON(&createAccountRequest)
	if bindingErr != nil {
		log.Println(bindingErr)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not process request. Try again later.",
		})
		return
	}

	userIdAsUUID, uuidParseErr := uuid.Parse(createAccountRequest.UserID)
	if uuidParseErr != nil {
		log.Println(uuidParseErr)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not process request. Try again later.",
		})
		return
	}

	err := h.AccountRepository.FindAccountByNameAndUserID(createAccountRequest.Name, createAccountRequest.UserID)

	if err != nil && err == repository.ErrAccountNameAlreadyUsed {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "An account with name selected already exists",
		})
		return
	}

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not process request. Try again later",
		})
		return
	}

	account.Balance = decimal.NewFromInt(100)
	log.Println("Account balance", account.Balance)
	account.Name = createAccountRequest.Name
	account.Currency = createAccountRequest.Currency
	account.CreatedAt = time.Now()
	account.LastModified = time.Now()
	account.UserID = userIdAsUUID

	log.Println(account)

	err = h.AccountRepository.CreateAccount(&account)
	if err == repository.ErrAccountCannotCreated {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not process request. Try again later.",
		})
		return
	}
	//TODO: Map account to new "type" and eliminate sensitive fields
	c.JSON(http.StatusCreated, gin.H{
		"message": "Account successfully created",
		"data":    account,
	})
}
