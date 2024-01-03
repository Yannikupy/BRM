package httpserver

import (
	"auth/internal/app"
	"github.com/gin-gonic/gin"
)

// @Summary		Обновление токенов
// @Description	Обновляет access и refresh-токены, старая пара становится непригодной
// @Accept			json
// @Produce		json
// @Param			input	body		refreshRequest	true	"Refresh-токен"
// @Success		200		{object}	tokensResponse	"Успешное обновление токенов"
// @Failure		500		{object}	tokensResponse	"Проблемы на стороне сервера"
// @Failure		400		{object}	tokensResponse	"Неверный формат входных данных"
// @Router			/refresh [post]
func refresh(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO implement
	}
}

// @Summary		Получение токенов
// @Description	Получает access и refresh-токены, используя авторизацию по логину и паролю
// @Accept			json
// @Produce		json
// @Param			input	body		authorizeRequest	true	"Логин и пароль"
// @Success		200		{object}	tokensResponse		"Успешное получение токенов"
// @Failure		500		{object}	tokensResponse		"Проблемы на стороне сервера"
// @Failure		400		{object}	tokensResponse		"Неверный формат входных данных"
// @Router			/authorize [post]
func authorize(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO implement
	}
}
