package middlewares

import (
	"fx-bank/internal/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "no token found in authorization header"})
			return
		}
		if !strings.HasPrefix(tokenString, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "invalid format in authorization header"})
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		token, claims, err := utils.VerifyJWT(tokenString)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "token contains an invalid number of segments"})
			return
		}

		if !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("user_id", claims.UserID)

	}
}
