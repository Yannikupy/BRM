package httpserver

import (
	"auth/internal/app"
	"auth/pkg/logger"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"

	"github.com/gin-gonic/gin"
)

func New(addr string, a app.App, logs logger.Logger) *http.Server {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("api/v1/auth")
	appRouter(api, a, logs)

	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}
