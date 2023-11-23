package admin

import (
	"github.com/PMBULBI/types/schemas"
	"gorm.io/gorm"
)

type EmailPassword struct {
	db *gorm.DB
}

func NewEmailPassword(db *gorm.DB) *EmailPassword {
	return &EmailPassword{
		db: db.Session(&gorm.Session{}).Model(&schemas.Admin{}),
	}
}
