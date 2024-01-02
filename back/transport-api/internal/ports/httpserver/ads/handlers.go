package ads

import (
	"github.com/gin-gonic/gin"
	"transport-api/internal/app"
)

// @Summary		Добавление нового объявления
// @Description	Добавляет новое объявление
// @Tags			ads
// @Accept			json
// @Produce		json
// @Param			input	body		addAdRequest	true	"Новое объявление в JSON"
// @Success		200		{object}	adResponse		"Успешное добавление объявления"
// @Failure		500		{object}	adResponse		"Проблемы на стороне сервера"
// @Failure		400		{object}	adResponse		"Неверный формат входных данных"
// @Router			/ads [post]
func AddAd(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO implement
	}
}

// @Summary		Получение объявления
// @Description	Получает объявление по id
// @Tags			ads
// @Produce		json
// @Param			id	path		int			true	"id объявления"
// @Success		200	{object}	adResponse	"Успешное получение объявления"
// @Failure		500	{object}	adResponse	"Проблемы на стороне сервера"
// @Failure		400	{object}	adResponse	"Неверный формат входных данных"
// @Failure		404	{object}	adResponse	"Объявление не найдено"
// @Router			/ads/{id} [get]
func GetAd(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO implement
	}
}

// @Summary		Получение списка объявлений
// @Description	Получает список объявлений с использованием фильтрации и пагинации
// @Tags			ads
// @Produce		json
// @Param			limit		query		int				true	"Limit"
// @Param			offset		query		int				true	"Offset"
// @Param			name		query		string			true	"Поиск по названию/тексту"
// @Param			industry	query		int				true	"Поиск по отрасли"
// @Success		200			{object}	adListResponse	"Успешное получение объявлений"
// @Failure		500			{object}	adListResponse	"Проблемы на стороне сервера"
// @Failure		400			{object}	adListResponse	"Неверный формат входных данных"
// @Router			/ads [get]
func GetAdsList(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO implement
	}
}

// @Summary		Редактирование объявления
// @Description	Изменяет одно или несколько полей объявления
// @Tags			ads
// @Accept			json
// @Produce		json
// @Param			id		path		int				true	"id объявления"
// @Param			input	body		updateAdRequest	true	"Новые поля"
// @Success		200		{object}	adResponse		"Успешное обновление данных"
// @Failure		500		{object}	adResponse		"Проблемы на стороне сервера"
// @Failure		400		{object}	adResponse		"Неверный формат входных данных"
// @Failure		404		{object}	adResponse		"Объявление не найдено"
// @Router			/ads/{id} [put]
func UpdateAd(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO implement
	}
}

// @Summary		Удаление объявления
// @Description	Безвозвратно удаляет объявление
// @Tags			ads
// @Produce		json
// @Param			id	path		int			true	"id объявления"
// @Success		200	{object}	adResponse	"Успешное удаление объявления"
// @Failure		500	{object}	adResponse	"Проблемы на стороне сервера"
// @Failure		400	{object}	adResponse	"Неверный формат входных данных"
// @Failure		404	{object}	adResponse	"Объявление не найдено"
// @Router			/ads/{id} [delete]
func DeleteAd(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO implement
	}
}
