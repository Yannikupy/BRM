package tasks

import (
	"github.com/gin-gonic/gin"
	"transport-api/internal/app"
)

// @Summary		Добавление новой задачи
// @Description	Добавляет новую задачу
// @Tags			tasks
// @Accept			json
// @Produce		json
// @Param			input	body		addTaskRequest	true	"Новая задача в JSON"
// @Success		200		{object}	taskResponse	"Успешное добавление задачи"
// @Failure		500		{object}	taskResponse	"Проблемы на стороне сервера"
// @Failure		400		{object}	taskResponse	"Неверный формат входных данных"
// @Router			/tasks [post]
func AddTask(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO implement
	}
}

// @Summary		Получение списка задач
// @Description	Получает список задач с использованием фильтрации и пагинации
// @Tags			tasks
// @Produce		json
// @Param			limit	query		int					true	"Limit"
// @Param			offset	query		int					true	"Offset"
// @Param			name	query		string				true	"Поиск по названию/тексту"
// @Param			stage	query		int					true	"Поиск по этапу"
// @Success		200		{object}	taskListResponse	"Успешное получение задач"
// @Failure		500		{object}	taskListResponse	"Проблемы на стороне сервера"
// @Failure		400		{object}	taskListResponse	"Неверный формат входных данных"
// @Router			/tasks [get]
func GetTasksList(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO implement
	}
}

// @Summary		Получение задачи
// @Description	Получает задачу по id
// @Tags			tasks
// @Produce		json
// @Param			id	path		int				true	"id задачи"
// @Success		200	{object}	taskResponse	"Успешное получение задачи"
// @Failure		500	{object}	taskResponse	"Проблемы на стороне сервера"
// @Failure		400	{object}	taskResponse	"Неверный формат входных данных"
// @Failure		404	{object}	taskResponse	"Задача не найдена"
// @Router			/tasks/{id} [get]
func GetTask(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO implement
	}
}

// @Summary		Редактирование задачи
// @Description	Изменяет одно или несколько полей задачи
// @Tags			tasks
// @Accept			json
// @Produce		json
// @Param			id		path		int					true	"id задачи"
// @Param			input	body		updateTaskRequest	true	"Новые поля"
// @Success		200		{object}	taskResponse		"Успешное обновление данных"
// @Failure		500		{object}	taskResponse		"Проблемы на стороне сервера"
// @Failure		400		{object}	taskResponse		"Неверный формат входных данных"
// @Failure		404		{object}	taskResponse		"Задача не найдена"
// @Router			/tasks/{id} [put]
func UpdateTask(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO implement
	}
}

// @Summary		Удаление задачи
// @Description	Безвозвратно удаляет задачу
// @Tags			tasks
// @Produce		json
// @Param			id	path		int				true	"id задачи"
// @Success		200	{object}	taskResponse	"Успешное удаление задачи"
// @Failure		500	{object}	taskResponse	"Проблемы на стороне сервера"
// @Failure		400	{object}	taskResponse	"Неверный формат входных данных"
// @Failure		404	{object}	taskResponse	"Задача не найдена"
// @Router			/tasks/{id} [delete]
func DeleteTask(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO implement
	}
}

// @Summary		Получение этапов задачи
// @Description	Возвращает словарь из этапов и их id
// @Tags			tasks
// @Produce		json
// @Success		200	{object}	stageResponse	"Успешное получение данных"
// @Failure		500	{object}	stageResponse	"Проблемы на стороне сервера"
// @Failure		400	{object}	stageResponse	"Неверный формат входных данных"
// @Router			/tasks/stages [get]
func GetStagesMap(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO implement
	}
}
