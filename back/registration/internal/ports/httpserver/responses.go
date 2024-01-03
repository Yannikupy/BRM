package httpserver

type companyData struct {
	Title     string `json:"title"`
	Industry  string `json:"industry"`
	OwnerId   int    `json:"owner_id"`
	CreatedAt uint   `json:"created_at"`
}

type ownerData struct {
	CompanyId    int    `json:"company_id"`
	EmployeeId   int    `json:"employee_id"`
	FirstName    string `json:"first_name"`
	SecondName   string `json:"second_name"`
	JobTitle     string `json:"job_title"`
	Department   string `json:"department"`
	CreationDate uint   `json:"creation_date"`
}

type addCompanyAndOwnerData struct {
	CompanyData *companyData `json:"company_data"`
	OwnerData   *ownerData   `json:"owner_data"`
}

type addCompanyAndOwnerResponse struct {
	Data *addCompanyAndOwnerData `json:"data"`
	Err  *string                 `json:"error"`
}
