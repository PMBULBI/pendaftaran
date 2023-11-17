package admin_user

import (
	"context"
	pmbulbi "github.com/PMBULBI/types/schemas"
	"gorm.io/gorm"
)

func GetAllAdmin(conn *gorm.DB, ctx context.Context) (val []pmbulbi.Admin, err error) {
	err = conn.
		WithContext(ctx).
		Find(&val).
		Error
	return
}

func GetOneAdmin(conn *gorm.DB, ctx context.Context, id string) (val pmbulbi.Admin, err error) {
	err = conn.
		WithContext(ctx).
		Where("id_admin = ? ", id).
		First(&val).
		Error
	return
}

func InsertAdmin(conn *gorm.DB, ctx context.Context, val pmbulbi.Admin) (err error) {
	err = conn.
		WithContext(ctx).
		Create(&val).
		Error
	return
}
