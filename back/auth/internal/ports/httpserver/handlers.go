package httpserver

import (
	"auth/internal/app"

	"github.com/gin-gonic/gin"
)

//	@Summary		Обновление токенов
//	@Description	Обновляет access и refresh-токены, старая пара становится непригодной
//	@Accept			json
//	@Produce		json
//	@Param			input	body		refreshRequest	true	"Пара токенов"
//	@Success		200		{object}	tokensResponse	"Успешное обновление токенов"
//	@Failure		500		{object}	tokensResponse	"Проблемы на стороне сервера"
//	@Failure		400		{object}	tokensResponse	"Неверный формат входных данных"
//	@Router			/refresh [post]
func refresh(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO implement
	}
}

//	@Summary		Получение токенов
//	@Description	Получает access и refresh-токены, используя аутентификацию по логину и паролю
//	@Accept			json
//	@Produce		json
//	@Param			input	body		loginRequest	true	"Логин и пароль"
//	@Success		200		{object}	tokensResponse	"Успешное получение токенов"
//	@Failure		500		{object}	tokensResponse	"Проблемы на стороне сервера"
//	@Failure		400		{object}	tokensResponse	"Неверный формат входных данных"
//	@Router			/login [post]
func login(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO implement
	}
}

//	@Summary		Выход из аккаунта
//	@Description	Удаляет пару токенов
//	@Accept			json
//	@Produce		json
//	@Param			input	body		logoutRequest	true	"Пара токенов"
//	@Success		200		{object}	tokensResponse	"Успешный выход из аккаунта"
//	@Failure		500		{object}	tokensResponse	"Проблемы на стороне сервера"
//	@Failure		400		{object}	tokensResponse	"Неверный формат входных данных"
//	@Router			/logout [post]
func logout(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO implement
	}
}
