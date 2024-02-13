package httpserver

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "transport-api/docs"
	"transport-api/internal/app"
	"transport-api/internal/ports/httpserver/core/companies"
	"transport-api/internal/ports/httpserver/core/contacts"
	"transport-api/internal/ports/httpserver/core/employees"
)

func appRouter(r *gin.RouterGroup, a app.App) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/companies/:id", companies.GetCompany(a))
	r.GET("/companies/:id/mainpage", companies.GetCompanyMainPage(a))
	r.PUT("/companies/:id", companies.UpdateCompany(a))
	r.DELETE("/companies/:id", companies.DeleteCompany(a))
	r.GET("/companies/industries", companies.GetIndustriesMap(a))

	r.POST("/employees", employees.AddEmployee(a))
	r.GET("/employees/:id", employees.GetEmployee(a))
	r.GET("/employees", employees.GetEmployeesList(a))
	r.PUT("/employees/:id", employees.UpdateEmployee(a))
	r.DELETE("/employees/:id", employees.DeleteEmployee(a))

	r.POST("/contacts", contacts.AddContact(a))
	r.GET("/contacts/:id", contacts.GetContact(a))
	r.GET("/contacts", contacts.GetContactsList(a))
	r.PUT("/contacts/:id", contacts.UpdateContact(a))
	r.DELETE("/contacts/:id", contacts.DeleteContact(a))

	//r.POST("/tasks", tasks.AddTask(a))
	//r.GET("/tasks/:id", tasks.GetTask(a))
	//r.GET("/tasks", tasks.GetTasksList(a))
	//r.PUT("/tasks/:id", tasks.UpdateTask(a))
	//r.DELETE("/tasks/:id", tasks.DeleteTask(a))
	//r.GET("/tasks/stages", tasks.GetStagesMap(a))
	//
	//r.POST("/ads", ads.AddAd(a))
	//r.GET("/ads/:id", ads.GetAd(a))
	//r.GET("/ads", ads.GetAdsList(a))
	//r.PUT("/ads/:id", ads.UpdateAd(a))
	//r.DELETE("/ads/:id", ads.DeleteAd(a))
	//
	//r.POST("/leads", leads.AddLead(a))
	//r.GET("/leads/:id", leads.GetLead(a))
	//r.GET("/leads", leads.GetLeadsList(a))
	//r.PUT("/leads/:id", leads.UpdateLead(a))
	//r.DELETE("/leads/:id", leads.DeleteLead(a))
	//r.GET("/leads/stages", leads.GetStagesMap(a))
}
