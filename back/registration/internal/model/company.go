package model

type Company struct {
	Id           uint64
	Name         string
	Description  string
	Industry     uint64
	OwnerId      uint64
	Rating       float64
	CreationDate int64
	IsDeleted    bool
}
