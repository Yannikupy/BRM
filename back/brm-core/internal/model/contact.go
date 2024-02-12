package model

import "time"

type Contact struct {
	Id           uint
	OwnerId      uint
	EmployeeId   uint
	Notes        string
	CreationDate time.Time
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
