package handlers

import (
	"fx-bank/internal/domain/models"
	repository "fx-bank/internal/domain/repositories"
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
			"message": "could not parse request. check API documentation",
		})
		return
	}

	userIdAsUUID, uuidParseErr := uuid.Parse(createAccountRequest.UserID)
	if uuidParseErr != nil {
		log.Println(uuidParseErr)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid user id, check API documentation",
		})
		return
	}

	err := h.AccountRepository.FindAccountByNameAndUserID(createAccountRequest.Name, createAccountRequest.UserID)

	if err != nil && err == repository.ErrAccountNameAlreadyUsed {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "account name has been used already",
		})
		return
	}

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "couldn't perform operation, check back later",
		})
		return
	}

	account.Balance = 0
	account.Name = createAccountRequest.Name
	account.Currency = createAccountRequest.Currency
	account.CreatedAt = time.Now()
	account.LastModified = time.Now()
	account.UserID = userIdAsUUID

	err = h.AccountRepository.CreateAccount(&account)
	if err == repository.ErrAccountCannotCreated {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "could not create account, try again later",
		})
		return
	}
	//TODO: Map account to new "type" and eliminate sensitive fields
	c.JSON(http.StatusCreated, gin.H{
		"message": "account successfully created",
		"data":    account,
	})
}
