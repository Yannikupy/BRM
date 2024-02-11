package companies

type addCompanyAndOwnerRequest struct {
	Company addCompanyData `json:"company"`
	Owner   addOwnerData   `json:"owner"`
}

type addCompanyData struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Industry    uint   `json:"industry"`
}

type addOwnerData struct {
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Email      string `json:"email"`
	JobTitle   string `json:"job_title"`
	Department string `json:"department"`
}

type updateCompanyRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Industry    uint   `json:"industry"`
	OwnerId     uint   `json:"owner_id"`
}
