package buku

import "gorm.io/gorm"

type Service interface {
	FindAll() ([]Buku, error)
	FindByID(ID int) (Buku, error)
	Create(buku BukuRequest) (Buku, error)
}
type service struct {
	repository Repository
}

func NewService(db *gorm.DB) *repository {
	return &repository{db}
}
