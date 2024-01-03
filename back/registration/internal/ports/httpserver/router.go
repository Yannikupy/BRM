package httpserver

import (
	"github.com/gin-gonic/gin"
	"registration/internal/app"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "registration/docs"
)

func appRouter(r *gin.RouterGroup, a app.App) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("", addCompanyWithOwner(a))
}
