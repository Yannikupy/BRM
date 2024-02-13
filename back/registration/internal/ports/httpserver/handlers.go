package httpserver

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"registration/internal/app"
	"registration/internal/model"
)

// @Summary		Добавление новой компании и владельца
// @Description	Добавляет новую компанию и её владельца, который является её первым сотрудником, одним запросом
// @Accept			json
// @Produce		json
// @Param			input	body		addCompanyAndOwnerRequest	true	"Информация о компании и её владельце"
// @Success		200		{object}	companyAndOwnerResponse	"Успешное добавление компании с владельцем"
// @Failure		500		{object}	companyAndOwnerResponse	"Проблемы на стороне сервера"
// @Failure		400		{object}	companyAndOwnerResponse	"Неверный формат входных данных"
// @Router			/register [post]
func addCompanyWithOwner(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req addCompanyAndOwnerRequest
		if err := c.BindJSON(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.ErrInvalidInput)
			return
		}

		company, owner, err := a.CreateCompanyAndOwner(c,
			model.Company{
				Id:           0,
				Name:         req.Company.Name,
				Description:  req.Company.Description,
				Industry:     req.Company.Industry,
				OwnerId:      0,
				Rating:       0,
				CreationDate: 0,
				IsDeleted:    false,
			},
			model.Employee{
				Id:           0,
				CompanyId:    0,
				FirstName:    req.Owner.FirstName,
				SecondName:   req.Owner.SecondName,
				Email:        req.Owner.Email,
				JobTitle:     req.Owner.JobTitle,
				Department:   req.Owner.Department,
				CreationDate: 0,
				IsDeleted:    false,
			},
		)

		switch {
		case err == nil:
			c.JSON(http.StatusOK, companyAndOwnerResponse{
				Data: &companyAndOwnerData{
					Company: companyToCompanyData(company),
					Owner:   ownerToOwnerData(owner),
				},
				Err: nil,
			})
		case errors.Is(err, model.ErrCoreError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrCoreError))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrCoreUnknown))
		}
	}
}
