package prodi

import (
	"context"
	pmbulbi "github.com/PMBULBI/types/schemas"
	"gorm.io/gorm"
)

func GetMAllProdi(conn *gorm.DB, ctx context.Context) (val []pmbulbi.ProgramStudi, err error) {
	err = conn.
		WithContext(ctx).
		Find(&val).
		Error
	return
}

func GetMOneProdi(conn *gorm.DB, ctx context.Context, id string) (val pmbulbi.ProgramStudi, err error) {
	err = conn.
		WithContext(ctx).
		Where("kode_program_studi = ? ", id).
		First(&val).
		Error
	return
}

func InsProdi(conn *gorm.DB, ctx context.Context, val pmbulbi.ProgramStudi) (err error) {
	err = conn.
		WithContext(ctx).
		Create(&val).
		Error
	return
}

func UpdProdi(conn *gorm.DB, ctx context.Context, id string, val pmbulbi.ProgramStudi) (err error) {
	err = conn.
		WithContext(ctx).
		Where("kode_program_studi = ?", id).
		Updates(&val).
		Error
	return
}
