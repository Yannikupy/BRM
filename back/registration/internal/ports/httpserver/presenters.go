package httpserver

type addCompanyRequest struct {
	Title    string `json:"title"`
	Industry int    `json:"industry"`
}

type addOwnerRequest struct {
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	JobTitle   string `json:"job_title"`
	Department string `json:"department"`
	Login      string `json:"login"`
	Password   string `json:"password"`
}

type addCompanyWithOwnerRequest struct {
	CompanyData *addCompanyRequest `json:"company_data"`
	OwnerData   *addOwnerRequest   `json:"owner_data"`
}
