package model

type Employee struct {
	Id           uint
	CompanyId    uint
	FirstName    string
	SecondName   string
	Email        string
	Password     string
	JobTitle     string
	Department   string
	CreationDate int64
	IsDeleted    bool
}
