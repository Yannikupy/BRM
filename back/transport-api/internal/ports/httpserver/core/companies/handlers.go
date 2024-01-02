package companies

import (
	"github.com/gin-gonic/gin"
	"transport-api/internal/app"
)

// @Summary		Получение информации о компании
// @Description	Возвращает название и статистику компании для главной страницы
// @Tags			core/companies
// @Produce		json
// @Param			id	path		int					true	"id компании"
// @Success		200	{object}	mainPageResponse	"Успешное получение данных"
// @Failure		500	{object}	mainPageResponse	"Проблемы на стороне сервера"
// @Failure		400	{object}	mainPageResponse	"Неверный формат входных данных"
// @Failure		404	{object}	mainPageResponse	"Компания не найдена"
// @Router			/companies/mainpage/{id} [get]
func GetCompanyMainPage(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO implement
	}
}

// @Summary		Получение информации о компании
// @Description	Возвращает поля компании для страницы редактирования
// @Tags			core/companies
// @Produce		json
// @Param			id	path		int				true	"id компании"
// @Success		200	{object}	companyResponse	"Успешное получение данных"
// @Failure		500	{object}	companyResponse	"Проблемы на стороне сервера"
// @Failure		400	{object}	companyResponse	"Неверный формат входных данных"
// @Failure		404	{object}	companyResponse	"Компания не найдена"
// @Router			/companies/{id} [get]
func GetCompany(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO implement
	}
}

// @Summary		Редактирование полей компании
// @Description	Изменяет одно или несколько полей компании
// @Tags			core/companies
// @Produce		json
// @Param			id		path		int						true	"id компании"
// @Param			input	body		updateCompanyRequest	true	"Новые поля"
// @Success		200		{object}	companyResponse			"Успешное обновление данных"
// @Failure		500		{object}	companyResponse			"Проблемы на стороне сервера"
// @Failure		400		{object}	companyResponse			"Неверный формат входных данных"
// @Failure		404		{object}	companyResponse			"Компания не найдена"
// @Router			/companies/{id} [put]
func UpdateCompany(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO implement
	}
}

// @Summary		Удаление компании
// @Description	Безвозвратно удаляет компанию и всё, что с этой компанией связано
// @Tags			core/companies
// @Produce		json
// @Param			id	path		int				true	"id компании"
// @Success		200	{object}	companyResponse	"Успешное удаление компании"
// @Failure		500	{object}	companyResponse	"Проблемы на стороне сервера"
// @Failure		400	{object}	companyResponse	"Неверный формат входных данных"
// @Failure		404	{object}	companyResponse	"Компания не найдена"
// @Router			/companies/{id} [delete]
func DeleteCompany(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO implement
	}
}

// @Summary		Получение отраслей
// @Description	Возвращает словарь из отраслей и их id
// @Tags			core/companies
// @Produce		json
// @Success		200	{object}	industryResponse	"Успешное получение данных"
// @Failure		500	{object}	industryResponse	"Проблемы на стороне сервера"
// @Failure		400	{object}	industryResponse	"Неверный формат входных данных"
// @Router			/companies/industries [get]
func GetIndustriesMap(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO implement
	}
}
