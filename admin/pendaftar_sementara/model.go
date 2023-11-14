package pendaftar_sementara

import (
	"context"
	pmbulbi "github.com/PMBULBI/types/schemas"
	"gorm.io/gorm"
)

func GetAllPendaftar(conn *gorm.DB, ctx context.Context) (val []pmbulbi.Pendaftaran, err error) {
	err = conn.
		WithContext(ctx).
		Order("id DESC").
		Find(&val).
		Error
	return
}

func GetSatuPendaftar(conn *gorm.DB, ctx context.Context, id string) (dest pmbulbi.Pendaftaran, err error) {
	err = conn.
		WithContext(ctx).
		Where("id = ? ", id).
		First(&dest).
		Error
	return
}
