package companies

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"transport-api/internal/app"
	"transport-api/internal/model"
	"transport-api/internal/model/core"
)

// @Summary		Создание компании и её владельца
// @Description	Возвращает созданную компанию и её владельца
// @Tags			core/companies
// @Produce		json
// @Param input body addCompanyAndOwnerRequest true "Информация о новой компании и её владельце"
// @Success		200	{object}	companyAndOwnerResponse	"Успешное создание компании и владельца"
// @Failure		500	{object}	companyAndOwnerResponse	"Проблемы на стороне сервера"
// @Failure		400	{object}	companyAndOwnerResponse	"Неверный формат входных данных"
// @Router			/companies [post]
func AddCompanyAndOwner(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req addCompanyAndOwnerRequest
		if err := c.BindJSON(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}

		company, owner, err := a.CreateCompanyAndOwner(c,
			core.Company{
				Id:           0,
				Name:         req.Company.Name,
				Description:  req.Company.Description,
				Industry:     req.Company.Industry,
				OwnerId:      0,
				Rating:       0,
				CreationDate: 0,
				IsDeleted:    false,
			},
			core.Employee{
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
		c.JSON(http.StatusOK, mainPageResponse{
			Data: &mainPageData{
				Title: "Какаято компания",
				Stats: mainPageStatsData{
					ActiveLeadsAmount:     20,
					ActiveLeadsPrice:      30,
					ClosedLeadsAmount:     40,
					ClosedLeadsPrice:      50,
					ActiveAdsAmount:       60,
					CompanyAbsoluteRating: 4.9,
					CompanyRelativeRating: 0.99,
				},
			},
			Err: nil,
		})
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
		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}

		company, err := a.GetCompany(c, uint(id))
		switch {
		case err == nil:
			data := companyToCompanyData(company)
			c.JSON(http.StatusOK, companyResponse{
				Data: &data,
				Err:  nil,
			})
		case errors.Is(err, model.ErrCompanyNotExists):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrCompanyNotExists))
		case errors.Is(err, model.ErrCoreError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrCoreError))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrCoreUnknown))
		}
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
		companyId, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}

		// TODO заменить header на auth и добавить поддержку 401 unauthorized
		ownerId, err := strconv.ParseUint(c.GetHeader("employee_id"), 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}

		var req updateCompanyRequest
		if err = c.BindJSON(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}

		company, err := a.UpdateCompany(c, uint(companyId), uint(ownerId), core.UpdateCompany{
			Name:        req.Name,
			Description: req.Description,
			Industry:    req.Industry,
			OwnerId:     req.OwnerId,
		})

		switch {
		case err == nil:
			data := companyToCompanyData(company)
			c.JSON(http.StatusOK, companyResponse{
				Data: &data,
				Err:  nil,
			})
		case errors.Is(err, model.ErrCompanyNotExists):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrCompanyNotExists))
		case errors.Is(err, model.ErrAuthorization):
			c.AbortWithStatusJSON(http.StatusForbidden, errorResponse(model.ErrAuthorization))
		case errors.Is(err, model.ErrCoreError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrCoreError))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrCoreUnknown))
		}
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
		companyId, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}

		// TODO заменить header на auth и добавить поддержку 401 unauthorized
		ownerId, err := strconv.ParseUint(c.GetHeader("employee_id"), 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}

		err = a.DeleteCompany(c, uint(companyId), uint(ownerId))
		switch {
		case err == nil:
			c.JSON(http.StatusOK, companyResponse{
				Data: nil,
				Err:  nil,
			})
		case errors.Is(err, model.ErrCompanyNotExists):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrCompanyNotExists))
		case errors.Is(err, model.ErrAuthorization):
			c.AbortWithStatusJSON(http.StatusForbidden, errorResponse(model.ErrAuthorization))
		case errors.Is(err, model.ErrCoreError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrCoreError))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrCoreUnknown))
		}
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
		c.JSON(http.StatusOK, industryResponse{
			Industries: []string{
				"802",
				"804",
				"805",
				"806",
			},
			Err: nil,
		})
	}
}
