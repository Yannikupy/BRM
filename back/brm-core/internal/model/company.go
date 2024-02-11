package model

type Company struct {
	Id           uint
	Name         string
	Description  string
	Industry     uint
	OwnerId      uint
	Rating       float64
	CreationDate int64
	IsDeleted    bool
}

type UpdateCompany struct {
	Name        string
	Description string
	Industry    uint
	OwnerId     uint
}
