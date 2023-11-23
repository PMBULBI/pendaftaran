package admin

import (
	"github.com/PMBULBI/types/schemas"
	"gorm.io/gorm"
)

type LevelRepository struct {
	db *gorm.DB
}

func NewLevelRepository(db *gorm.DB) *LevelRepository {
	return &LevelRepository{
		db: db.Session(&gorm.Session{}).Model(&schemas.AdminLevel{}),
	}
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db.Session(&gorm.Session{}).Model(&schemas.Admin{}),
	}
}
