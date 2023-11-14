package pendaftaran

import (
	"context"
	"github.com/PMBULBI/types/schemas"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func (r *Repository) GetByEmailPass(ctx context.Context, email, password string) (data schemas.Pendaftaran, err error) {
	err = r.db.WithContext(ctx).
		Where("email_mhs = ? AND password = ?", email, password).
		First(&data).
		Error
	return

}

func (r *Repository) GetByPhoneNumPass(ctx context.Context, email, password string) (data schemas.Pendaftaran, err error) {
	err = r.db.WithContext(ctx).
		Where("hp_mhs = ? AND password = ?", email, password).
		First(&data).
		Error
	return
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}
