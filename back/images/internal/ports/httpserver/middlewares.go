package httpserver

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"images/internal/model"
	"images/pkg/logger"
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

func corsMiddleware(originAddr string) gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{originAddr}
	return cors.New(config)
}
