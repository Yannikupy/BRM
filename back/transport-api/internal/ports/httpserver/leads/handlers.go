package leads

import (
	"github.com/gin-gonic/gin"
	"transport-api/internal/app"
)

// @Summary		Добавление новой сделки
// @Description	Добавляет новую сделку
// @Tags			leads
// @Accept			json
// @Produce		json
// @Param			input	body		addLeadRequest	true	"Новая сделка в JSON"
// @Success		200		{object}	leadResponse	"Успешное добавление задачи"
// @Failure		500		{object}	leadResponse	"Проблемы на стороне сервера"
// @Failure		400		{object}	leadResponse	"Неверный формат входных данных"
// @Router			/leads [post]
func AddLead(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO implement
	}
}

// @Summary		Получение сделки
// @Description	Получает сделку по id
// @Tags			leads
// @Produce		json
// @Param			id	path		int				true	"id сделки"
// @Success		200	{object}	leadResponse	"Успешное получение сделки"
// @Failure		500	{object}	leadResponse	"Проблемы на стороне сервера"
// @Failure		400	{object}	leadResponse	"Неверный формат входных данных"
// @Failure		404	{object}	leadResponse	"Задача не найдена"
// @Router			/leads/{id} [get]
func GetLead(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO implement
	}
}

// @Summary		Получение списка сделок
// @Description	Получает список сделок с использованием фильтрации и пагинации
// @Tags			leads
// @Produce		json
// @Param			limit	query		int					true	"Limit"
// @Param			offset	query		int					true	"Offset"
// @Param			name	query		string				true	"Поиск по названию/тексту"
// @Param			stage	query		int					true	"Поиск по этапу"
// @Success		200		{object}	leadListResponse	"Успешное получение сделок"
// @Failure		500		{object}	leadListResponse	"Проблемы на стороне сервера"
// @Failure		400		{object}	leadListResponse	"Неверный формат входных данных"
// @Router			/leads [get]
func GetLeadsList(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO implement
	}
}

// @Summary		Редактирование сделки
// @Description	Изменяет одно или несколько полей сделки
// @Tags			leads
// @Accept			json
// @Produce		json
// @Param			id		path		int					true	"id сделки"
// @Param			input	body		updateLeadRequest	true	"Новые поля"
// @Success		200		{object}	leadResponse		"Успешное обновление данных"
// @Failure		500		{object}	leadResponse		"Проблемы на стороне сервера"
// @Failure		400		{object}	leadResponse		"Неверный формат входных данных"
// @Failure		404		{object}	leadResponse		"Сделка не найдена"
// @Router			/leads/{id} [put]
func UpdateLead(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO implement
	}
}

// @Summary		Удаление сделки
// @Description	Безвозвратно удаляет сделку
// @Tags			leads
// @Produce		json
// @Param			id	path		int				true	"id сделки"
// @Success		200	{object}	leadResponse	"Успешное удаление сделки"
// @Failure		500	{object}	leadResponse	"Проблемы на стороне сервера"
// @Failure		400	{object}	leadResponse	"Неверный формат входных данных"
// @Failure		404	{object}	leadResponse	"Сделка не найдена"
// @Router			/leads/{id} [delete]
func DeleteLead(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO implement
	}
}

// @Summary		Получение этапов сделки
// @Description	Возвращает словарь из этапов и их id
// @Tags			leads
// @Produce		json
// @Success		200	{object}	stageResponse	"Успешное получение данных"
// @Failure		500	{object}	stageResponse	"Проблемы на стороне сервера"
// @Failure		400	{object}	stageResponse	"Неверный формат входных данных"
// @Router			/leads/stages [get]
func GetStagesMap(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO implement
	}
}
