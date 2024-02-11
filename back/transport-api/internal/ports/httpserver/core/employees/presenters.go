package employees

type addEmployeeRequest struct {
	CompanyId  int    `json:"company_id"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Email      string `json:"email"`
	JobTitle   string `json:"job_title"`
	Department string `json:"department"`
}

type updateEmployeeRequest struct {
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Email      string `json:"email"`
	JobTitle   string `json:"job_title"`
	Department string `json:"department"`
}
