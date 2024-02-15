package core

type Employee struct {
	Id           uint
	CompanyId    uint
	FirstName    string
	SecondName   string
	Email        string
	Passport     string
	JobTitle     string
	Department   string
	CreationDate int64
	IsDeleted    bool
}

type UpdateEmployee struct {
	FirstName  string
	SecondName string
	JobTitle   string
	Department string
}

type FilterEmployee struct {
	ByJobTitle bool
	JobTitle   string

	ByDepartment bool
	Department   string

	Limit  int
	Offset int
}

type EmployeeByName struct {
	Pattern string

	Limit  int
	Offset int
}
