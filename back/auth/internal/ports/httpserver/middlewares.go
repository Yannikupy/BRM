package httpserver

import (
	"auth/internal/model"
	"auth/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

// requestsCounter нужен для сопоставления логов о запросе и ответе на этот запрос
var requestsCounter uint64 = 0
var mu sync.Mutex

func logMiddleware(logs logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		mu.Lock()
		requestId := requestsCounter
		requestsCounter++
		mu.Unlock()

		logs.Info(logger.Fields{
			"path":   c.Request.URL.Path,
			"method": c.Request.Method,
		}, fmt.Sprintf("incoming request #%d", requestId))
		c.Next()

		logs.Info(logger.Fields{
			"path":   c.Request.URL.Path,
			"method": c.Request.Method,
			"status": c.Writer.Status(),
		}, fmt.Sprintf("response to #%d", requestId))
	}
}

func panicMiddleware(logs logger.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				logs.Error(logger.Fields{
					"path":   c.Request.URL.Path,
					"method": c.Request.Method,
				}, fmt.Sprintf("panic: %v", r))

				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"data": nil, "error": model.ErrServiceError.Error()})
			}
		}()
		c.Next()
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	}
}
