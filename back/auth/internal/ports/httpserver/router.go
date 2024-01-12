package httpserver

import (
	"auth/internal/app"

	"github.com/gin-gonic/gin"

	_ "auth/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func appRouter(r *gin.RouterGroup, a app.App) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("refresh", refresh(a))
	r.POST("login", login(a))
	r.POST("logout", logout(a))
}
