package kecamatan

import (
	pmbulbi "github.com/PMBULBI/types/schemas"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db.Session(&gorm.Session{}).Model(&pmbulbi.WilayahKecamatan{}),
	}
}
