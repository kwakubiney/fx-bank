package handlers

import (
	"fx-bank/internal/domain/models"
	repository "fx-bank/internal/domain/repositories"
	"fx-bank/internal/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token  string `json:"token"`
	UserId string `json:"user_id"`
}

func (h *Handler) Login(c *gin.Context) {
	var loginRequest LoginRequest
	var user *models.User
	bindingErr := c.BindJSON(&loginRequest)
	if bindingErr != nil {
		log.Println(bindingErr)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "could not parse request. check API documentation",
		})
		return
	}

	user, err, _ := h.UserRepository.FindUserAccountByUsername(loginRequest.Username)

	if err != nil {
		if err == repository.ErrUserAccountDoesNotExist {
			log.Println(bindingErr)
			c.JSON(http.StatusNotFound, gin.H{
				"message": "could not find user with this username",
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to get user",
			})
			return
		}
	}

	userPass := []byte(user.Password)
	dbPass := []byte(loginRequest.Password)

	passErr := bcrypt.CompareHashAndPassword(userPass, dbPass)
	if passErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Password don't match",
		})
		return
	}

	jwtToken, err := utils.GenerateJWT(user)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to generate jwt token",
		})
		return
	}

	log.Println(user.ID)
	c.JSON(http.StatusCreated, gin.H{
		"message": "User successfully signed in",
		"data":    newLoginResponse(jwtToken, user.ID),
	})
}

func newLoginResponse(token string, id string) *LoginResponse {
	return &LoginResponse{
		Token:  token,
		UserId: id,
	}
}