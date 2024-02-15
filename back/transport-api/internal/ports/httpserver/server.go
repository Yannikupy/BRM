package httpserver

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"transport-api/internal/app"
	"transport-api/pkg/tokenizer"
)

func New(addr string, a app.App, tkn tokenizer.Tokenizer) *http.Server {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("api/v1")
	appRouter(api, a, tkn)
	return &http.Server{
		Addr:    addr,
		Handler: router,
	}
}
