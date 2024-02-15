package model

import "time"

type Employee struct {
	Id           uint
	CompanyId    uint
	FirstName    string
	SecondName   string
	Email        string
	Password     string
	JobTitle     string
	Department   string
	CreationDate time.Time
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

type EmployeeCredentials struct {
	Email      string
	Password   string
	EmployeeId uint
	CompanyId  uint
}
