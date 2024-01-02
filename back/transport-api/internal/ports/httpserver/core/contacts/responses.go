package contacts

type contactData struct {
	ContactId    int    `json:"contact_id"`
	OwnerId      int    `json:"owner_id"`
	EmployeeId   int    `json:"employee_id"`
	FirstName    string `json:"first_name"`
	SecondName   string `json:"second_name"`
	CompanyName  string `json:"company_name"`
	Notes        string `json:"notes"`
	CreationDate uint   `json:"creation_date"`
}

type contactResponse struct {
	Data *contactData `json:"data"`
	Err  *string      `json:"error"`
}

type сontactListResponse struct {
	Data []contactData `json:"data"`
	Err  *string       `json:"error"`
}
