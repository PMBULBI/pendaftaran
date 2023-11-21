package fakultas

import (
	"context"
	pmbulbi "github.com/PMBULBI/types/schemas"
	"gorm.io/gorm"
)

func GetMAllFakultas(conn *gorm.DB, ctx context.Context) (val []pmbulbi.Fakultas, err error) {
	err = conn.
		WithContext(ctx).
		Find(&val).
		Error
	return
}

func GetMOneFakultas(conn *gorm.DB, ctx context.Context, id string) (val pmbulbi.Fakultas, err error) {
	err = conn.
		WithContext(ctx).
		Where("id_fakultas = ? ", id).
		First(&val).
		Error
	return
}

func InsFakultas(conn *gorm.DB, ctx context.Context, val pmbulbi.Fakultas) (err error) {
	err = conn.
		WithContext(ctx).
		Create(&val).
		Error
	return
}
