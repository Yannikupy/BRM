package httpserver

import (
	"auth/internal/app"

	"github.com/gin-gonic/gin"

	_ "auth/docs"
)

func appRouter(r *gin.RouterGroup, a app.App) {
	r.POST("refresh", refresh(a))
	r.POST("login", login(a))
	r.POST("logout", logout(a))
}
