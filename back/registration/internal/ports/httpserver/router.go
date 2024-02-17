package httpserver

import (
	"github.com/gin-gonic/gin"
	_ "registration/docs"
	"registration/internal/app"
)

func appRouter(r *gin.RouterGroup, a app.App) {
	r.POST("register", addCompanyWithOwner(a))
}
