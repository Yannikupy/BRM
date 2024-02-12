package employees

type employeeData struct {
	CompanyId    int    `json:"company_id"`
	EmployeeId   int    `json:"employee_id"`
	FirstName    string `json:"first_name"`
	SecondName   string `json:"second_name"`
	JobTitle     string `json:"job_title"`
	Department   string `json:"department"`
	CreationDate uint   `json:"creation_date"`
}

type employeeResponse struct {
	Data *employeeData `json:"data"`
	Err  *string       `json:"error"`
}

type employeeListResponse struct {
	Data []employeeData `json:"data"`
	Err  *string        `json:"error"`
}
