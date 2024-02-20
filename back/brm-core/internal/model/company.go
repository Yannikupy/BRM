package model

import "time"

type Company struct {
	Id           uint64
	Name         string
	Description  string
	Industry     uint64
	OwnerId      uint64
	Rating       float64
	CreationDate time.Time
	IsDeleted    bool
}

type UpdateCompany struct {
	Name        string
	Description string
	Industry    uint64
	OwnerId     uint64
}
