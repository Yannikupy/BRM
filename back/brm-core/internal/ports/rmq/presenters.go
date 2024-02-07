package rmq

type companyData struct {
	Id           uint    `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Industry     uint    `json:"industry"`
	OwnerId      uint    `json:"owner_id"`
	Rating       float64 `json:"rating"`
	CreationDate int64   `json:"creation_date"`
	IsDeleted    bool    `json:"is_deleted"`
}

type updateCompanyData struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Industry    uint   `json:"industry"`
	OwnerId     uint   `json:"owner_id"`
}

type employeeData struct {
	Id           uint   `json:"id"`
	CompanyId    uint   `json:"company_id"`
	FirstName    string `json:"first_name"`
	SecondName   string `json:"second_name"`
	Email        string `json:"email"`
	JobTitle     string `json:"job_title"`
	Department   string `json:"department"`
	CreationDate int64  `json:"creation_date"`
	IsDeleted    bool   `json:"is_deleted"`
}

type updateEmployeeData struct {
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Email      string `json:"email"`
	JobTitle   string `json:"job_title"`
	Department string `json:"department"`
}

type contactData struct {
	Id           uint         `json:"id"`
	OwnerId      uint         `json:"owner_id"`
	EmployeeId   uint         `json:"employee_id"`
	Notes        string       `json:"notes"`
	CreationDate int64        `json:"creation_date"`
	IsDeleted    bool         `json:"is_deleted"`
	Empl         employeeData `json:"empl"`
}

type updateContactData struct {
	Notes string `json:"notes"`
}

type createCompanyAndOwnerJob struct {
	Company companyData  `json:"company"`
	Owner   employeeData `json:"owner"`
}

type updateCompanyJob struct {
	CompanyId uint              `json:"company_id"`
	OwnerId   uint              `json:"owner_id"`
	Upd       updateCompanyData `json:"update_company"`
}

type deleteCompanyJob struct {
	CompanyId uint `json:"company_id"`
	OwnerId   uint `json:"owner_id"`
}

type createEmployeeJob struct {
	CompanyId uint         `json:"company_id"`
	OwnerId   uint         `json:"owner_id"`
	Employee  employeeData `json:"employee"`
}

type updateEmployeeJob struct {
	CompanyId  uint               `json:"company_id"`
	OwnerId    uint               `json:"owner_id"`
	EmployeeId uint               `json:"employee_id"`
	Upd        updateEmployeeData `json:"upd"`
}

type deleteEmployeeJob struct {
	CompanyId  uint `json:"company_id"`
	OwnerId    uint `json:"owner_id"`
	EmployeeId uint `json:"employee_id"`
}

type createContactJob struct {
	OwnerId    uint `json:"owner_id"`
	EmployeeId uint `json:"employee_id"`
}

type updateContactJob struct {
	OwnerId   uint              `json:"owner_id"`
	ContactId uint              `json:"contact_id"`
	Upd       updateContactData `json:"upd"`
}

type deleteContactJob struct {
	OwnerId   uint `json:"owner_id"`
	ContactId uint `json:"contact_id"`
}
