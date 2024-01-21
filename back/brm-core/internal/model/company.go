package model

import "time"

type Company struct {
	Id           uint
	Name         string
	Description  string
	Industry     uint
	OwnerId      uint
	Rating       float64
	CreationDate time.Time
	IsDeleted    bool
	DeletionDate time.Time
}

type UpdateCompany struct {
	Name        string
	Description string
	Industry    uint
	OwnerId     uint
}
