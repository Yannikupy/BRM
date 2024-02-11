package employees

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"transport-api/internal/app"
	"transport-api/internal/model"
	"transport-api/internal/model/core"
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
		// TODO заменить header на auth и добавить поддержку 401 unauthorized
		companyId, err := strconv.ParseUint(c.GetHeader("company_id"), 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}
		ownerId, err := strconv.ParseUint(c.GetHeader("employee_id"), 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}

		var req addEmployeeRequest
		if err = c.BindJSON(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}

		employee, err := a.CreateEmployee(c, uint(companyId), uint(ownerId), core.Employee{
			Id:           0,
			CompanyId:    uint(req.CompanyId),
			FirstName:    req.FirstName,
			SecondName:   req.SecondName,
			Email:        req.Email,
			JobTitle:     req.JobTitle,
			Department:   req.Department,
			CreationDate: 0,
			IsDeleted:    false,
		})

		switch {
		case err == nil:
			data := employeeToEmployeeData(employee)
			c.JSON(http.StatusOK, employeeResponse{
				Data: &data,
				Err:  nil,
			})
		case errors.Is(err, model.ErrCompanyNotExists):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrCompanyNotExists))
		case errors.Is(err, model.ErrEmployeeNotExists):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrEmployeeNotExists))
		case errors.Is(err, model.ErrAuthorization):
			c.AbortWithStatusJSON(http.StatusForbidden, errorResponse(model.ErrAuthorization))
		case errors.Is(err, model.ErrCoreError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrCoreError))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrCoreUnknown))
		}
	}
}

// @Summary		Получение списка сотрудников
// @Description	Получает список сотрудников компании с использованием фильтрации и пагинации
// @Tags			core/employees
// @Produce		json
// @Param			limit		query		int						true	"Limit"
// @Param			offset		query		int						true	"Offset"
// @Param			name		query		string					false	"Поиск по имени/фамилии"
// @Param			jobtitle	query		string					false	"Поиск по должности"
// @Param			department	query		string					false	"Поиск по названию отдела"
// @Success		200			{object}	employeeListResponse	"Успешное получение сотрудников"
// @Failure		500			{object}	employeeListResponse	"Проблемы на стороне сервера"
// @Failure		400			{object}	employeeListResponse	"Неверный формат входных данных"
// @Router			/employees [get]
func GetEmployeesList(a app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO заменить header на auth и добавить поддержку 401 unauthorized
		companyId, err := strconv.ParseUint(c.GetHeader("company_id"), 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}
		employeeId, err := strconv.ParseUint(c.GetHeader("employee_id"), 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}

		limit, err := strconv.Atoi(c.Query("limit"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}
		offset, err := strconv.Atoi(c.Query("offset"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}

		var employees []core.Employee

		if pattern, byName := c.GetQuery("name"); byName {
			employees, err = a.GetEmployeeByName(c,
				uint(companyId),
				uint(employeeId),
				core.EmployeeByName{
					Pattern: pattern,
					Limit:   limit,
					Offset:  offset,
				},
			)
		} else {
			var filter core.FilterEmployee
			filter.Limit = limit
			filter.Offset = offset
			filter.JobTitle, filter.ByJobTitle = c.GetQuery("jobtitle")
			filter.Department, filter.ByDepartment = c.GetQuery("department")
			employees, err = a.GetCompanyEmployees(c,
				uint(companyId),
				uint(employeeId),
				filter)
		}

		switch {
		case err == nil:
			c.JSON(http.StatusOK, employeeListResponse{
				Data: employeesToEmployeeDataList(employees),
				Err:  nil,
			})
		case errors.Is(err, model.ErrCompanyNotExists):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrCompanyNotExists))
		case errors.Is(err, model.ErrEmployeeNotExists):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrEmployeeNotExists))
		case errors.Is(err, model.ErrAuthorization):
			c.AbortWithStatusJSON(http.StatusForbidden, errorResponse(model.ErrAuthorization))
		case errors.Is(err, model.ErrCoreError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrCoreError))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrCoreUnknown))
		}
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
		// TODO заменить header на auth и добавить поддержку 401 unauthorized
		companyId, err := strconv.ParseUint(c.GetHeader("company_id"), 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}
		employeeId, err := strconv.ParseUint(c.GetHeader("employee_id"), 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}

		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}

		employee, err := a.GetEmployeeById(c,
			uint(companyId),
			uint(employeeId),
			uint(id))

		switch {
		case err == nil:
			data := employeeToEmployeeData(employee)
			c.JSON(http.StatusOK, employeeResponse{
				Data: &data,
				Err:  nil,
			})
		case errors.Is(err, model.ErrCompanyNotExists):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrCompanyNotExists))
		case errors.Is(err, model.ErrEmployeeNotExists):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrEmployeeNotExists))
		case errors.Is(err, model.ErrAuthorization):
			c.AbortWithStatusJSON(http.StatusForbidden, errorResponse(model.ErrAuthorization))
		case errors.Is(err, model.ErrCoreError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrCoreError))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrCoreUnknown))
		}
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
		// TODO заменить header на auth и добавить поддержку 401 unauthorized
		companyId, err := strconv.ParseUint(c.GetHeader("company_id"), 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}
		ownerId, err := strconv.ParseUint(c.GetHeader("employee_id"), 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}

		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}

		var req updateEmployeeRequest
		if err = c.BindJSON(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}

		employee, err := a.UpdateEmployee(c,
			uint(companyId),
			uint(ownerId),
			uint(id),
			core.UpdateEmployee{
				FirstName:  req.FirstName,
				SecondName: req.SecondName,
				Email:      req.Email,
				JobTitle:   req.JobTitle,
				Department: req.Department,
			},
		)
		switch {
		case err == nil:
			data := employeeToEmployeeData(employee)
			c.JSON(http.StatusOK, employeeResponse{
				Data: &data,
				Err:  nil,
			})
		case errors.Is(err, model.ErrCompanyNotExists):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrCompanyNotExists))
		case errors.Is(err, model.ErrEmployeeNotExists):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrEmployeeNotExists))
		case errors.Is(err, model.ErrAuthorization):
			c.AbortWithStatusJSON(http.StatusForbidden, errorResponse(model.ErrAuthorization))
		case errors.Is(err, model.ErrCoreError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrCoreError))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrCoreUnknown))
		}
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
		// TODO заменить header на auth и добавить поддержку 401 unauthorized
		companyId, err := strconv.ParseUint(c.GetHeader("company_id"), 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}
		ownerId, err := strconv.ParseUint(c.GetHeader("employee_id"), 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}

		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(model.ErrInvalidInput))
			return
		}

		err = a.DeleteEmployee(c,
			uint(companyId),
			uint(ownerId),
			uint(id))
		switch {
		case err == nil:
			c.JSON(http.StatusOK, employeeResponse{
				Data: nil,
				Err:  nil,
			})
		case errors.Is(err, model.ErrCompanyNotExists):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrCompanyNotExists))
		case errors.Is(err, model.ErrEmployeeNotExists):
			c.AbortWithStatusJSON(http.StatusNotFound, errorResponse(model.ErrEmployeeNotExists))
		case errors.Is(err, model.ErrAuthorization):
			c.AbortWithStatusJSON(http.StatusForbidden, errorResponse(model.ErrAuthorization))
		case errors.Is(err, model.ErrCoreError):
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrCoreError))
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(model.ErrCoreUnknown))
		}

	}
}
