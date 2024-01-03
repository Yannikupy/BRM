package httpserver

import (
	"github.com/gin-gonic/gin"
	"registration/internal/app"
)

// @Summary		Добавление новой компании и владельца
// @Description	Добавляет новую компанию и её владельца, который является её первым сотрудником, одним запросом
// @Accept			json
// @Produce		json
// @Param			input	body		addCompanyWithOwnerRequest	true	"Информация о компании и её владельце"
// @Success		200		{object}	addCompanyAndOwnerResponse	"Успешное добавление компании с владельцем"
// @Failure		500		{object}	addCompanyAndOwnerResponse	"Проблемы на стороне сервера"
// @Failure		400		{object}	addCompanyAndOwnerResponse	"Неверный формат входных данных"
// @Router			/registration [post]
func addCompanyWithOwner(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO implement
	}
}
