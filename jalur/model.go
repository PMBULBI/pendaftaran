package jalur

import (
	"context"
	pmbulbi "github.com/PMBULBI/types/schemas"
	"gorm.io/gorm"
)

func GetAllJalur(conn *gorm.DB, ctx context.Context) (val []pmbulbi.JalurPendaftaran, err error) {
	err = conn.
		WithContext(ctx).
		Find(&val).
		Error
	return
}

func GetOneJalur(conn *gorm.DB, ctx context.Context, id string) (val pmbulbi.JalurPendaftaran, err error) {
	err = conn.
		WithContext(ctx).
		Where("id_jalur = ? ", id).
		First(&val).
		Error
	return
}

func InsJalur(conn *gorm.DB, ctx context.Context, val pmbulbi.JalurPendaftaran) (err error) {
	err = conn.
		WithContext(ctx).
		Create(&val).
		Error
	return
}

func UpdJalur(conn *gorm.DB, ctx context.Context, id string, val pmbulbi.JalurPendaftaran) (err error) {
	err = conn.
		WithContext(ctx).
		Where("id_jalur = ?", id).
		Updates(&val).
		Error
	return
}
