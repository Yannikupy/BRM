package employees

import (
	"github.com/gin-gonic/gin"
	"transport-api/internal/app"
)

// @Summary		Добавление нового сотрудника
// @Description	Добавляет нового сотрудника в компанию
// @Tags			core/employees
// @Accept			json
// @Produce		json
// @Param			input	body		addEmployeeRequest	true	"id сотрудника, которого добавляют в контакты"
// @Success		200		{object}	employeeResponse	"Успешное добавление сотрудника"
// @Failure		500		{object}	employeeResponse	"Проблемы на стороне сервера"
// @Failure		400		{object}	employeeResponse	"Неверный формат входных данных"
// @Router			/employees [post]
func AddEmployee(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO implement
	}
}

// @Summary		Получение списка сотрудников
// @Description	Получает список сотрудников компании с использованием фильтрации и пагинации
// @Tags			core/employees
// @Produce		json
// @Param			limit		query		int						true	"Limit"
// @Param			offset		query		int						true	"Offset"
// @Param			name		query		string					true	"Поиск по имени/фамилии"
// @Param			jobtitle	query		string					true	"Поиск по должности"
// @Param			department	query		string					true	"Поиск по названию отдела"
// @Success		200			{object}	employeeListResponse	"Успешное получение сотрудников"
// @Failure		500			{object}	employeeListResponse	"Проблемы на стороне сервера"
// @Failure		400			{object}	employeeListResponse	"Неверный формат входных данных"
// @Router			/employees [get]
func GetEmployeesList(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO implement
	}
}

// @Summary		Получение сотрудника
// @Description	Получает сотрудника по id
// @Tags			core/employees
// @Produce		json
// @Param			id	path		int					true	"id сотрудника"
// @Success		200	{object}	employeeResponse	"Успешное получение сотрудника"
// @Failure		500	{object}	employeeResponse	"Проблемы на стороне сервера"
// @Failure		400	{object}	employeeResponse	"Неверный формат входных данных"
// @Failure		404	{object}	employeeResponse	"Сотрудник не найден"
// @Router			/employees/{id} [get]
func GetEmployee(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO implement
	}
}

// @Summary		Редактирование сотрудника
// @Description	Изменяет одно или несколько полей сотрудника
// @Tags			core/employees
// @Accept			json
// @Produce		json
// @Param			id		path		int						true	"id сотрудника"
// @Param			input	body		updateEmployeeRequest	true	"Новые поля"
// @Success		200		{object}	employeeResponse		"Успешное обновление данных"
// @Failure		500		{object}	employeeResponse		"Проблемы на стороне сервера"
// @Failure		400		{object}	employeeResponse		"Неверный формат входных данных"
// @Failure		404		{object}	employeeResponse		"Контакт не найден"
// @Router			/employees/{id} [put]
func UpdateEmployee(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO implement
	}
}

// @Summary		Удаление сотрудника
// @Description	Безвозвратно удаляет сотрудника и все его поля
// @Tags			core/employees
// @Produce		json
// @Param			id	path		int					true	"id сотрудника"
// @Success		200	{object}	employeeResponse	"Успешное удаление контакта"
// @Failure		500	{object}	employeeResponse	"Проблемы на стороне сервера"
// @Failure		400	{object}	employeeResponse	"Неверный формат входных данных"
// @Failure		404	{object}	employeeResponse	"Контакт не найден"
// @Router			/employees/{id} [delete]
func DeleteEmployee(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO implement
	}
}
