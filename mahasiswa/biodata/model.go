package biodata

import (
	"context"
	pmbulbi "github.com/PMBULBI/types/schemas"
	"gorm.io/gorm"
)

func InsertBioJalur(conn *gorm.DB, ctx context.Context, val pmbulbi.BiodataJalur) (err error) {
	err = conn.
		WithContext(ctx).
		Create(&val).
		Error
	return
}

func InsertBioProdi(conn *gorm.DB, ctx context.Context, val pmbulbi.BiodataProdi) (err error) {
	err = conn.
		WithContext(ctx).
		Create(&val).
		Error
	return
}

func InsertBioDiri(conn *gorm.DB, ctx context.Context, val pmbulbi.BiodataDataDiri) (err error) {
	err = conn.
		WithContext(ctx).
		Create(&val).
		Error
	return
}
func InsertBioOrtu(conn *gorm.DB, ctx context.Context, val pmbulbi.BiodataDataOrtu) (err error) {
	err = conn.
		WithContext(ctx).
		Create(&val).
		Error
	return
}
