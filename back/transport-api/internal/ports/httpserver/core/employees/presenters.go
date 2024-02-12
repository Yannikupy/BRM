package employees

type addEmployeeRequest struct {
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	CompanyId  int    `json:"company_id"`
	JobTitle   string `json:"job_title"`
	Department string `json:"department"`
	Login      string `json:"login"`
	Password   string `json:"password"`
}

type updateEmployeeRequest struct {
	EmployeeId int    `json:"employee_id"`
	JobTitle   string `json:"job_title"`
	Department string `json:"department"`
}
