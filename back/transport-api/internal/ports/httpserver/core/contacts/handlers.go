package contacts

import (
	"github.com/gin-gonic/gin"
	"transport-api/internal/app"
)

// @Summary		Добавление нового контакта
// @Description	Добавляет новый контакт в список контактов сотрудника
// @Tags			core/contacts
// @Accept			json
// @Produce		json
// @Param			input	body		addContactRequest	true	"id сотрудника, которого добавляют в контакты"
// @Success		200		{object}	contactResponse		"Успешное добавление контакта"
// @Failure		500		{object}	contactResponse		"Проблемы на стороне сервера"
// @Failure		400		{object}	contactResponse		"Неверный формат входных данных"
// @Failure		404		{object}	contactResponse		"Добавляемый сотрудник не найден"
// @Router			/contacts [post]
func AddContact(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO implement
	}
}

// @Summary		Получение списка контактов
// @Description	Получает список контактов сотрудника с использованием фильтрации и пагинации
// @Tags			core/contacts
// @Produce		json
// @Param			limit	query		int				true	"Limit"
// @Param			offset	query		int				true	"Offset"
// @Param			name	query		string			true	"Поиск по имени/фамилии"
// @Param			company	query		string			true	"Поиск по компании"
// @Success		200		{object}	contactResponse	"Успешное получение контактов"
// @Failure		500		{object}	contactResponse	"Проблемы на стороне сервера"
// @Failure		400		{object}	contactResponse	"Неверный формат входных данных"
// @Router			/contacts [get]
func GetContactsList(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO implement
	}
}

// @Summary		Получение контакта
// @Description	Получает контакт по id
// @Tags			core/contacts
// @Produce		json
// @Param			id	path		int				true	"id контакта"
// @Success		200	{object}	contactResponse	"Успешное получение контакта"
// @Failure		500	{object}	contactResponse	"Проблемы на стороне сервера"
// @Failure		400	{object}	contactResponse	"Неверный формат входных данных"
// @Failure		404	{object}	contactResponse	"Контакт не найден"
// @Router			/contacts/{id} [get]
func GetContact(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO implement
	}
}

// @Summary		Редактирование контакта
// @Description	Изменяет одно или несколько полей контакта
// @Tags			core/contacts
// @Accept			json
// @Produce		json
// @Param			id		path		int						true	"id контакта"
// @Param			input	body		updateContactRequest	true	"Новые поля"
// @Success		200		{object}	contactResponse			"Успешное обновление данных"
// @Failure		500		{object}	contactResponse			"Проблемы на стороне сервера"
// @Failure		400		{object}	contactResponse			"Неверный формат входных данных"
// @Failure		404		{object}	contactResponse			"Контакт не найден"
// @Router			/contacts/{id} [put]
func UpdateContact(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO implement
	}
}

// @Summary		Удаление контакта
// @Description	Безвозвратно удаляет контакт и все его поля
// @Tags			core/contacts
// @Produce		json
// @Param			id	path		int				true	"id контакта"
// @Success		200	{object}	contactResponse	"Успешное удаление контакта"
// @Failure		500	{object}	contactResponse	"Проблемы на стороне сервера"
// @Failure		400	{object}	contactResponse	"Неверный формат входных данных"
// @Failure		404	{object}	contactResponse	"Контакт не найден"
// @Router			/contacts/{id} [delete]
func DeleteContact(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO implement
	}
}
