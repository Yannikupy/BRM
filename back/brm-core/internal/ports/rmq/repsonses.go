package rmq

type createCompanyAndOwnerResponse struct {
	Company companyData  `json:"company"`
	Owner   employeeData `json:"owner"`
}

type updateCompanyResponse struct {
	Company companyData `json:"company"`
}

type createEmployeeResponse struct {
	Employee employeeData `json:"employee"`
}

type updateEmployeeResponse struct {
	Employee employeeData `json:"employee"`
}

type createContactResponse struct {
	Contact contactData `json:"contact"`
}

type updateContactResponse struct {
	Contact contactData `json:"contact"`
}
