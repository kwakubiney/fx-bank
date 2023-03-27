package handlers

import (
	"fx-bank/internal/domain/models"
	repository "fx-bank/internal/domain/repositories"
	"fx-bank/internal/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type SignUpRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) SignUp(c *gin.Context) {
	var signUpRequest SignUpRequest
	var user models.User
	bindingErr := c.BindJSON(&signUpRequest)
	if bindingErr != nil {
		log.Println(bindingErr)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse request. check API documentation",
		})
		return
	}

	hashedPassword := utils.GetHash([]byte(signUpRequest.Password))
	user.Username = signUpRequest.Username
	user.Password = hashedPassword
	user.LastModified = time.Now()
	user.CreatedAt = time.Now()

	_, err, rowsAffected := h.UserRepository.FindUserAccountByUsername(user.Username)
	if err != nil && err != repository.ErrUserAccountDoesNotExist {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not process request. Try again later.",
		})
		return
	}

	if rowsAffected > 0 {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User with username selected already exists.",
		})
		return
	}

	createAccountErr := h.UserRepository.CreateUserAccount(&user)
	if createAccountErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not process request. Try again later.",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User successfully created",
	})
	return

}
