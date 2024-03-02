package model

import "time"

type Employee struct {
	Id           uint64
	CompanyId    uint64
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

	Limit  uint
	Offset uint
}

type EmployeeByName struct {
	Pattern string

	Limit  uint
	Offset uint
}

type EmployeeCredentials struct {
	Email      string
	Password   string
	EmployeeId uint64
	CompanyId  uint64
}
