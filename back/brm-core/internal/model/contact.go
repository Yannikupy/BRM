package model

type Contact struct {
	Id           uint
	OwnerId      uint
	EmployeeId   uint
	Notes        string
	CreationDate int64
	IsDeleted    bool
	Empl         Employee
}

type UpdateContact struct {
	Notes string
}

type GetContacts struct {
	Limit  int
	Offset int
}
