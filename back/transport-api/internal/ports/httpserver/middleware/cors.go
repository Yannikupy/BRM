package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors(originAddr string) gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{originAddr}
	return cors.New(config)
}
