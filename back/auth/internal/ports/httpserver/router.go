package httpserver

import (
	"auth/internal/app"
	"auth/pkg/logger"

	"github.com/gin-gonic/gin"

	_ "auth/docs"
)

func appRouter(r *gin.RouterGroup, a app.App, logs logger.Logger) {
	r.Use(corsMiddleware())
	r.Use(panicMiddleware(logs))
	r.Use(logMiddleware(logs))

	r.POST("refresh", refresh(a))
	r.POST("login", login(a))
	r.POST("logout", logout(a))
}
