package httpserver

import (
	"auth/internal/app"
	"net/http"

	"github.com/gin-gonic/gin"
)

func New(addr string, a app.App) *http.Server {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	api := router.Group("api/v1/auth")
	appRouter(api, a)
	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}
