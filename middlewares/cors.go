package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "Cache-Control"}
	return cors.New(config)
}
