package httpserver

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"registration/internal/app"
)

func New(addr string, a app.App) *http.Server {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("api/v1")
	appRouter(api, a)

	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}
