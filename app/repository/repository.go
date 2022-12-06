package repository

import "gorm.io/gorm"

type Repository struct {
	DB     *gorm.DB
	DbUser string
	DbPass string
	DbPort string
	DbName string
	DbHost string
}

// creates new repository
func NewRepo() *Repository {
	r := Repository{}
	return &r
}
