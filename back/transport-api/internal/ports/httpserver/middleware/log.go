package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"sync"
	"transport-api/pkg/logger"
)

// requestsCounter нужен для сопоставления логов о запросе и ответе на этот запрос
var requestsCounter uint64 = 0
var mu sync.Mutex

func Log(logs logger.Logger) gin.HandlerFunc {
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
