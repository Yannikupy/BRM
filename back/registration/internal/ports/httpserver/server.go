package httpserver

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"registration/internal/app"
)

func New(addr string, a app.App) *http.Server {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	api := router.Group("api/v1")
	appRouter(api, a)
	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}
