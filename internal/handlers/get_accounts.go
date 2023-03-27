package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAccounts(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not process request. Try again later",
		})
		return
	}

	accounts, err := h.AccountRepository.FindAllAccountsByUserID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not process request. Try again later",
		})
		return
	}

	//TODO: Map account to new "type" and eliminate sensitive fields
	c.JSON(http.StatusOK, gin.H{
		"message": "Account successfully retrieved",
		"data":    accounts,
	})
}
